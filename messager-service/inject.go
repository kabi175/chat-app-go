package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kabi175/chat-app-go/messager/api"
	"github.com/kabi175/chat-app-go/messager/controller"
	"github.com/kabi175/chat-app-go/messager/repository"
	"github.com/kabi175/chat-app-go/messager/service"
)

func inject() *gin.Engine {

	postgresConn, err := NewPostgresClient()
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := repository.NewUserRepo(postgresConn)
	if err != nil {
		log.Fatal(err)
	}

	tokenService := service.NewJwtTokenService()

	userService := service.NewDefaultUserService(userRepo, tokenService)
	userController := controller.NewGinUserController(userService, tokenService)

	redisConn := NewRedisClient()
	messageRepo := repository.NewResidMessageRepo(redisConn)
	messageService := service.NewDefaultMessageService(messageRepo)
	messageController := controller.NewGorillaMessageController(messageService)
	api := api.NewRouter(userController, tokenService, messageController)

	return api.Router
}
