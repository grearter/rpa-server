package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/dao/agent"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if err := agent.Delete(id); err != nil {
		logrus.Errorf("delete agent err: %s, id: %s", err.Error(), id)
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}
