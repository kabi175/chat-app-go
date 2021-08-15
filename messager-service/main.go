package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

}

func main() {
	HOST := os.Getenv("HOST")
	log.Println(HOST)
	// dependency injection
	router := inject()

	server := http.Server{
		Addr:        HOST,
		Handler:     router,
		ReadTimeout: 10 * time.Second,
	}

	// Create and register close signal
	closeSignal := make(chan os.Signal)
	signal.Notify(closeSignal)

	//Listen and serve go-routene
	go func() {
		log.Println("Stating Server on host:", HOST)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Wait for close signal, then gracefull shutdown server
	<-closeSignal
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server Shutdown...")
}
