package interfaces

import "github.com/alanzorzi/crud-go/app/model"

type AuthServiceInterface interface {
	LoginUserServices(email string, password string) (model.User, string, error)
}
