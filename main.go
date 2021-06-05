package main

import (
	"github.com/kabi175/chat-app-go/http"
	"github.com/kabi175/chat-app-go/repository"
	"github.com/kabi175/chat-app-go/source"
)

func main() {
	source := source.NewSource()
	service := repository.NewService(&source)
	handler := http.NewHandler(&service)
	server := http.NewServer(&handler)
	server.Config("0.0.0.0:8000")
	server.Start()
}
