// Package chatroom provides room to communicate
package chatroom

import (
	"errors"
	"sort"
	"sync"

	"github.com/kabi175/chat-app-go/user"
)

const (
	echoBuff = 4 // maximum messages that can be sent to the  chatroom
)

// Create ChatRoom, user.UserId is used as default
// admin for the  chatroom
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

// Add a member a to the chatroom
func (r *ChatRoom) Join(memberId user.UserId) error {
	r.mutex.Lock()
	if index := r.search(memberId); index != len(r.members) {
		return errors.New("User already exist")
	}
	r.members = append(r.members, memberId)
	r.mutex.Unlock()
	return nil
}

// Remove a member from the chatroom
func (r *ChatRoom) Leave(memberId user.UserId) error {
	r.mutex.Lock()
	if index := r.search(memberId); index != len(r.members) {
		r.members = append(r.members[:index-1], r.members[index+1:]...)
		return nil
	}
	r.mutex.Unlock()
	return errors.New("User not exist")
}

func (r *ChatRoom) search(id user.UserId) int {
	return sort.Search(len(r.members), func(index int) bool {
		return r.members[index] == id
	})
}

// Echo the message to all the members of the chatroom
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
