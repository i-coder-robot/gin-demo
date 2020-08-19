package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/service"
)

type ProductHandler struct {
	ProductSrv service.ProductSrv
}

func (h *ProductHandler) ProductListHandler(c *gin.Context) {

}

func (h *ProductHandler) AddProductHandler(c *gin.Context) {

}

func (h *ProductHandler) EditProductHandler(c *gin.Context) {
}

func (h *ProductHandler) DeleteProductHandler(c *gin.Context) {

}
