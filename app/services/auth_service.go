package services

import (
	"errors"
	"github.com/alanzorzi/crud-go/app/model"
	repositoryInterface "github.com/alanzorzi/crud-go/app/repository/interfaces"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET_KEY = "JWT_SECRET_KEY"

type authService struct {
	userRepo repositoryInterface.UserRepositoryInterface
}

func NewAuthService(repo repositoryInterface.UserRepositoryInterface) *authService {
	return &authService{userRepo: repo}
}

func (ud *authService) LoginUserServices(email string, password string) (model.User, string, error) {

	user, err := ud.userRepo.GetUserByEmailAndPassword(
		email,
		password,
	)
	if err != nil {
		return model.User{}, "", err
	}

	if len(user) < 1 {
		return model.User{}, "", errors.New("Dados inválidos")
	}

	token, err := GenerateJWT(user[0].ID)

	if err != nil {
		return model.User{}, "", err
	}

	return *user[0], token, nil
}

var SecretKey = []byte(JWT_SECRET_KEY)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token válido por 24 horas
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
