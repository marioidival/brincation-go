package portdomainrepo_test

import (
	"context"
	"testing"

	"github.com/marioidival/brincation-go/internal/repo/portdomainrepo"
	"github.com/marioidival/brincation-go/pkg/db/database"
	"github.com/stretchr/testify/assert"

	rpc "github.com/marioidival/brincation-go/rpc"
)

func TestCreatePort(t *testing.T) {
	repo := portdomainrepo.NewRepoMemor()

	newPort := &rpc.Port{
		Id:          "newId",
		Name:        "Fake",
		City:        "Fake",
		Country:     "Fake",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{12.53, 34.76},
		Province:    "",
		Timezone:    "America/SaoPaulo",
		Unlocs:      []string{},
		Code:        "FK",
	}

	port := repo.CreatePort(context.Background(), newPort)

	assert.NotNil(t, port)
}

func TestGetPort(t *testing.T) {
	repo := portdomainrepo.NewRepoMemor()

	newPort := &rpc.Port{
		Id:          "newId",
		Name:        "Fake",
		City:        "Fake",
		Country:     "Fake",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{12.53, 34.76},
		Province:    "",
		Timezone:    "America/SaoPaulo",
		Unlocs:      []string{},
		Code:        "FK",
	}

	ctx := context.Background()
	_ = repo.CreatePort(ctx, newPort)

	getPort, err := repo.GetPort(ctx, "myFake")
	assert.Nil(t, getPort)
	assert.ErrorIs(t, err, database.DbNotFound)

	getPort, err = repo.GetPort(ctx, newPort.GetId())
	assert.NotNil(t, getPort)
	assert.Nil(t, err)
}

func TestUpdatePort(t *testing.T) {
	repo := portdomainrepo.NewRepoMemor()

	newPort := &rpc.Port{
		Id:          "newId",
		Name:        "Fake",
		City:        "Fake",
		Country:     "Fake",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{12.53, 34.76},
		Province:    "",
		Timezone:    "America/SaoPaulo",
		Unlocs:      []string{},
		Code:        "FK",
	}

	ctx := context.Background()
	_ = repo.CreatePort(ctx, newPort)

	updatedPort, err := repo.UpdatePort(ctx, &rpc.UpdatePortRequest{
		Id:   "myFake",
		Port: newPort,
	})
	assert.Nil(t, updatedPort)
	assert.ErrorIs(t, err, database.DbNotFound)

	updatedPort, err = repo.UpdatePort(ctx, &rpc.UpdatePortRequest{
		Id: newPort.GetId(),
		Port: &rpc.Port{
			Name: "No Fake",
		},
	})
	assert.NotNil(t, updatedPort)
	assert.Nil(t, err)
	assert.Equal(t, "No Fake", updatedPort.GetName())

}
