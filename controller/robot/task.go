package robot

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/dao/robottask"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Task(c *gin.Context) {

	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, "id无效"))
		return
	}

	taskAPI, err := robottask.Get(bson.ObjectIdHex(id))
	if err != nil {
		logrus.Errorf("get task err: %s, id: %s", err.Error(), id)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewRespWithData(taskAPI))
	return
}
