package repository

import (
	"fmt"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/query"
	"github.com/i-coder-robot/gin-demo/utils"
	"github.com/jinzhu/gorm"
)

type BannerRepository struct {
	DB *gorm.DB
}

type BannerRepoInterface interface {
	List(req query.ListQuery) (banners []*model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(banner *model.Banner) (*model.Banner, error)
	Exist(banner *model.Banner) bool
	Add(banner *model.Banner) (*model.Banner, error)
	Edit(banner model.Banner) (bool, error)
	Delete(id string) (bool, error)
}

func (repo *BannerRepository) List(req query.ListQuery) (banners []*model.Banner, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := utils.Page(req.Limit, req.Page) // 分页
	sort := utils.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}

func (repo *BannerRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var banners []model.Banner
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&banners).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *BannerRepository) Get(banner *model.Banner) (*model.Banner, error) {
	if err := repo.DB.Where(&banner).Find(&banner).Error; err != nil {
		return nil, err
	}
	return banner, nil
}

func (repo *BannerRepository) Exist(banner *model.Banner) bool {
	var count int
	if banner.Url != "" {
		repo.DB.Model(&banner).Where("url= ?", banner.Url).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo *BannerRepository) Add(banner *model.Banner) (*model.Banner, error) {
	if exist := repo.Exist(banner); exist == true {
		return banner, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(banner).Error
	if err != nil {

		return banner, fmt.Errorf("用户注册失败")
	}
	return banner, nil
}

func (repo *BannerRepository) Edit(banner model.Banner) (bool, error) {
	if banner.BannerID == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	id := &model.Banner{
		BannerID: banner.BannerID,
	}
	err := repo.DB.Model(id).Update(banner).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *BannerRepository) Delete(id string) (bool, error) {
	temp := &model.Banner{BannerID: id}
	err := repo.DB.Delete(temp).Error
	if err != nil {
		return false, err
	}
	return true, nil
}