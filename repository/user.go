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
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user *model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	ExistByMobile(mobile string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(u model.User) (bool, error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users []*model.User, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page) // 分页
	//sort := utils.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	if err := db.Order("id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
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

func (repo *UserRepository) Exist(user model.User) *model.User {
	var count int
	repo.DB.Find(&user).Where("nick_name = ?", user.NickName)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByMobile(mobile string) *model.User{
	var count int
	var user model.User
	repo.DB.Find(&user).Where("mobile = ?", mobile)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var count int
	var user model.User
	repo.DB.Find(&user).Where("user_id = ?", id).Count(&count)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist !=nil {
		return nil, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}

func (repo *UserRepository) Edit(user model.User) (bool, error) {

	id := &model.User{
		UserId: user.UserId,
	}
	err := repo.DB.Model(id).Update(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) Delete(u model.User) (bool, error) {
	err := repo.DB.Model(&u).Update("is_deleted",u.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
