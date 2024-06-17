package main

import (
	"context"
	"dm/internal/config"
	"dm/internal/db"
	"dm/internal/server"
	"dm/internal/user/usercore"
	"dm/internal/user/userservice"
	"dm/internal/user/userstorage"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

func main() {
	var appConfig config.Config

	err := cleanenv.ReadConfig("coordinator.toml", &appConfig)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	database, err := db.Connect(context.TODO(), appConfig.Database)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	webServer := server.NewServer(appConfig.Server)

	webServer.AddService(userservice.New(usercore.New(userstorage.New(database))))

	if err = webServer.Run(context.Background()); err != nil {
		log.Fatalf("Error running web server: %v", err)
	}
}
