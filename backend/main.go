package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies(nil)

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		fmt.Println("name:", name)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %s", name),
			"name":    name,
		})
	})

	fmt.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
