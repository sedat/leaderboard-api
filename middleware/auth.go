package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Token string `json:"token"`
}

func TokenVerify(r *http.Request) error {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenVerify(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
