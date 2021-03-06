package grpc

import (
	org_dash_platform_dapi_v0 "github.com/10xcryptodev/dapi-client-go/grpc/protos"
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

func GetCoreClient() (org_dash_platform_dapi_v0.CoreClient, error) {
	gRPCconn, err := GetConnection()

	if err != nil {
		return nil, err
	}

	coreClient := org_dash_platform_dapi_v0.NewCoreClient(gRPCconn)

	return coreClient, nil
}

func GetPlatformClient() (org_dash_platform_dapi_v0.PlatformClient, error) {
	gRPCconn, err := GetConnection()

	if err != nil {
		return nil, err
	}

	platformClient := org_dash_platform_dapi_v0.NewPlatformClient(gRPCconn)

	return platformClient, nil
}

func GetTransactionStreamClient() (org_dash_platform_dapi_v0.TransactionsFilterStreamClient, error) {
	gRPCconn, err := GetConnection()

	if err != nil {
		return nil, err
	}

	platformClient := org_dash_platform_dapi_v0.NewTransactionsFilterStreamClient(gRPCconn)

	return platformClient, nil
}
