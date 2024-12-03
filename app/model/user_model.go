package model

type User struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email" example:"test@test.com"`
	Age      int    `json:"age" validate:"required"`
	Password string `json:"password" validate:"required"`
}
