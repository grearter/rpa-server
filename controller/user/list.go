package user

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/user"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func List(c *gin.Context) {
	userAPIs, err := user.List()
	if err != nil {
		logrus.Errorf("list user err: %s", err.Error())
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	if userAPIs == nil {
		userAPIs = make([]*api.User, 0)
	}

	c.JSON(http.StatusOK, util.NewRespWithData(userAPIs))
	return
}
