package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/yash91989201/superfast-delivery-api/common/utils"
	authenticationService "github.com/yash91989201/superfast-delivery-api/services/authentication"
)

type Config struct {
	DatabaseUrl  string `envconfig:"DATABASE_URL" required:"true"`
	ServiceUrl   string `envconfig:"SERVICE_URL" required:"true"`
	JwtSecretKey string `envconfig:"JWT_SECRET_KEY" required:"true"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	tokenManager := utils.NewTokenManager(cfg.JwtSecretKey)

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
	log.Fatal(authenticationService.StartGRPCServer(s, tokenManager, cfg.ServiceUrl))
}
