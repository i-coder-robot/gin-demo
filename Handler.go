package main

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/model"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func CategoryListHandler(c *gin.Context) {

	var firstCategories []*model.Category
	var secondCategories []*model.Category
	var thirdCategories []*model.Category

	c31 := &model.Category{
		ID:       uuid.NewV4().String(),
		Name:     "苹果",
		Desc:     "",
		Children: nil,
	}
	c32 := &model.Category{
		ID:       uuid.NewV4().String(),
		Name:     "橙子",
		Desc:     "",
		Children: nil,
	}
	thirdCategories = append(thirdCategories, c31)
	thirdCategories = append(thirdCategories, c32)
	c21 := &model.Category{
		ID:       uuid.NewV4().String(),
		Name:     "热销水果",
		Desc:     "",
		Children: thirdCategories,
	}
	secondCategories = append(secondCategories, c21)

	category := &model.Category{
		ID:       uuid.NewV4().String(),
		Name:     "新鲜水果",
		Desc:     "草莓 | 水蜜桃 | 车厘子",
		Order:    200,
		Children: secondCategories,
	}
	firstCategories = append(firstCategories, category)
	c.JSON(http.StatusOK, gin.H{"category_list": firstCategories})
}

func AddCategoryHandler(c *gin.Context)  {
	var category model.Category
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{})
	}



}

func EditCategoryHandler(c *gin.Context)  {

}

func DeleteCategoryHandler(c *gin.Context)  {

}