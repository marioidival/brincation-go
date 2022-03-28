package portdomainrepo

import (
	"context"

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

func (r *Repo) CreatePort(ctx context.Context, newPort *rpc.Port) *rpc.Port {
	value := r.db.GetOrCreate(ctx, newPort.GetId(), newPort)
	return value
}

func (r *Repo) GetPort(ctx context.Context, id string) (*rpc.Port, error) {
	return r.db.Get(ctx, id)
}

func (r *Repo) UpdatePort(ctx context.Context, req *rpc.UpdatePortRequest) (*rpc.Port, error) {
	return r.db.Update(ctx, req.Id, req.Port)
}
