package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/casbin/casbin/v2"
	"github.com/yash91989201/superfast-delivery-api/common/clients"
)

type Server struct {
	AuthenticationClient *clients.AuthenticationClient
	UserClient           *clients.UserClient
	ShopClient           *clients.ShopClient
	ProductClient        *clients.ProductClient
	InventoryClient      *clients.InventoryClient
	GeolocationClient    *clients.GeolocationClient
	CasbinEnforcer       *casbin.Enforcer
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

	casbinEnforcer, err := casbin.NewEnforcer("rbac/model.conf", "rbac/policy.csv")
	if err != nil {
		return nil, err
	}

	return &Server{
		AuthenticationClient: authenticationClient,
		UserClient:           userClient,
		ShopClient:           shopClient,
		ProductClient:        productClient,
		InventoryClient:      inventoryClient,
		GeolocationClient:    geolocationClient,
		CasbinEnforcer:       casbinEnforcer,
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
	schemaConfig := Config{
		Resolvers: s,
	}

	directives := NewGQLDirective(s.CasbinEnforcer)

	schemaConfig.Directives.RequireAuthRole = directives.RequireAuthRole

	return NewExecutableSchema(schemaConfig)
}
