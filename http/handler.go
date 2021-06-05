package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/domain"
)

type handler struct {
	service Service
}

func NewHandler(service Service) handler {
	return handler{
		service: service,
	}
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user struct {
		UserId   string `json:"userID"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	token, err := h.service.SignUp(user.UserId, user.Password)
	if err != nil {
		http.Error(w, "", http.StatusUnauthorized)
	}
	w.Header().Set("content-type", "application/json")

	resp := make(map[string]string)
	resp["status"] = "successfull created"
	rs, _ := json.Marshal(resp)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "auth", Value: token, Expires: expiration}
	http.SetCookie(w, &cookie)
	w.Write(rs)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var user struct {
		UserId   string `json:"userID"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	token, err := h.service.LogIn(user.UserId, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "auth", Value: token, Expires: expiration}
	http.SetCookie(w, &cookie)

	w.Write([]byte("auth successfull"))
}

func (h *handler) CreateRoom(w http.ResponseWriter, r *http.Request) {

	var reqBody struct {
		UserId string `json:"userId"`
		RoomId string `json:"roomId"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	auth, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}

	userId, err := h.service.Validate(auth.Value)
	if err != nil || userId != reqBody.UserId {

		http.Error(w, "Not authorized", http.StatusNonAuthoritativeInfo)
		return
	}

	h.service.SpinRoom(domain.UserId(reqBody.UserId), domain.RoomId(reqBody.RoomId))
	w.Write([]byte("Room Created"))
}

func (h *handler) JoinRoom(w http.ResponseWriter, r *http.Request) {

	var reqBody struct {
		UserId string `json:"userId"`
		RoomId string `json:"roomId"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}
	auth, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}

	userId, err := h.service.Validate(auth.Value)
	if err != nil || userId != reqBody.UserId {
		http.Error(w, "not authorized", http.StatusNonAuthoritativeInfo)
		return
	}

	err = h.service.SubscribeRoom(domain.UserId(reqBody.UserId), domain.RoomId(reqBody.RoomId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.Write([]byte(" Successfull Joint the room"))
}

func (h *handler) Upgrader(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
	)

	upgrader := websocket.Upgrader{

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	_, auth, err := conn.ReadMessage()
	if err != nil {
		return
	}
	jwt := string(auth)
	userId, err = h.service.Validate(jwt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	h.service.SpinUser(domain.UserId(userId), conn)

	var msg domain.Message

	for {

		err := conn.ReadJSON(&msg)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		h.service.ConsumeMessage(msg)

	}
}
