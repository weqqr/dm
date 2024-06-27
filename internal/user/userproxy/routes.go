package userproxy

import (
	"dm/internal/rpc"

	"github.com/gin-gonic/gin"
)

type UserProxy struct {
	Client rpc.UserClient
}

func Routes(router *gin.Engine, proxy *UserProxy) error {
	user := router.Group("/user")
	{
		user.POST("/create", proxy.create)
		user.GET("/:id", proxy.get)
		user.PATCH("/:id", proxy.update)
		user.DELETE("/:id", proxy.delete)
	}

	return nil
}
