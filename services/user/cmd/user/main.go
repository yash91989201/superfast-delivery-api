package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	userService "github.com/yash91989201/superfast-delivery-api/services/user"
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

	var r userService.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = userService.NewMysqlRepository(cfg.DatabaseUrl)

		if err != nil {
			log.Println(err)
		}

		return
	})

	defer r.Close()

	log.Printf("User service started at %s", cfg.ServiceUrl)

	s := userService.New(r)
	log.Fatal(userService.Start(s, cfg.ServiceUrl))
}
