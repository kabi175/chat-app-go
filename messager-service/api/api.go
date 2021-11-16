package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type API struct {
	user    domain.UserController
	auth    domain.TokenService
	message domain.MessageController
	Router  *gin.Engine
}

func NewRouter(user domain.UserController, auth domain.TokenService, message domain.MessageController) *API {
	api := &API{
		user:    user,
		auth:    auth,
		message: message,
		Router:  gin.Default(),
	}
	api.RegisterUserAPI()
	api.RegisterChatAPI()
	return api
}
