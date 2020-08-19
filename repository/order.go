package repository

import (
	"fmt"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/utils"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	List(req query.ListQuery) (orders []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(order *model.Order) (*model.Order, error)
	Exist(order *model.Order) bool
	Add(order *model.Order) (*model.Order, error)
	Edit(order model.Order) (bool, error)
	Delete(id string) (bool, error)
}


func (repo *OrderRepository) List(req query.ListQuery) (order []*model.Order, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page) // 分页
	sort := utils.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *OrderRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var orders []model.Order
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&orders).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *OrderRepository) Get(order *model.Order) (*model.Order, error) {
	if err := repo.DB.Where(&order).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *OrderRepository) Exist(order *model.Order) bool {
	var count int
	if order.OrderId != "" {
		repo.DB.Model(&order).Where("order_id= ?", order.OrderId).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo *OrderRepository) Add(order *model.Order) (*model.Order, error) {
	if exist := repo.Exist(order); exist == true {
		return order, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(order).Error
	if err != nil {

		return order, fmt.Errorf("用户注册失败")
	}
	return order, nil
}

func (repo *OrderRepository) Edit(order model.Order) (bool, error) {
	if order.OrderId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &model.Order{
		OrderId: order.OrderId,
	}
	err := repo.DB.Model(id).Update(order).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *OrderRepository) Delete(id string) (bool, error) {
	temp := &model.Order{OrderId: id}
	err := repo.DB.Delete(temp).Error
	if err != nil {
		return false, err
	}
	return true, nil
}