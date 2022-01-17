package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func parmsHandle(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}

func nameHandle(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "the string is :%s", name)
}

func main() {
	router := gin.Default()
	router.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Gin!")
	})
	v1 := router.Group("/api/v1")
	{
		v1.GET("/welcome", parmsHandle)
		v1.GET("/welcome/:name", nameHandle)
	}
	router.Run(":8000")
}
