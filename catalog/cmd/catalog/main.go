package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"go-grpc-graphql-microservices/catalog"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			return err
		}
		return nil
	})
	defer r.Close()
	log.Println("Listenning on Port 8080...")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, 8080))
}
