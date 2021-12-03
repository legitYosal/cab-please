package middleware

import (
	"net/http"
	"surge/identity/m/auth"

	"github.com/gin-gonic/gin"
)

var TokenInHeaderName string = "Authorization"

type TokenUser struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func AuthorizationMiddleware(c *gin.Context) {
	var user TokenUser
	if len(c.Request.Header[TokenInHeaderName]) > 0 {
		userID, userName, notValid := auth.DecodeJWT(c.Request.Header[TokenInHeaderName][0])
		if notValid == false {
			user.ID = userID
			user.Username = userName
			c.Set("user", user)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT is not valid!"})
			return
		}
	}
	c.Next()
}
