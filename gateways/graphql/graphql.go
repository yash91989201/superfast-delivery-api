package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/yash91989201/superfast-delivery-api/common/clients"
)

type Server struct {
	authenticationClient *clients.AuthenticationClient
	userClient           *clients.UserClient
	shopClient           *clients.ShopClient
	productClient        *clients.ProductClient
	inventoryClient      *clients.InventoryClient
	geolocationClient    *clients.GeolocationClient
}

type ServerConfig struct {
	AuthenticationServiceUrl string
	UserServiceUrl           string
	ShopServiceUrl           string
	ProductServiceUrl        string
	InventoryServiceUrl      string
	GeolocationServiceUrl    string
}

func NewGraphQLServer(cfg ServerConfig) (*Server, error) {
	authenticationClient, err := clients.NewAuthenticationClient(cfg.AuthenticationServiceUrl)
	if err != nil {
		return nil, err
	}

	userClient, err := clients.NewUserClient(cfg.UserServiceUrl)
	if err != nil {
		return nil, err
	}

	shopClient, err := clients.NewShopClient(cfg.ShopServiceUrl)
	if err != nil {
		return nil, err
	}

	productClient, err := clients.NewProductClient(cfg.ProductServiceUrl)
	if err != nil {
		return nil, err
	}

	inventoryClient, err := clients.NewInventoryClient(cfg.InventoryServiceUrl)
	if err != nil {
		return nil, err
	}

	geolocationClient, err := clients.NewGeolocationClient(cfg.GeolocationServiceUrl)
	if err != nil {
		return nil, err
	}

	return &Server{
		authenticationClient: authenticationClient,
		userClient:           userClient,
		shopClient:           shopClient,
		productClient:        productClient,
		inventoryClient:      inventoryClient,
		geolocationClient:    geolocationClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
