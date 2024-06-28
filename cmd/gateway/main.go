package main

import (
	"context"
	"dm/internal/config"
	"log"

	"dm/internal/user/userproxy"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var appConfig config.Config

	err := cleanenv.ReadConfig("gateway.toml", &appConfig)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var u userproxy.Gateway
	if err = u.Run(context.Background(), appConfig.Gateway); err != nil {
		log.Fatalf("Error launching Gateway: %v", err)
	}
}
