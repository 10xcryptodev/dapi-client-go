package grpc

import (
	"google.golang.org/grpc"
)

var gRPCServer = "seed.evonet.networks.dash.org:3010"

func GetConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(gRPCServer, opts...)
	return conn, err
}
