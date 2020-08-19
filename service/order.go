package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
)

type OrderSrv interface {
	List(req query.ListQuery) (orders []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(order *model.Order) (*model.Order, error)
	Exist(order *model.Order) bool
	Add(order *model.Order) (*model.Order, error)
	Edit(order model.Order) (bool, error)
	Delete(id string) (bool, error)
}

type OrderService struct {
	Repo repository.OrderRepoInterface
}

func (srv *OrderService) List(req query.ListQuery) (orders []*model.Order, err error){
	return srv.Repo.List(req)
}
func (srv *OrderService) GetTotal(req *query.ListQuery) (total int, err error){
	return srv.Repo.GetTotal(req)
}
func (srv *OrderService) Get(order *model.Order) (*model.Order, error){
	return srv.Repo.Get(order)
}
func (srv *OrderService) Exist(order *model.Order) bool{
	return srv.Repo.Exist(order)
}
func (srv *OrderService) Add(order *model.Order) (*model.Order, error){
	return srv.Repo.Add(order)
}
func (srv *OrderService) Edit(order model.Order) (bool, error){
	return srv.Repo.Edit(order)
}
func (srv *OrderService) Delete(id string) (bool, error){
	return srv.Repo.Delete(id)
}
