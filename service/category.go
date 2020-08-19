package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
)

type CategorySrv interface {
	List(req query.ListQuery) (categories []*model.Category, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(category *model.Category) (*model.Category, error)
	Exist(category *model.Category) bool
	Add(category *model.Category) (*model.Category, error)
	Edit(category model.Category) (bool, error)
	Delete(id string) (bool, error)
}

type CategoryService struct {
	Repo repository.CategoryRepoInterface
}

func (srv *CategoryService) List(req query.ListQuery) (categories []*model.Category, err error){
	return srv.Repo.List(req)
}
func (srv *CategoryService) GetTotal(req *query.ListQuery) (total int, err error){
	return srv.Repo.GetTotal(req)
}
func (srv *CategoryService) Get(category *model.Category) (*model.Category, error){
	return srv.Repo.Get(category)
}
func (srv *CategoryService) Exist(category *model.Category) bool{
	return srv.Repo.Exist(category)
}
func (srv *CategoryService) Add(category *model.Category) (*model.Category, error){
	return srv.Repo.Add(category)
}
func (srv *CategoryService) Edit(category model.Category) (bool, error){
	return srv.Repo.Edit(category)
}
func (srv *CategoryService) Delete(id string) (bool, error){
	return srv.Repo.Delete(id)
}