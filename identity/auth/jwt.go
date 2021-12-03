package auth

import (
	"surge/identity/m/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtExpiryInSeconds int = 600

func CreateJWT(userID uint, userName string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"user_name": userName,
		"exp":       time.Now().UTC().Add(time.Second * time.Duration(jwtExpiryInSeconds)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(utils.GetEnv("SECRET_KEY")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// func DecodeJWT(tokenString string) (uint, string, error) {

// }
