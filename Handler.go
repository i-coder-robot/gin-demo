package main

import (
	"fmt"
	"github.com/i-coder-robot/gin-demo/handler"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/repository"
	"github.com/i-coder-robot/gin-demo/service"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var (
	DB              *gorm.DB
	BannerHandler   handler.BannerHandler
	CategoryHandler handler.CategoryHandler
	OrderHandler    handler.OrderHandler
	ProductHandler  handler.ProductHandler
	UserHandler     handler.UserHandler
)

func Init() {

	fmt.Println("数据库 init")
	var err error
	conf := &model.DBConf{
		Driver:   viper.GetString("DB_DRIVER"),
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DbName:   viper.GetString("DB_NAME"),
		Charset:  viper.GetString("DB_CHARSET"),
	}
	DB, err = gorm.Open("mysql", conf)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	fmt.Println("数据库 init 结束...")

	BannerHandler = handler.BannerHandler{
		BannerSrv: &service.BannerService{
			Repo: &repository.BannerRepository{
				DB: DB,
			},
		}}

	CategoryHandler = handler.CategoryHandler{
		CategorySrv: &service.CategoryService{
			Repo: &repository.CategoryRepository{
				DB: DB,
			},
		},
	}

	OrderHandler = handler.OrderHandler{
		OrderSrv: &service.OrderService{
			Repo: &repository.OrderRepository{
				DB: DB,
			},
		}}

	ProductHandler = handler.ProductHandler{
		ProductSrv: &service.ProductService{
			Repo: &repository.ProductRepository{
				DB: DB,
			},
		}}

	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: DB,
			},
		}}

}

//config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
//username,
//password,
//addr,
//name,
//true,
//"Local")
//
//db, err := gorm.Open("mysql", config)
