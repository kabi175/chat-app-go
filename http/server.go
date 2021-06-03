package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type server struct {
	appServer *http.Server
	handler   Handler
}

func New(handler Handler) *server {
	return &server{handler: handler}
}

func (s *server) config() {

	router := mux.NewRouter()
	s.routes(router)
	headers := handlers.AllowedHeaders(
		[]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization",
		},
	)

	methods := handlers.AllowedMethods(
		[]string{
			"GET",
			"POST",
		},
	)

	origns := handlers.AllowedOrigins(
		[]string{
			"*",
		},
	)

	addr := os.Getenv("adderss")

	s.appServer = &http.Server{
		Addr:         addr,
		Handler:      handlers.CORS(headers, methods, origns)(router),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

func (s *server) routes(router *mux.Router) {

	router.HandleFunc("/ws", s.handler.Upgrader)
	router.HandleFunc("/login", s.handler.Login)

}

func (s *server) Start(addr string) {

	go s.ShutDown()

	log.Println("started serving on ", addr)

	err := s.appServer.ListenAndServe()

	if err != nil {
		log.Println(err)
	}

	log.Println("Shutting Down")
}

func (s *server) ShutDown() {
	intr := make(chan os.Signal, 1)
	signal.Notify(intr, os.Interrupt)

	<-intr

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	s.appServer.Shutdown(ctx)
}
