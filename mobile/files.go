package mobile

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	ipld "gx/ipfs/QmZtNq8dArGfnpCZfx2pUNY7UcjGhVp5qqwQ4hH6mpTMRQ/go-ipld-format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/proto"
	"github.com/textileio/textile-go/pb"

	"github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/ipfs"
	"github.com/textileio/textile-go/mill"
	"github.com/textileio/textile-go/repo"
	"github.com/textileio/textile-go/schema"
)

// FileData is a wrapper around a file data url
type FileData struct {
	Url string `json:"url"`
}

// AddSchema adds a new schema via schema mill
func (m *Mobile) AddSchema(jsonstr string) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}
	added, err := m.addSchema(jsonstr)
	if err != nil {
		return "", err
	}

	return toJSON(added)
}

// PrepareFiles processes a file by path for a thread, but does NOT share it
func (m *Mobile) PrepareFiles(path string, threadId string) ([]byte, error) {
	thrd := m.node.Thread(threadId)
	if thrd == nil {
		return nil, core.ErrThreadNotFound
	}

	if thrd.Schema == nil {
		return nil, core.ErrThreadSchemaRequired
	}

	mdir := &pb.MobilePreparedFiles{
		Dir: &pb.Directory{
			Files: make(map[string]*pb.File),
		},
		Pin: make(map[string]string),
	}

	writeDir := m.RepoPath + "/tmp/"

	mil, err := getMill(thrd.Schema.Mill, thrd.Schema.Opts)
	if err != nil {
		return nil, err
	}
	if mil != nil {
		conf, err := m.getFileConfig(mil, path, "")
		if err != nil {
			return nil, err
		}

		added, err := m.node.AddFile(mil, *conf)
		if err != nil {
			return nil, err
		}

		file, err := toProtoFile(added)
		if err != nil {
			return nil, err
		}
		mdir.Dir.Files[schema.SingleFileTag] = file

		if added.Size >= m.node.Config().Cafe.Client.Mobile.P2PWireLimit {
			mdir.Pin[added.Hash] = writeDir + added.Hash
		}

	} else if len(thrd.Schema.Links) > 0 {

		// determine order
		steps, err := schema.Steps(thrd.Schema.Links)
		if err != nil {
			return nil, err
		}

		// send each link
		for _, step := range steps {
			mil, err := getMill(step.Link.Mill, step.Link.Opts)
			if err != nil {
				return nil, err
			}
			var conf *core.AddFileConfig

			if step.Link.Use == schema.FileTag {
				conf, err = m.getFileConfig(mil, path, "")
				if err != nil {
					return nil, err
				}

			} else {
				if mdir.Dir.Files[step.Link.Use] == nil {
					return nil, errors.New(step.Link.Use + " not found")
				}

				conf, err = m.getFileConfig(mil, path, mdir.Dir.Files[step.Link.Use].Hash)
				if err != nil {
					return nil, err
				}
			}

			added, err := m.node.AddFile(mil, *conf)
			if err != nil {
				return nil, err
			}

			file, err := toProtoFile(added)
			if err != nil {
				return nil, err
			}
			mdir.Dir.Files[step.Name] = file

			if added.Size >= m.node.Config().Cafe.Client.Mobile.P2PWireLimit {
				mdir.Pin[added.Hash] = writeDir + added.Hash
			}
		}
	} else {
		return nil, schema.ErrEmptySchema
	}

	for hash, pth := range mdir.Pin {
		if err := m.writeFileData(hash, pth); err != nil {
			return nil, err
		}
	}

	return proto.Marshal(mdir)
}

// PrepareFilesAsync is the async flavor of PrepareFiles
func (m *Mobile) PrepareFilesAsync(path string, threadId string, cb Callback) {
	go func() {
		cb.Call(m.PrepareFiles(path, threadId))
	}()
}

