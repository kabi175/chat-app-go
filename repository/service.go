package service

import "github.com/kabi175/chat-app-go/domain"

type Datasource interface {
	GetUser(domain.UserId) domain.User
	Store(domain.UserId, domain.Message)
}

type service struct {
	datasource Datasource
}

func NewService(datasource Datasource) *service {
	return &service{
		datasource: datasource,
	}
}

func (s *service) sendMesage(id domain.UserId, msg domain.Message) {
	user := s.datasource.GetUser(id)
	if ok := user.Send(msg); ok != true {
		s.datasource.Store(id, msg)
	}
	return
}
