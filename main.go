package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create gin router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", healthHandler)

	r.GET("/divide", divideHandler)
	r.POST("/submit", submitHandler)
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running!",
	})
}

func divideHandler(c *gin.Context) {
	// Get query params
	a, errA := strconv.Atoi(c.Query("a"))
	b, errB := strconv.Atoi(c.Query("b"))
	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can not convert to integer",
		})
		return
	}
	if b == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Division by zero is not allowed"})
		return
	}
	result := float64(a) / float64(b)

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("%.2f", result)})
}

func submitHandler(c *gin.Context) {

	request := new(User)

	//c.ShouldBindJSON()
	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind request",
		})
		return
	}
	c.JSON(http.StatusOK, request)
}

type User struct {
	Username string
	Password string
}
