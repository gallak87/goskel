package server

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type User struct {
	ID   int64
	Name string
}

type UserService interface {
	GetUser(id int64) (User, error)
}
