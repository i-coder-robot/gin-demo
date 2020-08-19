package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func Cors() gin.HandlerFunc{
	return func(c *gin.Context) {
		method:=c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	//users:=[]User{{ID:123,Name:"张三丰"},{ID:456,Name:"张无忌"}}
	r:=gin.Default()
	r.Use(Cors())

	r.GET("/api/category_list",CategoryListHandler)
	r.POST("api/add_category",AddCategoryHandler)
	r.POST("api/edit_category",EditCategoryHandler)
	r.POST("api/delete_category",DeleteCategoryHandler)

	r.Run(":8889")

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"name":"小红",
	//		"desc":"漂亮的小姐姐，我来了，今天是520，你有人约会吗？",
	//	})
	//})
	////http://127.0.0.1:8889/123?name=老王
	//r.GET("/users/:id", func(context *gin.Context) {
	//	id:=context.Param("id")
	//	fmt.Println("我的Id 是"+id)
	//	name:=context.Query("name")
	//
	//	context.JSON(http.StatusOK,gin.H{
	//		"name":name,
	//		"desc":"小姐姐，我的 ID是"+id,
	//
	//	})
	//})
	//r.GET("/slice", func(context *gin.Context) {
	//	context.JSON(http.StatusOK,context.QueryArray("media"))
	//})
	//r.GET("/map", func(context *gin.Context) {
	//	context.JSON(http.StatusOK,context.QueryMap("ids"))
	//})
	//r.POST("/users", func(context *gin.Context) {
	//	//创建一个新用户
	//	var user model.User
	//	err := context.BindJSON(&user)
	//	if err != nil {
	//		panic(err)
	//	}
	//	context.JSON(http.StatusOK,gin.H{"user":user})
	//
	//})
	//r.GET("/users", func(context *gin.Context) {
	//	//获取全部用户
	//})
	//r.PUT("/user/:id", func(context *gin.Context) {
	//	//更新用户id 为 xx 的信息
	//})
	//r.DELETE("/user/:id", func(context *gin.Context) {
	//	//删除用户 id 为 xx
	//})
}