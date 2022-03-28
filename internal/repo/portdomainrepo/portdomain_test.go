package portdomainrepo_test

import (
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

	port := repo.CreatePort(newPort)

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

	_ = repo.CreatePort(newPort)

	getPort, err := repo.GetPort("myFake")
	assert.Nil(t, getPort)
	assert.ErrorIs(t, err, database.DbNotFound)

	getPort, err = repo.GetPort(newPort.GetId())
	assert.NotNil(t, getPort)
	assert.Nil(t, err)

}
