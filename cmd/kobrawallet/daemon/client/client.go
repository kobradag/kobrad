package client

import (
	"context"
	"time"

	"github.com/kobradag/kobrad/cmd/kobrawallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/kobradag/kobrad/cmd/kobrawallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kobrawalletd server, and returns the client instance
func Connect(address string) (pb.PyrinwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("kobrawallet daemon is not running, start it with `kobrawallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewPyrinwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
