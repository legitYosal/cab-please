package permission

import (
	"github.com/usefss/cab-please/identity/middleware"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) (*middleware.TokenUser, bool) {
	user, exist := c.Get("user")
	if exist {
		user := user.(middleware.TokenUser)
		return &user, true
	} else {
		return nil, false
	}
}
