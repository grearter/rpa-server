package user

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/dao/user"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ResetPassword(c *gin.Context) {
	id := c.Param("id")

	auth := generateAuthData(id, defaultPassword)

	if err := user.UpdateAuth(id, auth); err != nil {
		logrus.Errorf("update user auth err: %s, id: %s", err.Error(), id)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}
