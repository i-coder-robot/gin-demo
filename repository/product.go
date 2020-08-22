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
	List(req *query.ListQuery) (Products []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Product model.Product) (*model.Product, error)
	Exist(Product model.Product) *model.Product
	ExistByProductID(id string) *model.Product
	Add(Product model.Product) (*model.Product, error)
	Edit(Product model.Product) (bool, error)
	Delete(u model.Product) (bool, error)
}

func (repo *ProductRepository) List(req *query.ListQuery) (products []*model.Product, err error) {
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

func (repo *ProductRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
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

func (repo *ProductRepository) Get(product model.Product) (*model.Product, error) {
	if err := repo.DB.Where(&product).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) Exist(product model.Product) *model.Product {
	if product.ProductName != "" {
		repo.DB.Model(&product).Where("product_name= ?", product.ProductName)
		return &product
	}
	return nil
}

func (repo *ProductRepository) ExistByProductID(id string) *model.Product {
	var p model.Product
	repo.DB.Where("product_id = ?", id).First(&p)
	return &p
}

func (repo *ProductRepository) Add(product model.Product) (*model.Product, error) {
	exist := repo.Exist(product)
	if exist !=nil {
		return &product, fmt.Errorf("商品已存在")
	}
	err := repo.DB.Create(product).Error
	if err != nil {
		return nil, fmt.Errorf("商品添加失败")
	}
	return &product, nil
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

func (repo *ProductRepository) Delete(product model.Product) (bool, error) {

	err := repo.DB.Model(&product).Where("product_id = ?", product.ProductId).Update("is_deleted", product.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