// AddThreadFiles adds a prepared file to a thread
func (m *Mobile) AddThreadFiles(dir []byte, threadId string, caption string) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	thrd := m.node.Thread(threadId)
	if thrd == nil {
		return "", core.ErrThreadNotFound
	}

	var node ipld.Node
	var keys core.Keys

	mdir := new(pb.Directory)
	if err := proto.Unmarshal(dir, mdir); err != nil {
		return "", err
	}
	if len(mdir.Files) == 0 {
		return "", errors.New("no files found")
	}

	var err error
	if mdir.Files[schema.SingleFileTag] != nil {
		file, err := toRepoFile(mdir.Files[schema.SingleFileTag])
		if err != nil {
			return "", err
		}

		node, keys, err = m.node.AddNodeFromFiles([]repo.File{*file})
		if err != nil {
			return "", err
		}

	} else {

		rdir := make(core.Directory)
		for k, v := range mdir.Files {
			file, err := toRepoFile(v)
			if err != nil {
				return "", err
			}
			rdir[k] = *file
		}

		node, keys, err = m.node.AddNodeFromDirs([]core.Directory{rdir})
		if err != nil {
			return "", err
		}
	}

	if node == nil {
		return "", errors.New("no files found")
	}

	hash, err := thrd.AddFiles(node, caption, keys)
	if err != nil {
		return "", err
	}

	return m.blockInfo(hash)
}

// AddThreadFilesByTarget adds a prepared file to a thread by referencing its top level hash,
// which is the target of an existing files block.
func (m *Mobile) AddThreadFilesByTarget(target string, threadId string, caption string) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	thrd := m.node.Thread(threadId)
	if thrd == nil {
		return "", core.ErrThreadNotFound
	}

	// just need one block w/ the same target, doesn't matter what thread
	blocks := m.node.BlocksByTarget(target)
	if len(blocks) == 0 {
		return "", errors.New("target not found")
	}

	fsinfo, err := m.node.ThreadFile(blocks[0].Id)
	if err != nil {
		return "", err
	}

	var dirs []core.Directory
	var files []repo.File

	for _, info := range fsinfo.Files {
		if len(info.Links) > 0 {
			dirs = append(dirs, info.Links)
		} else if info.File != nil {
			files = append(files, *info.File)
		}
	}

	var node ipld.Node
	var keys core.Keys

	if len(dirs) > 0 {
		node, keys, err = m.node.AddNodeFromDirs(dirs)
		if err != nil {
			return "", err
		}
	} else if len(files) > 0 {
		node, keys, err = m.node.AddNodeFromFiles(files)
		if err != nil {
			return "", err
		}
	}

	if node == nil {
		return "", errors.New("no files found")
	}

	hash, err := thrd.AddFiles(node, caption, keys)
	if err != nil {
		return "", err
	}

	return m.blockInfo(hash)
}

// ThreadFiles calls core ThreadFiles
func (m *Mobile) ThreadFiles(offset string, limit int, threadId string) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	files, err := m.node.ThreadFiles(offset, limit, threadId)
	if err != nil {
		return "", err
	}

	return toJSON(files)
}

// FileData returns a data url of a raw file under a path
func (m *Mobile) FileData(hash string) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	reader, file, err := m.node.FileData(hash)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	prefix := "data:" + file.Media + ";base64,"
	img := &FileData{
		Url: prefix + base64.StdEncoding.EncodeToString(data),
	}

	return toJSON(img)
}

type img struct {
	hash  string
	width int
}

