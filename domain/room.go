package domain

import (
	"errors"
	"sort"
	"sync"
)

const (
	echoBuff = 4 // maximum messages that can be sent to the roomEcho
)

type RoomId string

type Room struct {
	mutex    sync.Mutex
	name     RoomId
	admin    UserId
	members  []UserId
	echoChnl chan Message
	service  Service
}

func CreateRoom(name RoomId, admin UserId, service Service) *Room {
	return &Room{
		mutex:    sync.Mutex{},
		name:     name,
		admin:    admin,
		members:  []UserId{admin},
		echoChnl: make(chan Message, echoBuff),
		service:  service,
	}
}

func (r *Room) Id() RoomId {
	return r.name
}

func (r *Room) Join(memberId UserId) error {
	r.mutex.Lock()

	if index := r.search(memberId); index != len(r.members) {
		return errors.New("User already exist")
	}
	r.members = append(r.members, memberId)

	r.mutex.Unlock()
	return nil
}

func (r *Room) Leave(memberId UserId) error {
	r.mutex.Lock()

	if index := r.search(memberId); index != len(r.members) {
		r.members = append(r.members[:index-1], r.members[index+1:]...)
		return nil
	}

	r.mutex.Unlock()
	return errors.New("User not exist")
}

func (r *Room) search(id UserId) int {
	return sort.Search(len(r.members), func(index int) bool {
		return r.members[index] == id
	})
}

func (r *Room) Echo() {
	var msg Message
	for {

		msg = <-r.echoChnl

		r.mutex.Lock()
		for _, memberId := range r.members {
			r.service.SendMessage(memberId, msg)
		}

		r.mutex.Unlock()

	}
}

func (r *Room) Consume(message Message) {
	r.echoChnl <- message
}
