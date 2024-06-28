package userproxy

func (g *Gateway) Roots() error {

	users := g.router.Group("/user")
	{
		users.GET("/test", test)

		users.POST("/create/:name", g.create)
		users.GET("/:id/*get", g.get)
		users.PATCH("/:id/*change", g.change)
		users.DELETE("/:id/*delete", g.delete)
	}
	return nil
}
