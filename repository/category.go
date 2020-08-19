package repository

import (
	"fmt"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/utils"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	List(req query.ListQuery) (categories []*model.Category, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(category *model.Category) (*model.Category, error)
	Exist(category *model.Category) bool
	Add(category *model.Category) (*model.Category, error)
	Edit(category model.Category) (bool, error)
	Delete(id string) (bool, error)
}

func (repo *CategoryRepository) List(req query.ListQuery) (categories []*model.Category, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page) // 分页
	sort := utils.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (repo *CategoryRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var categories []model.Category
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&categories).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *CategoryRepository) Get(category *model.Category) (*model.Category, error) {
	if err := repo.DB.Where(&category).Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *CategoryRepository) Exist(category *model.Category) bool {
	var count int
	if category.Name != "" {
		repo.DB.Model(&category).Where("name= ?", category.Name).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo *CategoryRepository) Add(category *model.Category) (*model.Category, error) {
	if exist := repo.Exist(category); exist == true {
		return category, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(category).Error
	if err != nil {

		return category, fmt.Errorf("用户注册失败")
	}
	return category, nil
}

func (repo *CategoryRepository) Edit(category model.Category) (bool, error) {
	if category.ID == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &model.Category{
		ID: category.ID,
	}
	err := repo.DB.Model(id).Update(category).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *CategoryRepository) Delete(id string) (bool, error) {
	temp := &model.Category{ID: id}
	err := repo.DB.Delete(temp).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
