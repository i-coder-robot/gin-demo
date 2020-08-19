package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/service"
)

type BannerHandler struct {
	BannerSrv service.BannerSrv
}

func (h *BannerHandler) BannerListHandler(c *gin.Context) {

}

func (h *BannerHandler) AddBannerHandler(c *gin.Context) {

}

func (h *BannerHandler) EditBannerHandler(c *gin.Context) {
}

func (h *BannerHandler) DeleteBannerHandler(c *gin.Context) {

}
