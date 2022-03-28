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

func (mem *MemClient) Update(key string, updatedValue *rpc.Port) (*rpc.Port, error) {
	currentValue, err := mem.Get(key)
	if err != nil {
		return nil, err
	}
	mem.Lock()
	defer mem.Unlock()

	if updatedValue.Name != "" {
		currentValue.Name = updatedValue.Name
	}
	if updatedValue.Country != "" {
		currentValue.Country = updatedValue.Country
	}
	if updatedValue.Province != "" {
		currentValue.Province = updatedValue.Province
	}
	if updatedValue.Timezone != "" {
		currentValue.Timezone = updatedValue.Timezone
	}
	if updatedValue.Code != "" {
		currentValue.Code = updatedValue.Code
	}
	if updatedValue.Alias != nil {
		currentValue.Alias = updatedValue.Alias
	}
	if updatedValue.Regions != nil {
		currentValue.Regions = updatedValue.Regions
	}
	if updatedValue.Coordinates != nil {
		currentValue.Coordinates = updatedValue.Coordinates
	}
	if updatedValue.Unlocs != nil {
		currentValue.Unlocs = updatedValue.Unlocs
	}
	mem.m[key] = currentValue
	return currentValue, nil
}
