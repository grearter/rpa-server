package robot

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/robot"
	"github.com/grearter/rpa-server/dao/robottask"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Start(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, "id无效"))
		return
	}

	robotAPI, err := robot.Get(bson.ObjectIdHex(id))
	if err != nil {
		logrus.Errorf("get robot err: %s, id: %s", err.Error(), id)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	if robotAPI.Status == api.RobotStatusRunning {
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, "robot正在运行中"))
		return
	}

	uid := c.GetString("uid")
	unick := c.GetString("unick")
	now := time.Now()

	taskAPI := &api.RobotTask{
		ID:         bson.NewObjectId(),
		RobotID:    robotAPI.ID,
		CreateID:   uid,
		CreateName: unick,
		Ct:         now.UnixNano(),
		Ut:         now.UnixNano(),
		Messages:   nil,
	}

	if err := runTask(bson.ObjectIdHex(id), taskAPI); err != nil {
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}

func runTask(id bson.ObjectId, taskAPI *api.RobotTask) (err error) {
	if err = robot.UpdateStatus(id, api.RobotStatusRunning); err != nil {
		logrus.Errorf("update robot status err: %s", err.Error())
		return
	}

	if err = robottask.Add(taskAPI); err != nil {
		logrus.Errorf("add task err: %s", err.Error())
		_ = robot.UpdateStatus(id, api.RobotStatusStopped)
		return
	}

	return
}
