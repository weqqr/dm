package userproxy

func (g *Gateway) Routes() error {

	g.router.POST("/user/create", g.create)
	g.router.GET("/user/:id", g.get)
	g.router.PATCH("/user/:id", g.update)
	g.router.DELETE("/user/:id", g.delete)

	return nil
}
