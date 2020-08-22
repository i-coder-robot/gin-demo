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
	List(req *query.ListQuery) (Categorys []*model.Category, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Category model.Category) (*model.Category, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(Category model.Category) (*model.Category, error)
	Edit(Category model.Category) (bool, error)
	Delete(c model.Category) (bool, error)
}

func (repo *CategoryRepository) List(req *query.ListQuery) (categories []*model.Category, err error) {
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

func (repo *CategoryRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
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

func (repo *CategoryRepository) Get(category model.Category) (*model.Category, error) {
	err := repo.DB.Where(&category).Find(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *CategoryRepository) Exist(category model.Category) *model.Category {
	if category.Name != "" {
		repo.DB.Model(&category).Where("name= ?", category.Name)
		return &category
	}
	return nil
}

func (repo *CategoryRepository) ExistByCategoryID(id string) *model.Category {
	var c model.Category
	repo.DB.Where("category_id = ?", id).First(&c)
	return &c

}

func (repo *CategoryRepository) Add(category model.Category) (*model.Category, error) {
	exist := repo.Exist(category);
	if exist !=nil {
		return nil, fmt.Errorf("分类已存在")
	}
	err := repo.DB.Create(category).Error
	if err != nil {

		return nil, fmt.Errorf("用户注册失败")
	}
	return &category, nil
}

func (repo *CategoryRepository) Edit(category model.Category) (bool, error) {
	if category.CategoryID == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &model.Category{
		CategoryID: category.CategoryID,
	}
	err := repo.DB.Model(id).Update(category).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *CategoryRepository) Delete(c model.Category) (bool, error) {
	temp := &model.Category{CategoryID: c.CategoryID}
	err := repo.DB.Delete(temp).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
