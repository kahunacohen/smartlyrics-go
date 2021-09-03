package main

import "github.com/gin-gonic/gin"
import "net/http"

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"arist"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/artists", getAlbums)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
