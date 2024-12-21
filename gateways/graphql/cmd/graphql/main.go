package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
	graphql "github.com/yash91989201/superfast-delivery-api/gateways/graphql"
)

type Config struct {
	Port                     int    `envconfig:"PORT" required:"true"`
	AuthenticationServiceUrl string `envconfig:"AUTHENTICATION_SERVICE_URL" required:"true"`
	UserServiceUrl           string `envconfig:"USER_SERVICE_URL" required:"true"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("Failed to get env: %w", err)
	}

	s, err := graphql.NewGraphQLServer(cfg.AuthenticationServiceUrl, cfg.UserServiceUrl)
	if err != nil {
		log.Fatal("Failed to start graphql server: %w", err)
	}

	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("Superfast Delivery Graphql Gateway", "/graphql"))

	log.Printf("Graphql server started at :%d", cfg.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
