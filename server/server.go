// Package server provides app-server
package server

import (
	"context"
	"encoding/json"
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
}

func New() *server {
	return &server{}
}

func hello(w http.ResponseWriter, r *http.Request) {
	type User struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unable to login", http.StatusNotAcceptable)
	}
	fmt.Println(user)
	w.Write([]byte("Secreg key"))
	return
}

func (s *server) Start(addr string) {

	router := mux.NewRouter()
	router.HandleFunc("/login", hello).Methods("POST")
	router.HandleFunc("/ws", wsHandler)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origns := handlers.AllowedOrigins([]string{"*"})

	s.appServer = &http.Server{
		Addr:         addr,
		Handler:      handlers.CORS(headers, methods, origns)(router),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
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
