package grpc_clients

import (
	"os"

	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverServiceClient struct {
	Client pb.DriverServiceClient
	conn   *grpc.ClientConn
}

func NewDriverServiceClient() (*DriverServiceClient, error) {
	driverserviceurl := os.Getenv("DRIVER_SERVICE_URL")
	if driverserviceurl == "" {
		driverserviceurl = "localhost:8084"
	}

	conn, err := grpc.NewClient(driverserviceurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewDriverServiceClient(conn)

	return &DriverServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *DriverServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
