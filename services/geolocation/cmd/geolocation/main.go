package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/yash91989201/superfast-delivery-api/services/geolocation"
	"googlemaps.github.io/maps"
)

type Config struct {
	RedisUrl        string `envconfig:"REDIS_URL" required:"true"`
	GoogleMapApiKey string `envconfig:"GOOGLE_MAP_API_KEY" required:"true"`
	ServiceUrl      string `envconfig:"SERVICE_URL" required:"true"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	var r geolocation.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = geolocation.NewRedisRepository(cfg.RedisUrl, 30*24*time.Hour)
		if err != nil {
			log.Println(err)
		}
		return
	})

	defer r.Close()

	mapClient, err := maps.NewClient(maps.WithAPIKey(cfg.GoogleMapApiKey))
	if err != nil {
		log.Fatalf("failed to create google map client: %+v", err)
	}

	log.Printf("geolocation service started at %s", cfg.ServiceUrl)

	s := geolocation.New(mapClient, r)
	log.Fatal(geolocation.StartGRPCServer(s, cfg.ServiceUrl))
}
