package dhtmlform

import "github.com/mitoteam/mttools"

type formDataStorer interface {
	Set(fd *FormData)
	Get(build_id string) *FormData
	Remove(build_id string)
}

var formDataStore formDataStorer

func init() {
	formDataStore = newFormDataStore()
}

// default implementation: very simple, primitive map based for now
// TODO: simpleDataStore records expiration
type simpleDataStore struct {
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

	ds.store[fd.build_id] = fd
}

func (ds *simpleDataStore) Get(build_id string) *FormData {
	if fd, ok := ds.store[build_id]; ok {
		return fd
	}

	return nil
}

func (ds *simpleDataStore) Remove(build_id string) {
	delete(ds.store, build_id)
}
