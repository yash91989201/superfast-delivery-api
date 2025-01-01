package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
	"github.com/vektah/gqlparser/v2/ast"
	graphql "github.com/yash91989201/superfast-delivery-api/gateways/graphql"
)

type Config struct {
	Port                     int    `envconfig:"PORT" required:"true"`
	AuthenticationServiceUrl string `envconfig:"AUTHENTICATION_SERVICE_URL" required:"true"`
	UserServiceUrl           string `envconfig:"USER_SERVICE_URL" required:"true"`
	ShopServiceUrl           string `envconfig:"SHOP_SERVICE_URL" required:"true"`
	ProductServiceUrl        string `envconfig:"PRODUCT_SERVICE_URL" required:"true"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("Failed to get env: %w", err)
	}

	s, err := graphql.NewGraphQLServer(cfg.AuthenticationServiceUrl, cfg.UserServiceUrl, cfg.ShopServiceUrl, cfg.ProductServiceUrl)
	if err != nil {
		log.Fatal("Failed to start graphql server: %w", err)
	}

	srv := handler.New(s.ToExecutableSchema())

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/graphql", srv)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("Graphql server started at :%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
