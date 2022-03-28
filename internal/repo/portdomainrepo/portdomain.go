package portdomainrepo

import (
	"github.com/marioidival/brincation-go/pkg/db/database"
	rpc "github.com/marioidival/brincation-go/rpc"
)

type Repo struct {
	db *database.MemClient
}

func NewRepoMemor() *Repo {
	return &Repo{
		db: database.NewMemoryDatabase(),
	}
}

func (r *Repo) CreatePort(newPort *rpc.Port) *rpc.Port {
	value := r.db.GetOrCreate(newPort.GetId(), newPort)
	return value
}

func (r *Repo) GetPort(id string) (*rpc.Port, error) {
	return r.db.Get(id)
}
