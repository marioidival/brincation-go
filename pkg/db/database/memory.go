package database

import (
	"sync"

	rpc "github.com/marioidival/brincation-go/rpc"
)

type MemClient struct {
	sync.Mutex
	m map[string]*rpc.Port
}

func NewMemoryDatabase() *MemClient {
	return &MemClient{
		m: make(map[string]*rpc.Port),
	}
}

func (mem *MemClient) Get(key string) (*rpc.Port, error) {
	mem.Lock()
	defer mem.Unlock()
	value, ok := mem.m[key]
	if !ok {
		return nil, DbNotFound
	}
	return value, nil
}

func (mem *MemClient) GetOrCreate(key string, value *rpc.Port) *rpc.Port {
	mem.Lock()
	defer mem.Unlock()
	v, ok := mem.m[key]
	if !ok {
		mem.m[key] = value
		return value
	}
	return v
}
