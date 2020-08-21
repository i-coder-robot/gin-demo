package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/enum"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/resp"
	"github.com/i-coder-robot/gin-demo/service"
	"net/http"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery
	err := c.ShouldBindQuery(&q)
	if err != nil {
		panic(err)
	}
	list, err := h.UserSrv.List(&q)
	total, err := h.UserSrv.GetTotal(&q)

	if err != nil {
		panic(err)
	}
	if q.Limit == 0 {
		q.Limit = 5
	}
	ret := int(total % q.Limit)
	ret2 := int(total / q.Limit)
	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	} else {
		totalPage = ret2 + 1
	}

	entity := resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     total,
		TotalPage: totalPage,
		Data:      list,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.Operate_Fail),
		Msg:   enum.Operate_Fail.String(),
		Total: 0,
		Data:  nil,
	}
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.UserSrv.Add(u)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.UserId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.Operate_OK)
	entity.Msg = enum.Operate_OK.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

func (h *UserHandler) EditUserHandler(c *gin.Context) {
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		panic(err)
	}
	b, err := h.UserSrv.Edit(u)
	if err != nil {
		panic(err)
	}
	entity := resp.Entity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: 0,
		Data:  b,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	b, err := h.UserSrv.Delete(id)
	entity := resp.Entity{
		Code:  int(enum.Operate_Fail),
		Msg:   enum.Operate_Fail.String(),
		Total: 0,
		Data:  nil,
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
	if b {
		entity.Code = int(enum.Operate_OK)
		entity.Msg = enum.Operate_OK.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
