package gateway

func (r *Router) Roots() error {
	users := r.router.Group("/user")
	{
		users.GET("/test", test)
	}
	return nil
}
