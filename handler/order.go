package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/service"
)

type OrderHandler struct {
	OrderSrv service.OrderSrv
}

func (h *OrderHandler) OrderListHandler(c *gin.Context) {

}

func (h *OrderHandler) AddOrderHandler(c *gin.Context) {

}

func (h *OrderHandler) EditOrderHandler(c *gin.Context) {
}

func (h *OrderHandler) DeleteOrderHandler(c *gin.Context) {

}