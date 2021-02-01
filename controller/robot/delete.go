package robot

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/dao/robot"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, "id无效"))
		return
	}

	if err := robot.Delete(bson.ObjectIdHex(id)); err != nil {
		logrus.Errorf("delete robot err: %s", err.Error())
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	return
}
