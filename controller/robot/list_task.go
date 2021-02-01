package robot

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/robottask"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

type listTaskReq struct {
	RobotID string `json:"robotId"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

type listTaskResp struct {
	Total  int              `json:"total"`
	Offset int              `json:"offset"`
	Tasks  []*api.RobotTask `json:"tasks"`
}

func ListTask(c *gin.Context) {
	req := new(listTaskReq)

	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("parse query param err: %s", err.Error())
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, err.Error()))
		return
	}

	if !bson.IsObjectIdHex(req.RobotID) {
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, "id无效"))
		return
	}

	total, taskAPIs, err := robottask.List(bson.ObjectIdHex(req.RobotID), req.Offset, req.Limit)
	if err != nil {
		logrus.Errorf("list robot tasks err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	resp := &listTaskResp{
		Total:  total,
		Offset: req.Offset + len(taskAPIs),
		Tasks:  taskAPIs,
	}

	if resp.Tasks == nil {
		resp.Tasks = make([]*api.RobotTask, 0)
	}

	c.JSON(http.StatusOK, util.NewRespWithData(resp))
	return
}
