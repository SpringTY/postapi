package main

import (
	"postapi/handler"
	"postapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LogerMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/deals/:dataSource/:dealId", service.GetDealsHandler)
	r.Run("0.0.0.0:7675") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
