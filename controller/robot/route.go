package robot

import "github.com/gin-gonic/gin"

func InitRoute(engine *gin.Engine) {
	engine.POST("/robots", Add)
	engine.DELETE("/robots/:id", Delete)
	engine.POST("/robots/start/:id", Start)
	engine.POST("/robots/stop/:id", Stop)
	engine.GET("/robottasks", ListTask)
	engine.GET("/robottasks/:id", Task)

	return
}
