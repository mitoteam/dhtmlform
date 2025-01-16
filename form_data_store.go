package dhtmlform

import (
	"sync"
	"time"

	"github.com/mitoteam/mttools"
)

type formDataStorer interface {
	Set(fd *FormData)
	Get(build_id string) *FormData
	Remove(build_id string)

	//Called time to time to check and expire outdated form data
	Expire()
}

var formDataStore formDataStorer

func init() {
	formDataStore = newFormDataStore()
}

// default implementation: very simple, primitive map based for now
// TODO: simpleDataStore records expiration
type simpleDataStore struct {
	m sync.RWMutex

	store map[string]*FormData
}

func newFormDataStore() *simpleDataStore {
	return &simpleDataStore{
		store: make(map[string]*FormData),
	}
}

func (ds *simpleDataStore) Set(fd *FormData) {
	mttools.AssertNotNil(fd)
	mttools.AssertNotEmpty(fd.build_id)

	ds.m.Lock()
	defer ds.m.Unlock()

	ds.store[fd.build_id] = fd
}

func (ds *simpleDataStore) Get(build_id string) *FormData {
	ds.m.RLock()
	defer ds.m.RUnlock()

	if fd, ok := ds.store[build_id]; ok {
		return fd
	}

	return nil
}

func (ds *simpleDataStore) Remove(build_id string) {
	ds.m.Lock()
	defer ds.m.Unlock()

	delete(ds.store, build_id)
}

func (ds *simpleDataStore) Expire() {
	//one week by default
	expireEdge := time.Now().Add(-7 * 24 * time.Hour)

	ds.m.Lock()
	defer ds.m.Unlock()

	for build_id, fd := range ds.store {
		if fd.created.Before(expireEdge) {
			delete(ds.store, build_id)
		}
	}
}
