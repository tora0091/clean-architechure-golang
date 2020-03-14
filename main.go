package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/music-list")
		v1.GET("/artist-list")
		v1.GET("/company-list")
	}

	r.Run(":8080")
}
