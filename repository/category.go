package repository

import (
	"fmt"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	List(req *query.ListQuery) (Categorys []*model.CategoryResult, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(Category model.Category) (*model.Category, error)
	Edit(Category model.Category) (bool, error)
	Delete(c model.Category) (bool, error)
}

func (repo *CategoryRepository) List(req *query.ListQuery) (categories []*model.CategoryResult, err error) {
	//fmt.Println(req)
	//db := repo.DB
	//limit, offset := utils.Page(req.Limit, req.Page) // 分页
	//sort := utils.Sort(req.Sort)                     // 排序 默认 created_at desc
	//if req.Where != "" {
	//	db = db.Where(req.Where)
	//}
	//
	//if err := db.Order(sort).Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
	//	return nil, err
	//}
	//return categories, nil
	var list []*model.CategoryResult
	err = repo.DB.Raw("SELECT c1.category_id as c1_category_id,c1.name as c1_name,c1.desc as c1_desc,c1.order as c1_order,c1.parent_id as c1_parent_id, c2.category_id as c2_category_id,c2.name as c2_name,c2.order as c2_order,c2.parent_id as c2_parent_id,c3.category_id as c3_category_id,c3.name as c3_name,c3.order as c3_order,c3.parent_id as c3_parent_id FROM category c1 join category c2 on c1.category_id = c2.parent_id join category c3 on c2.category_id=c3.parent_id").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *CategoryRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
	//var categories []model.Category
	//db := repo.DB
	//if req.Where != "" {
	//	db = db.Where(req.Where)
	//}
	//if err := db.Find(&categories).Count(&total).Error; err != nil {
	//	return total, err
	//}
	//return total, nil

	err = repo.DB.Raw("SELECT count(c3.category_id) FROM category c1 join category c2 on c1.category_id = c2.parent_id join category c3 on c2.category_id=c3.parent_id").Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *CategoryRepository) Get(id string) ([]*model.CategoryResult, error) {
	var list []*model.CategoryResult
	err := repo.DB.Raw("SELECT c1.category_id as c1_category_id,c1.name as c1_name,c1.desc as c1_desc,c1.order as c1_order,c2.category_id as c2_category_id,c2.name as c2_name,c2.order as c2_order,c3.category_id as c3_category_id,c3.name as c3_name,c3.order as c3_order FROM category c1 join category c2 on c1.category_id = c2.parent_id join category c3 on c2.category_id=c3.parent_id where c3.category_id = ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *CategoryRepository) Exist(category model.Category) *model.Category {
	var c model.Category
	if category.Name != "" {
		repo.DB.Model(&c).Where("name= ?", category.Name)
		return &c
	}
	return nil
}

func (repo *CategoryRepository) ExistByCategoryID(id string) *model.Category {
	var c model.Category
	repo.DB.Where("category_id = ?", id).First(&c)
	return &c

}

func (repo *CategoryRepository) Add(category model.Category) (*model.Category, error) {
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
