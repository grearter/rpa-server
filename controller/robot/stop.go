package robot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/agent"
	"github.com/grearter/rpa-server/dao/robot"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Stop(c *gin.Context) {
	id := c.Param("id")

	if err := stopTask(bson.ObjectIdHex(id)); err != nil {
		logrus.Errorf("stop robot err: %s, id: %s", err.Error(), id)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}

func stopTask(robotID bson.ObjectId) (err error) {
	robotAPI, err := robot.Get(robotID)
	if err != nil {
		logrus.Errorf("get robot err: %s, id: %s", err.Error(), robotID.Hex())
		return
	}

	if robotAPI.Status != api.RobotStatusRunning {
		err = fmt.Errorf("robot has stopped")
		return
	}

	agentAPI, err := agent.Get(robotAPI.AgentID)
	if err != nil {
		logrus.Errorf("get agent err: %s, id: %s", err.Error(), robotAPI.AgentID.Hex())
		return
	}

	if err = stopRobotByAgent(agentAPI.IP, robotID.Hex()); err != nil {
		return
	}

	return
}

func stopRobotByAgent(agentIP string, robotID string) (err error) {
	url := fmt.Sprintf("http://%s:9000/robots", agentIP)

	resp, err := http.Post(url, "Content-Type/json", nil)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		logrus.Errorf("do http post for stop robot err: %s, url: %s", err.Error(), url)
		return
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("resp.statusCode: %d, url: %s", resp.StatusCode, url)
		err = fmt.Errorf("resp.statusCode=%d", resp.StatusCode)
		return
	}

	return
}
