package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/resp"
	"github.com/i-coder-robot/gin-demo/service"
	"github.com/i-coder-robot/gin-demo/utils"
	"net/http"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) UserListHandler(c *gin.Context) {

}

func (h *UserHandler) AddUserHandler(c *gin.Context) {
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		panic(err)
	}
	now:=utils.GetNow()
	u.CreateAt=now
	u.UpdateAt=now
	r, err := h.UserSrv.Add(&u)
	if err != nil {
		panic(err)
	}
	entity := resp.Entity{
		Code: http.StatusOK,
		Msg:  "OK",
		Data: r.UserId,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

func (h *UserHandler) EditUserHandler(c *gin.Context) {
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {

}
