package user

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/middleware"
)

func InitRoute(engine *gin.Engine) {

	// login
	public := engine.Group("/users")
	public.POST("login", Login)

	// private
	private := engine.Group("users")
	private.Use(middleware.Auth)

	private.POST("", Add)
	private.DELETE("", Delete)
	private.GET("", List)
	private.PUT("/:id", Update)
	private.PUT("/reset_password/:id", ResetPassword)
	private.POST("/logout", Logout)
}
