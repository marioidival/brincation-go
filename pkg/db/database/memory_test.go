package database

import (
	"testing"

	rpc "github.com/marioidival/brincation-go/rpc"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	db := NewMemoryDatabase()
	key := "my-id"

	savedValue, err := db.Get(key)
	assert.ErrorIs(t, err, DbNotFound)
	assert.Nil(t, savedValue)
}

func TestGetOrCreate(t *testing.T) {
	db := NewMemoryDatabase()
	key := "my-id"
	value := &rpc.Port{
		Id:          key,
		Name:        "Fake",
		City:        "Fake",
		Country:     "Fake",
		Alias:       []string{"fake", "false"},
		Regions:     []string{},
		Coordinates: []float64{1.090, 44.77},
		Province:    "",
		Timezone:    "America/Recife",
		Unlocs:      []string{"other-key"},
		Code:        "FK",
	}

	savedValue := db.GetOrCreate(key, value)
	assert.Equal(t, value, savedValue, "it's not equal")
}

func TestUpdate(t *testing.T) {
	db := NewMemoryDatabase()
	key := "my-id"
	value := &rpc.Port{
		Id:          key,
		Name:        "Fake",
		City:        "Fake",
		Country:     "Fake",
		Alias:       []string{"fake", "false"},
		Regions:     []string{},
		Coordinates: []float64{1.090, 44.77},
		Province:    "",
		Timezone:    "America/Recife",
		Unlocs:      []string{"other-key"},
		Code:        "FK",
	}

	savedValue := db.GetOrCreate(key, value)
	assert.Equal(t, value, savedValue, "it's not equal")

	toUpdate := &rpc.Port{
		Name: "No Fake",
		Code: "NOFK",
	}

	updated, err := db.Update(key, toUpdate)
	assert.Nil(t, err)
	assert.Equal(t, updated.GetId(), key)
	assert.Equal(t, updated.GetName(), savedValue.GetName())
}
