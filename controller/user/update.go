package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

type updateReq struct {
	Nick  string   `json:"nick" binding:"nick"`
	Mail  string   `json:"mail"`
	Phone string   `json:"phone"`
	Role  api.Role `json:"role" binding:"role"`
}

func Update(c *gin.Context) {
	id := c.Param("id")
	req := new(updateReq)

	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		logrus.Errorf("parse req err: %s, id: %s, req: %+v", err.Error(), id, req)
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}
