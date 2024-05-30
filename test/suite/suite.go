package suite

import (
	"context"
	"github.com/cheef/hw-final-project/internal/config"
	pb "github.com/cheef/hw-final-project/pkg/server/grpc/api/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strconv"
	"testing"
	"time"
)

type Suite struct {
	*testing.T
	Cfg    *config.Config
	Client pb.BFAProtectionClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	cfg, err := config.Load()

	if err != nil {
		t.Fatalf("config loading failed: %v", err)
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 300*time.Second)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	grpcAddress := net.JoinHostPort("localhost", strconv.Itoa(cfg.GRPC.Port))
	cc, err := grpc.NewClient(grpcAddress, opts...)

	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	client := pb.NewBFAProtectionClient(cc)

	return ctx, &Suite{
		T:      t,
		Cfg:    cfg,
		Client: client,
	}
}
