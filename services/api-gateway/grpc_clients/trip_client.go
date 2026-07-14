package grpc_clients

import(
	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/trip"
	"google.golang.org/grpc"
)

type TripServiceClient struct {
	Client pb.TripServiceClient
	conn *grpc.ClientConn
}

func NewTripServiceClient() (*TripServiceClient, error)  {
	tripServiceUrl:="trip-service:9093"
	
	conn, err := grpc.NewClient(tripServiceUrl)
	if err!=nil{
		return nil, err
	}

	client := pb.NewTripServiceClient(conn)

	return &TripServiceClient{
		Client : client,
		conn : conn,
	}, nil
}

func (c *TripServiceClient) Close() {
	if c.conn!=nil{
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}

