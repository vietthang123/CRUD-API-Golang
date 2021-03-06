package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"developer-orientenergy-golang/internal/pkg/constants"
	jwt2 "developer-orientenergy-golang/internal/pkg/jwt"
)

type Authentication interface {
	GenerateToken(user UserInfo) (string, error)
}

type authentication struct {
}

func NewAuthentication() *authentication {
	return &authentication{}
}

func (au *authentication) GenerateToken(user UserInfo) (string, error) {
	if !user.IsActive {
		return "", errors.New("User Is Disable Or not found ")
	}
	claims := jwt2.CustomJWTClaims{
		UserInfo: "test",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24, // expire time
			Issuer:    "OrientSchool",              //signal issuer
		},
	}
	authToken, _ := jwt2.CreateJWTToken(claims, []byte(constants.AuthTokenSalt))
	return authToken, nil
}
