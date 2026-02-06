package acoustics

import (
	"context"

	api "github.com/YumikoKawaii/rpc.com/protobuf/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	client api.EchoClient
}

func NewGRPCClient(endpoint string) (Client, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &grpcClient{
		client: api.NewEchoClient(conn),
	}, nil
}

func (c *grpcClient) Charge(ctx context.Context, request ChargeRequest) (ChargeResponse, error) {
	_, err := c.client.Charge(ctx, &api.ChargeRequest{})
	if err != nil {
		return ChargeResponse{}, err
	}
	return ChargeResponse{}, nil
}

func (c *grpcClient) Discharge(ctx context.Context, request DischargeRequest) (DischargeResponse, error) {
	_, err := c.client.Discharge(ctx, &api.DischargeRequest{})
	if err != nil {
		return DischargeResponse{}, err
	}
	return DischargeResponse{}, nil
}
