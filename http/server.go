package http

import (
	"context"
	"fmt"
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

func NewServer(handler Handler) *server {
	return &server{handler: handler}
}

func (s *server) Config(addr string) {

	router := mux.NewRouter()
	s.routes(router)
	headers := handlers.AllowedHeaders(
		[]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization",
		},
	)
	fmt.Println("Started serving on", addr)
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

	s.appServer = &http.Server{
		Addr:         addr,
		Handler:      handlers.CORS(headers, methods, origns)(router),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

func (s *server) routes(router *mux.Router) {

	router.HandleFunc("/chat", s.handler.Upgrader)
	router.HandleFunc("/user/login", s.handler.Login)
	router.HandleFunc("/user/signup", s.handler.SignUp)
	router.HandleFunc("/room/create", s.handler.CreateRoom)
	router.HandleFunc("/room/join", s.handler.JoinRoom)
}

func (s *server) Start() {

	go s.ShutDown()

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
