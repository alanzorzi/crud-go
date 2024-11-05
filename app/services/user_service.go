package services

import (
	"github.com/alanzorzi/crud-go/app/model"
	repositoryInterface "github.com/alanzorzi/crud-go/app/repository/interfaces"
	"github.com/alanzorzi/crud-go/app/services/interfaces"
)

type userService struct {
	userRepo repositoryInterface.UserRepositoryInterface
}

var _ interfaces.UserServiceInterface = &userService{}

func NewUserService(repo repositoryInterface.UserRepositoryInterface) interfaces.UserServiceInterface {
	return &userService{userRepo: repo}
}

func (s *userService) GetUserById(id string) ([]*model.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) CreateUser(user *model.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}
