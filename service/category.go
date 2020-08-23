package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
)

type CategorySrv interface {
	List(req *query.ListQuery) (Categories []*model.CategoryResult, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(Category model.Category) (*model.Category, error)
	Edit(Category model.Category) (bool, error)
	Delete(c model.Category) (bool, error)
}

type CategoryService struct {
	Repo repository.CategoryRepoInterface
}

func (srv *CategoryService) List(req *query.ListQuery) (categories []*model.CategoryResult, err error){
	return srv.Repo.List(req)
}
func (srv *CategoryService) GetTotal(req *query.ListQuery) (total int64, err error){
	return srv.Repo.GetTotal(req)
}
func (srv *CategoryService) Get(id string) ([]*model.CategoryResult, error){
	return srv.Repo.Get(id)
}
func (srv *CategoryService) Exist(category model.Category) *model.Category{
	return srv.Repo.Exist(category)
}
func (srv *CategoryService) ExistByCategoryID(id string) *model.Category{
	return srv.Repo.ExistByCategoryID(id)
}

func (srv *CategoryService) Add(category model.Category) (*model.Category, error){
	return srv.Repo.Add(category)
}
func (srv *CategoryService) Edit(category model.Category) (bool, error){
	return srv.Repo.Edit(category)
}
func (srv *CategoryService) Delete(c model.Category) (bool, error){
	return srv.Repo.Delete(c)
}