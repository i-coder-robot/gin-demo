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
	List(req *query.ListQuery) (Banners []*model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Banner model.Banner) (*model.Banner, error)
	Exist(Banner model.Banner) *model.Banner
	ExistByBannerID(id string) *model.Banner
	Add(Banner model.Banner) (*model.Banner, error)
	Edit(Banner model.Banner) (bool, error)
	Delete(id string) (bool, error)
}

func (repo *BannerRepository) List(req *query.ListQuery) (banners []*model.Banner, err error) {
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

func (repo *BannerRepository) GetTotal(req *query.ListQuery) (total int64, err error) {
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

func (repo *BannerRepository) Get(banner model.Banner) (*model.Banner, error) {
	if err := repo.DB.Where(&banner).Find(&banner).Error; err != nil {
		return nil, err
	}
	return &banner, nil
}

func (repo *BannerRepository) Exist(banner model.Banner) *model.Banner {
	if banner.Url != "" {
		repo.DB.Model(&banner).Where("url= ?", banner.Url)
		return &banner
	}
	return nil
}

func (repo *BannerRepository) ExistByBannerID(id string) *model.Banner {
	var b model.Banner
	repo.DB.Where("order_id = ?", id).First(&b)
	return &b
}

func (repo *BannerRepository) Add(banner model.Banner) (*model.Banner, error) {
	exist := repo.Exist(banner)
	if exist != nil {
		return nil, fmt.Errorf("轮播已存在")
	}
	err := repo.DB.Create(banner).Error
	if err != nil {
		return nil, fmt.Errorf("轮播添加失败")
	}
	return &banner, nil
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
	banner := repo.ExistByBannerID(id)
	err := repo.DB.Delete(banner).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
