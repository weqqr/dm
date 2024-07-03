package main

import (
	"context"
	"dm/db"
	"dm/internal/config"
	"dm/internal/rabbitmq"
	"dm/internal/server"
	"dm/internal/user/usercore"
	"dm/internal/user/userservice"
	"dm/internal/user/userstorage"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
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

	go func() {
		if err = rabbitmq.Receive(appConfig.RabbitMQ); err != nil {
			log.Fatalf("Error launging Consumer: %v", err)
		}
	}()

	if err = webServer.Run(context.Background()); err != nil {
		log.Fatalf("Error running web server: %v", err)
	}

	// if err = rabbitmq.Send(appConfig.RabbitMQ); err != nil {
	// 	log.Fatalf("Error launching Publisher: %v", err)
	// }

}
