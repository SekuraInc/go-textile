package core

import (
	"strings"

	"github.com/golang/protobuf/ptypes"
	mh "github.com/multiformats/go-multihash"
	"github.com/textileio/go-textile/pb"
)

// AddComment adds an outgoing comment block
func (t *Thread) AddComment(target string, body string) (mh.Multihash, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	if !t.annotatable(t.config.Account.Address) {
		return nil, ErrNotAnnotatable
	}

	body = strings.TrimSpace(body)
	msg := &pb.ThreadComment{
		Target: target,
		Body:   body,
	}

	res, err := t.commitBlock(msg, pb.Block_COMMENT, true, nil)
	if err != nil {
		return nil, err
	}

	err = t.indexBlock(&pb.Block{
		Id:      res.hash.B58String(),
		Thread:  t.Id,
		Author:  res.header.Author,
		Type:    pb.Block_COMMENT,
		Date:    res.header.Date,
		Parents: res.parents,
		Target:  target,
		Body:    body,
	})
	if err != nil {
		return nil, err
	}

	log.Debugf("added COMMENT to %s: %s", t.Id, res.hash.B58String())

	return res.hash, nil
}

// handleCommentBlock handles an incoming comment block
func (t *Thread) handleCommentBlock(hash mh.Multihash, block *pb.ThreadBlock) (string, string, error) {
	msg := new(pb.ThreadComment)
	err := ptypes.UnmarshalAny(block.Payload, msg)
	if err != nil {
		return "", "", err
	}

	if !t.readable(t.config.Account.Address) {
		return "", "", ErrNotReadable
	}
	if !t.annotatable(block.Header.Address) {
		return "", "", ErrNotAnnotatable
	}

	return msg.Target, msg.Body, nil
}
