package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/dao/agent"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Get(c *gin.Context) {
	id := c.Param("id")

	if bson.IsObjectIdHex(id) {
		logrus.Errorf("invalid id: %s", id)
		c.JSON(http.StatusOK, util.NewRespWithMsg(util.CodeParamErr, "id无效"))
		return
	}

	agentAPI, err := agent.Get(bson.ObjectIdHex(id))
	if err != nil {
		logrus.Errorf("get agent err: %s, id: %s", err.Error(), id)
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(agentAPI))
	return
}
