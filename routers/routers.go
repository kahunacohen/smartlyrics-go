package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kahunacohen/smartlyrics/handlers/album"
)

var router *gin.Engine = nil

func Get() *gin.Engine {
	if router == nil {
		router = gin.Default()
	}
	return router
}
func InitRoutes() {
	router := Get()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/albums", album.GetAll)
		v1.GET("/albums/:id", album.Get)
	}
}
