package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin_detail_list/models"
	"net/http"
)

/*
	url --> controller --> logic/service --> models
	请求 --> 控制层    --> 业务逻辑 --> 模型层的CRUD
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写代办事项, 会发送请求到这
	// 1. 从请求中拿到数据
	var todo models.Todo
	c.BindJSON(&todo)
	// 2. 存入数据库
	// 3. 返回响应
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 200,
		//	"msg":  "success",
		//	"data": todo,
		//})
	}
}

func GetTotoList(c *gin.Context) {
	// 1. 从数据库获取清单数据
	// 2. 返回给前端
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
		return
	}

	todo, err := models.GetATodo(&id)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	// 1. 获取要删除数据的ID
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
		return
	}

	// 2. 删除数据
	if err := models.DeleteATodo(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}