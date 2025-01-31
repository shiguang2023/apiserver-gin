// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 用户信息handler

package user

import (
	"apiserver-gin/internal/service"
	"apiserver-gin/pkg/constant"
	"apiserver-gin/pkg/errors"
	"apiserver-gin/pkg/errors/ecode"
	"apiserver-gin/pkg/response"
	"context"

	"github.com/gin-gonic/gin"
)

// Handler 用户业务handler
type Handler struct {
	userSrv service.UserService
}

func NewUserHandler(_userSrv service.UserService) *Handler {
	return &Handler{
		userSrv: _userSrv,
	}
}

func (uh *Handler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetInt64(constant.UserID)
		user, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, ecode.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.JSON(c, nil, user)
		}
	}
}
