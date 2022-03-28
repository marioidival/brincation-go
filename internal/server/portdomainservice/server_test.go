package portdomainservice_test

import (
	"testing"

	"github.com/marioidival/brincation-go/internal/server/portdomainservice"
	rpc "github.com/marioidival/brincation-go/rpc"
)

func TestServerImplementsTwirpServer(t *testing.T) {
	var i interface{} = new(portdomainservice.Server)
	if _, ok := i.(rpc.PortService); !ok {
		t.Fatal("expected portdomainservice.Server to implement rpc.Port")
	}
}
