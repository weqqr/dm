package gateway

import (
	"dm/internal/rpc"
	"dm/internal/user/userproxy"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address string `toml:"address"`
}

func Run(config Config) error {
	router := gin.Default()

	conn, err := grpc.NewClient("coordinator:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := rpc.NewUserClient(conn)
	var proxy userproxy.UserProxy
	proxy.Client = client

	if err := userproxy.Routes(router, &proxy); err != nil {
		return err
	}

	if err := router.Run(config.Address); err != nil {
		return err
	}

	return nil
}
