package clients

import (
	"github.com/yash91989201/superfast-delivery-api/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	conn    *grpc.ClientConn
	service pb.ShopServiceClient
}

func NewProductClient(serviceUrl string) (*ProductClient, error) {
	conn, err := grpc.NewClient(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := pb.NewShopServiceClient(conn)

	return &ProductClient{conn, s}, nil
}

func (c *ProductClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *ProductClient) Close() {
	c.conn.Close()
}
