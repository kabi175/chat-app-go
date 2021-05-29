package main

import "github.com/kabi175/chat-app-go/server"

func main() {
	server := server.New()
	server.Start("0.0.0.0:8000")
}
