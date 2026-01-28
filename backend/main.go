package main

import (
	"fmt"
	"net/http"
	router "todolist/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := router.New()

	r.POST("/post", func(c *gin.Context) {
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

	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %s", name),
			"name":    "vlad",
			"status":  http.StatusOK,
			"age":     21,
		})
	})

	fmt.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
