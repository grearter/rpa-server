package user

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/dao/user"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if err := user.Delete(id); err != nil {
		logrus.Errorf("delete user err: %s, id: %s", err.Error(), id)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}
