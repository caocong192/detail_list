package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dns := "caocong:cc1764..@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dns)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 创建数据库
	// sql: create Database bubble;
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer DB.Close()

	// 绑定模型
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	// 加载模板文件引用的静态文件
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// V1
	v1Group := r.Group("v1")
	{
		// 代办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写代办事项, 会发送请求到这
			// 1. 从请求中拿到数据
			var todo Todo
			c.BindJSON(&todo)
			// 2. 存入数据库
			// 3. 返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
				//c.JSON(http.StatusOK, gin.H{
				//	"code": 200,
				//	"msg":  "success",
				//	"data": todo,
				//})
			}
		})

		// 查看
		v1Group.GET("/todo", func(c *gin.Context) {
			// 1. 从数据库获取清单数据
			// 2. 返回给前端
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
				return
			}

			var todo Todo
			if err = DB.Where("id=?", id).Find(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}

			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			// 1. 获取要删除数据的ID
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
				return
			}

			// 2. 删除数据
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	r.Run()
}
