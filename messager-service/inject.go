package main

import (
	"github.com/gorilla/mux"
	"github.com/kabi175/chat-app-go/messager/handler"
	"github.com/kabi175/chat-app-go/messager/repository"
	"github.com/kabi175/chat-app-go/messager/service"
)

func inject() *mux.Router {

	router := mux.NewRouter()
	redisClient := redisDataSource()

	messageRepository := repository.NewMessageRepository(&repository.MessageRepositoryConfig{
		Redis: redisClient,
	})

	statusRepository := repository.NewUserStatusRepository(&repository.UserStatusRepositoryConfig{
		Redis: redisClient,
	})

	messageService := service.NewMessageRepository(&service.MessageServiceConfig{
		MessageRepository: messageRepository,
	})

	statusService := service.NewUserStatusRepository(&service.UserStatusServiceConfig{
		UserStatusRepository: statusRepository,
	})

	userService := service.NewUserService(&service.UserServiceConfig{
		MessageService:    messageService,
		UserStatusService: statusService,
	})

	handler.NewHandler(&handler.HandlerConfig{
		Router: router,
		Us:     userService,
	})

	return router
}