// ImageFileDataForMinWidth returns a data url of an image at or above requested size,
// or the next best option.
// Note: Now that consumers are in control of image sizes via schemas,
// handling this here doesn't feel right. We can eventually push this up to RN, Obj-C, Java.
// Note: pth is <target>/<index>, e.g., "Qm.../0"
func (m *Mobile) ImageFileDataForMinWidth(pth string, minWidth int) (string, error) {
	node, err := ipfs.NodeAtPath(m.node.Ipfs(), pth)
	if err != nil {
		return "", err
	}

	var imgs []img
	for _, link := range node.Links() {
		nd, err := ipfs.NodeAtLink(m.node.Ipfs(), link)
		if err != nil {
			return "", err
		}

		dlink := schema.LinkByName(nd.Links(), core.DataLinkName)
		if dlink == nil {
			continue
		}

		file, err := m.node.File(dlink.Cid.Hash().B58String())
		if err != nil {
			return "", err
		}

		if file.Mill == "/image/resize" {
			if width, ok := file.Meta["width"].(float64); ok {
				imgs = append(imgs, img{hash: file.Hash, width: int(width)})
			}
		}
	}

	if len(imgs) == 0 {
		return "", errors.New("no image files found")
	}

	sort.SliceStable(imgs, func(i, j int) bool {
		return imgs[i].width < imgs[j].width
	})

	var hash string
	for _, img := range imgs {
		if img.width >= minWidth {
			hash = img.hash
			break
		}
	}
	if hash == "" {
		hash = imgs[len(imgs)-1].hash
	}

	return m.FileData(hash)
}

func (m *Mobile) addSchema(jsonstr string) (*repo.File, error) {
	var node schema.Node
	if err := json.Unmarshal([]byte(jsonstr), &node); err != nil {
		return nil, err
	}
	data, err := json.Marshal(&node)
	if err != nil {
		return nil, err
	}

	conf := core.AddFileConfig{
		Input: data,
		Media: "application/json",
	}

	return m.node.AddFile(&mill.Schema{}, conf)
}

func (m *Mobile) getFileConfig(mil mill.Mill, path string, use string) (*core.AddFileConfig, error) {
	var reader io.ReadSeeker
	conf := &core.AddFileConfig{}

	if use == "" {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		reader = f
		_, file := filepath.Split(f.Name())
		conf.Name = file
	} else {
		var file *repo.File
		var err error
		reader, file, err = m.node.FileData(use)
		if err != nil {
			return nil, err
		}
		conf.Name = file.Name
		conf.Use = file.Checksum
	}

	media, err := m.node.GetMedia(reader, mil)
	if err != nil {
		return nil, err
	}
	conf.Media = media
	reader.Seek(0, 0)

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	conf.Input = data

	return conf, nil
}

func (m *Mobile) writeFileData(hash string, pth string) error {
	if err := os.MkdirAll(filepath.Dir(pth), os.ModePerm); err != nil {
		return err
	}

	data, err := m.node.DataAtPath(hash)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(pth, data, 0644)
}

func getMill(id string, opts map[string]string) (mill.Mill, error) {
	switch id {
	case "/blob":
		return &mill.Blob{}, nil
	case "/image/resize":
		width := opts["width"]
		if width == "" {
			return nil, errors.New("missing width")
		}
		quality := opts["quality"]
		if quality == "" {
			quality = "75"
		}
		return &mill.ImageResize{
			Opts: mill.ImageResizeOpts{
				Width:   width,
				Quality: quality,
			},
		}, nil
	case "/image/exif":
		return &mill.ImageExif{}, nil
	default:
		return nil, nil
	}
}

func toProtoFile(file *repo.File) (*pb.File, error) {
	added, err := ptypes.TimestampProto(file.Added)
	if err != nil {
		return nil, err
	}

	return &pb.File{
		Mill:     file.Mill,
		Checksum: file.Checksum,
		Source:   file.Source,
		Opts:     file.Opts,
		Hash:     file.Hash,
		Key:      file.Key,
		Media:    file.Media,
		Name:     file.Name,
		Size:     int64(file.Size),
		Added:    added,
		Meta:     pb.ToStruct(file.Meta),
	}, nil
}

func toRepoFile(file *pb.File) (*repo.File, error) {
	added, err := ptypes.Timestamp(file.Added)
	if err != nil {
		return nil, err
	}

	return &repo.File{
		Mill:     file.Mill,
		Checksum: file.Checksum,
		Source:   file.Source,
		Opts:     file.Opts,
		Hash:     file.Hash,
		Key:      file.Key,
		Media:    file.Media,
		Name:     file.Name,
		Size:     int(file.Size),
		Added:    added,
		Meta:     pb.ToMap(file.Meta),
	}, nil
}
