package api

import "github.com/gin-gonic/gin"

func (a *API) RegisterChatAPI() {
	chatRouter := a.Router.Group("/chat")
	chatRouter.GET("/", func(c *gin.Context) { a.message.Upgrade(c.Writer, c.Request) })
}
