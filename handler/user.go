package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/service"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) UserListHandler(c *gin.Context) {

}

func (h *UserHandler) AddUserHandler(c *gin.Context) {

}

func (h *UserHandler) EditUserHandler(c *gin.Context) {
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {

}

