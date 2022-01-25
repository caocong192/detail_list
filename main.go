package main

import (
	"github.com/gin_detail_list/dao"
	"github.com/gin_detail_list/models"
	"github.com/gin_detail_list/routers"
)



func main() {
	// 创建数据库
	// sql: create Database bubble;
	// 连接数据库
	if err := dao.InitMySQL(); err!=nil{
		panic(err)
	}

	// 关闭连接
	defer dao.Close()

	// 绑定模型
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routers.SetupRouter()

	// 启动服务
	r.Run("127.0.0.1:18080")
}
