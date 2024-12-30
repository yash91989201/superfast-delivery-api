package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	shopService "github.com/yash91989201/superfast-delivery-api/services/shop"
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

	var r shopService.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = shopService.NewPgRepository(cfg.DatabaseUrl)

		if err != nil {
			log.Println(err)
		}

		return
	})

	defer r.Close()

	log.Printf("Authentication service started at %s", cfg.ServiceUrl)

	s := shopService.New(r)
	log.Fatal(shopService.StartGRPCServer(s, cfg.ServiceUrl))
}
