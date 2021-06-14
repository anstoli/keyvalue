package inmemory

import (
	"sync"

	"github.com/tohast/keyvalue"
)

func NewStore() keyvalue.Store {
	return &store{}
}

type store struct {
	container sync.Map
}

func (s *store) Put(k keyvalue.Key, v keyvalue.Value) {
	s.container.Store(k, v)
}

func (s *store) Get(k keyvalue.Key) keyvalue.Value {
	vI, ok := s.container.Load(k)
	if !ok {
		return nil
	}
	return vI.(keyvalue.Value)
}

func (s *store) Delete(k keyvalue.Key) {
	s.container.Delete(k)
}




