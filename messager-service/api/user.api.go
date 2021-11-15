package api

func (a *API) RegisterUserAPI() {
	userRouter := a.Router.Group("/user")
	userRouter.GET("/:id", a.user.Get)
	userRouter.POST("/", a.user.Post)
	userRouter.PUT("/", a.user.Put)
	userRouter.DELETE("/:id", a.Protect(), a.user.Delete)
}
