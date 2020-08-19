package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
)

type ProductSrv interface {
	List(req query.ListQuery) (products []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(product *model.Product) (*model.Product, error)
	Exist(product *model.Product) bool
	Add(product *model.Product) (*model.Product, error)
	Edit(product model.Product) (bool, error)
	Delete(id string) (bool, error)
}

type ProductService struct {
	Repo repository.ProductRepoInterface
}

func (srv *ProductService) List(req query.ListQuery) (products []*model.Product, err error){
	return srv.Repo.List(req)
}
func (srv *ProductService) GetTotal(req *query.ListQuery) (total int, err error){
	return srv.Repo.GetTotal(req)
}
func (srv *ProductService) Get(product *model.Product) (*model.Product, error){
	return srv.Repo.Get(product)
}
func (srv *ProductService) Exist(product *model.Product) bool{
	return srv.Repo.Exist(product)
}
func (srv *ProductService) Add(product *model.Product) (*model.Product, error){
	return srv.Repo.Add(product)
}
func (srv *ProductService) Edit(product model.Product) (bool, error){
	return srv.Repo.Edit(product)
}
func (srv *ProductService) Delete(id string) (bool, error){
	return srv.Repo.Delete(id)
}
