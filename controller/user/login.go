package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/grearter/rpa-server/dao/user"
	"github.com/grearter/rpa-server/middleware"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

type loginReq struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登陆
func Login(c *gin.Context) {
	req := new(loginReq)

	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		logrus.Errorf("parse req err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusOK, util.NewRespWithMsg(util.CodeOK, err.Error()))
		return
	}

	auth := generateAuthData(req.ID, req.Password)

	ok, err := user.Auth(req.ID, auth)
	if err != nil {
		logrus.Errorf("auth user err: %s, id: %s", err.Error(), req.ID)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	if !ok {
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, "用户名或密码错误"))
		return
	}

	middleware.SetToken(c, req.ID)
	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	//c.Redirect(http.StatusFound, "/")
	return
}
