package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/usefss/cab-please/identity/utils"
)

var jwtExpiryInSeconds int = 600
var jwtSecretKey string

func CreateJWT(userID uint, userName string, isAdmin bool) string {
	expiryInUnix := time.Now().UTC().Add(time.Second * time.Duration(jwtExpiryInSeconds)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":        uuid.New().String(),
		"user_id":    strconv.FormatUint(uint64(userID), 10),
		"is_admin":   isAdmin,
		"user_name":  userName,
		"exp":        expiryInUnix,
		"token_type": "access",
	})

	if jwtSecretKey == "" {
		jwtSecretKey = utils.GetEnv("SECRET_KEY")
	}

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func DecodeJWT(tokenString string) (uint64, string, bool) {
	if jwtSecretKey == "" {
		jwtSecretKey = utils.GetEnv("SECRET_KEY")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
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
