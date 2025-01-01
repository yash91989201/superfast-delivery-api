package main

import (
	"context"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	productService "github.com/yash91989201/superfast-delivery-api/services/product"
)

type Config struct {
	DatabaseUrl string `envconfig:"DATABASE_URL" required:"true"`
	ServiceUrl  string `envconfig:"SERVICE_URL" required:"true"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	var r productService.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = productService.NewMongoRepository(cfg.DatabaseUrl)

		if err != nil {
			log.Println(err)
		}

		return
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	defer r.Close(ctx)

	log.Printf("Product service started at %s", cfg.ServiceUrl)

	s := productService.New(r)
	log.Fatal(productService.StartGRPCServer(s, cfg.ServiceUrl))
}
