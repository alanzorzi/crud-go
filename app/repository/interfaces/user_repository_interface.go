package interfaces

import "github.com/alanzorzi/crud-go/app/model"

type UserRepositoryInterface interface {
	GetAllUsers() ([]*model.User, error)
	GetUserById(id string) ([]*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
	GetUserByEmailAndPassword(email string, password string) ([]*model.User, error)
}
