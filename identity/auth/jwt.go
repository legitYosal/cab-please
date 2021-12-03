package auth

import (
	"fmt"
	"strconv"
	"surge/identity/m/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtExpiryInSeconds int = 600

func CreateJWT(userID uint, userName string) string {
	expiryInUnix := time.Now().UTC().Add(time.Second * time.Duration(jwtExpiryInSeconds)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   strconv.FormatUint(uint64(userID), 10),
		"user_name": userName,
		"exp":       expiryInUnix,
	})
	tokenString, err := token.SignedString([]byte(utils.GetEnv("SECRET_KEY")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func DecodeJWT(tokenString string) (uint64, string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.GetEnv("SECRET_KEY")), nil
	})

	if err != nil {
		return 0, "", true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(claims["user_id"].(string), 10, 32)
		if err != nil {
			panic(err)
		}
		userName := claims["user_name"].(string)
		return userID, userName, false
	} else {
		return 0, "", true
	}
}
