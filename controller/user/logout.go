package user

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/middleware"
	"net/http"
)

func Logout(c *gin.Context) {
	middleware.ClearToken(c)
	c.Redirect(http.StatusFound, "/login")
	return
}
