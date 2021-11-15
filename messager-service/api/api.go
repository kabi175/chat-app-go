package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type API struct {
	user   domain.UserController
	auth   domain.TokenService
	Router *gin.Engine
}

func NewRouter(user domain.UserController, auth domain.TokenService) *API {
	api := &API{
		user:   user,
		auth:   auth,
		Router: gin.Default(),
	}
	api.RegisterUserAPI()
	return api
}
