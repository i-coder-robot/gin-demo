package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
)

type BannerSrv interface {
	List(req query.ListQuery) (banners []*model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(banner *model.Banner) (*model.Banner, error)
	Exist(banner *model.Banner) bool
	Add(banner *model.Banner) (*model.Banner, error)
	Edit(banner model.Banner) (bool, error)
	Delete(id string) (bool, error)
}

type BannerService struct {
	Repo repository.BannerRepoInterface
}

func (srv *BannerService) List(req query.ListQuery) (banners []*model.Banner, err error){
	return srv.Repo.List(req)
}
func (srv *BannerService) GetTotal(req *query.ListQuery) (total int, err error){
	return srv.Repo.GetTotal(req)
}
func (srv *BannerService) Get(banner *model.Banner) (*model.Banner, error){
	return srv.Repo.Get(banner)
}
func (srv *BannerService) Exist(banner *model.Banner) bool{
	return srv.Repo.Exist(banner)
}
func (srv *BannerService) Add(banner *model.Banner) (*model.Banner, error){
	return srv.Repo.Add(banner)
}
func (srv *BannerService) Edit(banner model.Banner) (bool, error){
	return srv.Repo.Edit(banner)
}
func (srv *BannerService) Delete(id string) (bool, error){
	return srv.Repo.Delete(id)
}
