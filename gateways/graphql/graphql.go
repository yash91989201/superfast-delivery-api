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
}

func NewGraphQLServer(
	authenticationServiceUrl string,
	userServiceUrl string,
	shopServiceUrl string,
	productServiceUrl string,
) (*Server, error) {
	authenticationClient, err := clients.NewAuthenticationClient(authenticationServiceUrl)
	if err != nil {
		return nil, err
	}

	userClient, err := clients.NewUserClient(userServiceUrl)
	if err != nil {
		return nil, err
	}

	shopClient, err := clients.NewShopClient(shopServiceUrl)
	if err != nil {
		return nil, err
	}

	productClient, err := clients.NewProductClient(productServiceUrl)
	if err != nil {
		return nil, err
	}

	return &Server{
		authenticationClient: authenticationClient,
		userClient:           userClient,
		shopClient:           shopClient,
		productClient:        productClient,
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
