package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin_detail_list/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	// 加载模板文件引用的静态文件
	r.Static("/static", "static")

	r.GET("/", controller.IndexHandler)

	// V1
	v1Group := r.Group("v1")
	{
		// 代办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看
		v1Group.GET("/todo", controller.GetTotoList)

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
