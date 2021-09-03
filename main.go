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
func getAlbum(c *gin.Context) {
	//id := c.Param("id")
	c.IndentedJSON(http.StatusOK, albums[0])
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/albums", getAlbums)
		v1.GET("/albums/:id", getAlbum)

	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
