package interfaces

import "github.com/alanzorzi/crud-go/app/model"

type UserServiceInterface interface {
	GetUserById(id string) ([]*model.User, error)
	GetAllUsers() ([]*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
}
