package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type Handler struct {
	r  *mux.Router
	us domain.UserService
}

type HandlerConfig struct {
	Router *mux.Router
	Us     domain.UserService
}

func NewHandler(c *HandlerConfig) {
	h := Handler{
		r:  c.Router,
		us: c.Us,
	}
	h.r.HandleFunc("/ws/chat", h.Upgrader)
}

func (Handler) TodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo handler endpoint"))
}
