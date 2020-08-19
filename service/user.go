package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/repository"
)

type UserSrv interface {
	List(req query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(user *model.User) (*model.User, error)
	Exist(user *model.User) bool
	Add(user *model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}

func (srv *UserService) List(req query.ListQuery) (users []*model.User, err error){
	return srv.Repo.List(req)
}
func (srv *UserService) GetTotal(req *query.ListQuery) (total int, err error){
	return srv.Repo.GetTotal(req)
}
func (srv *UserService) Get(user *model.User) (*model.User, error){
	return srv.Repo.Get(user)
}
func (srv *UserService) Exist(user *model.User) bool{
	return srv.Repo.Exist(user)
}
func (srv *UserService) Add(user *model.User) (*model.User, error){
	return srv.Repo.Add(user)
}
func (srv *UserService) Edit(user model.User) (bool, error){
	return srv.Repo.Edit(user)
}
func (srv *UserService) Delete(id string) (bool, error){
	return srv.Repo.Delete(id)
}
