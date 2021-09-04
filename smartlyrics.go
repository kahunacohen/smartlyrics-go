package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kahunacohen/smartlyrics/album"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/albums", album.GetAll)
		v1.GET("/albums/:id", album.Get)

	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
