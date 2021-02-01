package robot

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/robot"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type addReq struct {
	AgentID  bson.ObjectId `json:"agentId" binding:"required"`
	Filepath string        `json:"filepath" binding:"required"`
	Name     string        `json:"name" binding:"required"`
}

func (req *addReq) regular() error {
	if !req.AgentID.Valid() {
		return errors.New("agentId无效")
	}

	if req.Filepath == "" {
		return errors.New("filepath无效")
	}

	return nil
}

func Add(c *gin.Context) {
	req := new(addReq)

	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		logrus.Errorf("parse param err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, err.Error()))
		return
	}

	if err := req.regular(); err != nil {
		logrus.Errorf("valid param err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, err.Error()))
		return
	}

	robotAPI := &api.Robot{
		ID:       bson.NewObjectId(),
		AgentID:  req.AgentID,
		Filepath: req.Filepath,
		Status:   api.RobotStatusStopped,
		Ct:       time.Now().UnixNano(),
		TaskIDs:  nil,
	}

	if err := robot.Add(robotAPI); err != nil {
		logrus.Errorf("add robot err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}
