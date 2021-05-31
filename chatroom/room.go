// Package chatroom provides room to communicate
package chatroom

import (
	"errors"
	"sort"
	"sync"

	"github.com/kabi175/chat-app-go/user"
)

const (
	echoBuff = 4
)

func Create(name string, admin user.UserId) *ChatRoom {
	return &ChatRoom{
		mutex:    sync.Mutex{},
		name:     name,
		admin:    admin,
		members:  []user.UserId{admin},
		echoChnl: make(chan Message, echoBuff),
	}
}

type ChatRoom struct {
	mutex    sync.Mutex
	name     string
	admin    user.UserId
	members  []user.UserId
	echoChnl chan Message
}

func KickOut() error {
	return nil
}

func (r *ChatRoom) Join(memberId user.UserId) error {
	if index := r.search(memberId); index != len(r.members) {
		return errors.New("User already exist")
	}
	r.members = append(r.members, memberId)
	return nil
}

func (r *ChatRoom) Leave(memberId user.UserId) error {
	if index := r.search(memberId); index != len(r.members) {
		r.members = append(r.members[:index-1], r.members[index+1:]...)
		return nil
	}
	return errors.New("User not exist")
}

func (r *ChatRoom) search(id user.UserId) int {
	return sort.Search(len(r.members), func(index int) bool {
		return r.members[index] == id
	})
}

func (r *ChatRoom) Echo() {
	var (
		msg Message
	)
	for {
		msg = <-r.echoChnl
		r.mutex.Lock()
		for _, memberId := range r.members {
			user.Ref(memberId).Send(msg)
		}
		r.mutex.Unlock()
	}
}
