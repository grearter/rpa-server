package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/agent"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Add(c *gin.Context) {
	hostname := c.Param("hostname")

	agentAPI := &api.Agent{
		ID:       bson.NewObjectId(),
		Hostname: hostname,
		IP:       c.ClientIP(),
		Ct:       time.Now().UnixNano(),
	}

	if err := agent.Add(agentAPI); err != nil {
		logrus.Errorf("add agent err: %s, agent: %+v", err.Error(), agentAPI)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, agentAPI.ID)
	return
}
