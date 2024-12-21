package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	authenticationService "github.com/yash91989201/superfast-delivery-api/services/authentication"
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

	var r authenticationService.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = authenticationService.NewMysqlRepository(cfg.DatabaseUrl)

		if err != nil {
			log.Println(err)
		}

		return
	})

	defer r.Close()

	log.Printf("Authentication service started at %s", cfg.ServiceUrl)

	s := authenticationService.New(r)
	log.Fatal(authenticationService.Start(s, cfg.ServiceUrl))
}
