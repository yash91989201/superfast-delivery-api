package clients

import (
	"context"

	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GeolocationClient struct {
	conn    *grpc.ClientConn
	service pb.GeolocationServiceClient
}

func NewGeolocationClient(serviceUrl string) (*GeolocationClient, error) {
	conn, err := grpc.NewClient(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewGeolocationServiceClient(conn)

	return &GeolocationClient{conn: conn, service: client}, nil
}

func (c *GeolocationClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *GeolocationClient) Close() {
	c.conn.Close()
}

func (c *GeolocationClient) ReverseGeocode(ctx context.Context, req *pb.ReverseGeocodeReq) (*pb.AddressDetail, error) {
	return c.service.ReverseGeocode(ctx, req)
}
