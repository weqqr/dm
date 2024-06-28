package userproxy

import (
	"context"

	"dm/internal/rpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address string `toml:"address"`
}

type Gateway struct {
	router *gin.Engine
	ctx    context.Context
	client rpc.UserClient
}

func (g *Gateway) Run(ctx context.Context, config Config) error {
	g.router = gin.Default()
	g.ctx = ctx

	conn, err := grpc.NewClient("coordinator:80", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	g.client = rpc.NewUserClient(conn)

	if err := g.Roots(); err != nil {
		return err
	}

	if err := g.router.Run(config.Address); err != nil {
		return err
	}

	return nil
}
