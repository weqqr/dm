package main

import (
	"dm/internal/config"
	"log"

	"dm/internal/gateway"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var appConfig config.Config

	err := cleanenv.ReadConfig("gateway.toml", &appConfig)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err = gateway.Run(appConfig.Gateway); err != nil {
		log.Fatalf("Error launching Gateway: %v", err)
	}
}
