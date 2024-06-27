package gateway

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Address string `toml:"address"`
}

type Router struct {
	router *gin.Engine
}

func (r *Router) Run(config Config) error {
	r.router = gin.Default()

	if err := r.Roots(); err != nil {
		return err
	}

	if err := r.router.Run(config.Address); err != nil {
		return err
	}
	return nil
}
