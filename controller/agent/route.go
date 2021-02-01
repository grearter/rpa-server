package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/middleware"
)

func InitRoute(engine *gin.Engine) {
	group := engine.Group("/agents")

	// public: agent注册
	group.POST("", Add)

	// private
	private := engine.Group("/agents")
	private.Use(middleware.Auth)

	private.GET("/:id", Get)
	private.GET("", List)
	private.DELETE("", Delete)

	return
}
