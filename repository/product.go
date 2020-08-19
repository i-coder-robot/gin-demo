package repository

import (
	"fmt"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/utils"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	List(req query.ListQuery) (products []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(product *model.Product) (*model.Product, error)
	Exist(product *model.Product) bool
	Add(product *model.Product) (*model.Product, error)
	Edit(product model.Product) (bool, error)
	Delete(id string) (bool, error)
}


func (repo *ProductRepository) List(req query.ListQuery) (products []*model.Product, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page) // 分页
	sort := utils.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var products []*model.Product
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&products).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *ProductRepository) Get(product *model.Product) (*model.Product, error) {
	if err := repo.DB.Where(&product).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *ProductRepository) Exist(product *model.Product) bool {
	var count int
	if product.ProductName != "" {
		repo.DB.Model(&product).Where("product_name= ?", product.ProductName).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo *ProductRepository) Add(product *model.Product) (*model.Product, error) {
	if exist := repo.Exist(product); exist == true {
		return product, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(product).Error
	if err != nil {

		return product, fmt.Errorf("用户注册失败")
	}
	return product, nil
}

func (repo *ProductRepository) Edit(product model.Product) (bool, error) {
	if product.ProductId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &model.Product{
		ProductId: product.ProductId,
	}
	err := repo.DB.Model(id).Update(product).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *ProductRepository) Delete(id string) (bool, error) {
	temp := &model.Product{ProductId: id}
	err := repo.DB.Delete(temp).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
