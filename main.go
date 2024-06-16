package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var user struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": user.Name})
	})

	// 写一个简单的两数相加的接口
	r.GET("/add", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")
		c.JSON(http.StatusOK, gin.H{
			"result": a + b,
		})
	})

	r.Run(":8080")
}
