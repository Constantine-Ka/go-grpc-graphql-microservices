package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"go-grpc-graphql-microservices/account"
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

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	defer r.Close()
	log.Println("Listenning on port 8080")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
