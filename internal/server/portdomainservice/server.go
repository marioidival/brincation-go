package portdomainservice

import (
	"context"

	"github.com/marioidival/brincation-go/internal/repo/portdomainrepo"
	rpc "github.com/marioidival/brincation-go/rpc"
)

type Server struct {
	repo portdomainrepo.Repo
}

func NewServer() (*Server, error) {
	return &Server{
		repo: *portdomainrepo.NewRepoMemor(),
	}, nil
}

func (s *Server) GetPort(ctx context.Context, portId *rpc.PortId) (*rpc.Port, error) {
	port, err := s.repo.GetPort(ctx, portId.GetId())
	if err != nil {
		return nil, err
	}
	return port, nil
}

func (s *Server) CreatePort(ctx context.Context, newPort *rpc.Port) (*rpc.Port, error) {
	return s.repo.CreatePort(ctx, newPort), nil
}

func (s *Server) UpdatePort(ctx context.Context, req *rpc.UpdatePortRequest) (*rpc.Port, error) {
	return s.repo.UpdatePort(ctx, req)
}
