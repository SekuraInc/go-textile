package db

import (
	"database/sql"
	"github.com/textileio/textile-go/repo"
	"sync"
)

type ThreadPeerDB struct {
	modelStore
}

func NewThreadPeerStore(db *sql.DB, lock *sync.Mutex) repo.ThreadPeerStore {
	return &ThreadPeerDB{modelStore{db, lock}}
}

func (c *ThreadPeerDB) Add(peer *repo.ThreadPeer) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	stm := `insert into thread_peers(id, threadId) values(?,?)`
	stmt, err := tx.Prepare(stm)
	if err != nil {
		log.Errorf("error in tx prepare: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		peer.Id,
		peer.ThreadId,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *ThreadPeerDB) List() []repo.ThreadPeer {
	c.lock.Lock()
	defer c.lock.Unlock()
	stm := "select * from thread_peers;"
	return c.handleQuery(stm)
}

func (c *ThreadPeerDB) ListById(id string) []repo.ThreadPeer {
	c.lock.Lock()
	defer c.lock.Unlock()
	stm := "select * from thread_peers where id='" + id + "';"
	return c.handleQuery(stm)
}

func (c *ThreadPeerDB) ListByThread(threadId string) []repo.ThreadPeer {
	c.lock.Lock()
	defer c.lock.Unlock()
	stm := "select * from thread_peers where threadId='" + threadId + "';"
	return c.handleQuery(stm)
}

func (c *ThreadPeerDB) Count(distinct bool) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	var stm string
	if distinct {
		stm = "select Count(distinct id) from thread_peers;"
	} else {
		stm = "select Count(*) from thread_peers;"
	}
	row := c.db.QueryRow(stm)
	var count int
	row.Scan(&count)
	return count
}

func (c *ThreadPeerDB) Delete(id string, threadId string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, err := c.db.Exec("delete from thread_peers where id=? and threadId=?", id, threadId)
	return err
}

func (c *ThreadPeerDB) DeleteById(id string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, err := c.db.Exec("delete from thread_peers where id=?", id)
	return err
}

func (c *ThreadPeerDB) DeleteByThread(threadId string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, err := c.db.Exec("delete from thread_peers where threadId=?", threadId)
	return err
}

func (c *ThreadPeerDB) handleQuery(stm string) []repo.ThreadPeer {
	var ret []repo.ThreadPeer
	rows, err := c.db.Query(stm)
	if err != nil {
		log.Errorf("error in db query: %s", err)
		return nil
	}
	for rows.Next() {
		var id, threadId string
		if err := rows.Scan(&id, &threadId); err != nil {
			log.Errorf("error in db scan: %s", err)
			continue
		}
		block := repo.ThreadPeer{
			Id:       id,
			ThreadId: threadId,
		}
		ret = append(ret, block)
	}
	return ret
}
