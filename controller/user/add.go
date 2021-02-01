package user

import (
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/grearter/rpa-server/api"
	"github.com/grearter/rpa-server/dao/user"
	"github.com/grearter/rpa-server/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

type addReq struct {
	ID       string   `json:"id" binding:"required"`
	Nick     string   `json:"nick" binding:"required"`
	Role     api.Role `json:"role" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Mail     string   `json:"mail"`
	Phone    string   `json:"phone"`
}

func Add(c *gin.Context) {
	req := new(addReq)

	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		logrus.Errorf("parse add err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusBadRequest, util.NewRespWithMsg(util.CodeParamErr, err.Error()))
		return
	}

	userAPI := &api.User{
		ID:    req.ID,
		Nick:  req.Nick,
		Mail:  req.Mail,
		Phone: req.Phone,
		Role:  req.Role,
		Auth:  generateAuthData(req.ID, req.Password),
	}

	if err := user.Add(userAPI); err != nil {
		logrus.Errorf("add user err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusInternalServerError, util.NewRespWithMsg(util.CodeSrvErr, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, util.NewRespWithData(nil))
	return
}

const (
	passwordSalt    = "sls403V587"
	defaultPassword = "rpa123456"
)

func generateAuthData(id, password string) string {
	authData := md5.Sum([]byte(id + passwordSalt + password))
	return string(authData[:md5.Size])
}
