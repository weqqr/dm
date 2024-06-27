package main

import (
	"context"
	"log"

	"dm/db"

	"github.com/ilyakaznacheev/cleanenv"

	"dm/internal/config"
	"dm/internal/server"
	"dm/internal/user/usercore"
	"dm/internal/user/userservice"
	"dm/internal/user/userstorage"
)

func main() {
	var appConfig config.Config

	if err := cleanenv.ReadConfig("coordinator.toml", &appConfig); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	database, err := db.Connect(context.TODO(), appConfig.Database)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = db.Migrate(appConfig.Database); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	webServer := server.NewServer(appConfig.Server)

	webServer.AddService(userservice.New(usercore.New(userstorage.New(database))))

	if err = webServer.Run(context.Background()); err != nil {
		log.Fatalf("Error running web server: %v", err)
	}
}
