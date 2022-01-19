package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // 返回默认的路由引擎

	//// 指定用户 GET 请求访问 /book
	//r.GET("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Get Book",
	//	})
	//})
	//
	//// 指定用户 PUT 请求访问 /hello
	//r.PUT("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Put Book",
	//	})
	//})
	//
	//// 指定用户 POST 请求访问 /hello
	//r.POST("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "POST Book",
	//	})
	//})
	//
	//// 指定用户 DELETE 请求访问 /hello
	//r.DELETE("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "DELETE Book",
	//	})
	//})

	

	// 运行
	r.Run("127.0.0.1:9090")
}
