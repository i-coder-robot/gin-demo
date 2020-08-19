package repository

import (
	"fmt"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/utils"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(user *model.User) (*model.User, error)
	Exist(user *model.User) bool
	Add(user *model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

func (repo *UserRepository) List(req query.ListQuery) (users []*model.User, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page) // 分页
	sort := utils.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var users []model.User
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *UserRepository) Get(user *model.User) (*model.User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Exist(user *model.User) bool {
	var count int
	if user.NickName != "" {
		repo.DB.Model(&user).Where("nike_name= ?", user.NickName).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo *UserRepository) Add(user *model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist == true {
		return user, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(user).Error
	if err != nil {
		return user, fmt.Errorf("用户注册失败")
	}
	return user, nil
}

func (repo *UserRepository) Edit(user model.User) (bool, error) {
	if user.UserId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &model.User{
		UserId: user.UserId,
	}
	err := repo.DB.Model(id).Update(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) Delete(id string) (bool, error) {
	temp := &model.User{UserId: id}
	err := repo.DB.Delete(temp).Error
	if err != nil {
		return false, err
	}
	return true, nil
}