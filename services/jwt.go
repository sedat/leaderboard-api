package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserLogin struct {
	UserID      string
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
}

func (userLogin *UserLogin) GenerateToken() (string, error) {
	var err error

	tokenClaims := jwt.MapClaims{}
	tokenClaims["authorized"] = true
	tokenClaims["user_id"] = userLogin.UserID
	tokenClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenCreated, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenCreated, nil
}
