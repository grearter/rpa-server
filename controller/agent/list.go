package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/agent"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

type listReq struct {
	Keyword string `json:"keyword"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

type listResp struct {
	Agents []*api.Agent `json:"agents"`
	Total  int          `json:"total"`
	Offset int          `json:"offset"`
}

func List(c *gin.Context) {
	req := new(listReq)

	if err := c.BindQuery(req); err != nil {
		logrus.Errorf("parse query err: %s, uri: %s", err.Error(), c.Request.RequestURI)
		c.JSON(http.StatusOK, util.NewRespWithMsg(util.CodeParamErr, err.Error()))
		return
	}

	total, agentAPIs, err := agent.List(req.Keyword, req.Offset, req.Limit)
	if err != nil {
		logrus.Errorf("list gent err: %s req: %+v", err.Error(), req)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	resp := &listResp{
		Agents: agentAPIs,
		Total:  total,
		Offset: req.Offset + len(agentAPIs),
	}

	if agentAPIs == nil {
		resp.Agents = make([]*api.Agent, 0)
	}

	c.JSON(http.StatusOK, util.NewRespWithData(resp))
	return
}
