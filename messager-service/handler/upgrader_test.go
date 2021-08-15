package handler

import (
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kabi175/chat-app-go/messager/model"
	"github.com/kabi175/chat-app-go/messager/model/mocks"
	"github.com/posener/wstest"
	"github.com/stretchr/testify/assert"
)

func TestUpgrader(t *testing.T) {

	godotenv.Load()
	HOST := os.Getenv("HOST")

	router := mux.NewRouter()
	mockUserService := &mocks.UserService{}
	NewHandler(&HandlerConfig{
		Router: router,
		Us:     mockUserService,
	})
	d := wstest.NewDialer(router)
	ws, _, err := d.Dial("ws://"+HOST+"/ws/chat", nil)
	assert.NoError(t, err)
	defer ws.Close()
	messageByClient := struct {
		Type    int8             `json:"type"`
		Message model.Message    `json:"message"`
		Status  model.UserStatus `json:"status"`
	}{}
	//[TODO]
	err = ws.WriteJSON(messageByClient)
	assert.NoError(t, err)
	//[TODO]
	var msg interface{}
	err = ws.ReadJSON(msg)
	assert.NoError(t, err)
}
