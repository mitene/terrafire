package server

import (
	"fmt"
	"github.com/mitene/terrafire/internal/api"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"sort"
	"sync"
	"testing"
)

func NewTestDB(t *testing.T) *DB {
	db, err := NewDB("sqlite3", fmt.Sprintf("file:%s.db?mode=memory&cache=shared", t.Name()))
	assert.NoError(t, err)
	return db
}

func (s *DB) createRecords(t *testing.T, records []interface{}) {
	for _, rec := range records {
		err := s.Create(rec).Error
		assert.NoError(t, err)
	}
}

func TestDB_Queue(t *testing.T) {
	// Use file based database, because in-memory database may cause concurrency problems.
	tf, _ := ioutil.TempFile("", "")
	defer func() { _ = os.Remove(tf.Name()) }()
	db, err := NewDB("sqlite3", tf.Name())
	assert.NoError(t, err)

	var exp []string
	for i := 0; i < 100; i++ {
		pj := fmt.Sprintf("pj-%02d", i)
		err := db.enqueue(&api.GetActionResponse{
			Type:      api.GetActionResponse_NONE,
			Project:   pj,
			Workspace: "ws",
		})
		assert.NoError(t, err)
		exp = append(exp, pj)
	}

	wg := &sync.WaitGroup{}
	wg.Add(10)
	mtx := &sync.Mutex{}
	var act []string
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m, err := db.dequeue()
				assert.NoError(t, err)
				if m == nil {
					wg.Done()
					return
				}

				mtx.Lock()
				act = append(act, m.Project)
				mtx.Unlock()
			}
		}()
	}
	wg.Wait()

	sort.Slice(act, func(i, j int) bool { return act[i] < act[j] })
	assert.Equal(t, exp, act)
}
