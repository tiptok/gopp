package model

import "github.com/tiptok/gopp/pkg/domain"

type Role struct {
	Id       int
	RoleName string
}

type ListUser struct {
	Total int             `json:"total"`
	Users []*domain.Users `json:"users"`
}
