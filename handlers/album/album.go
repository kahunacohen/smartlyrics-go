package album

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kahunacohen/smartlyrics/db"
	"github.com/kahunacohen/smartlyrics/models/album"
)

var as = []album.Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane"},
	{ID: 2, Title: "Jerus", Artist: "Gerdy Mulligan"},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan"},
}

func GetAll(c *gin.Context) {
	var albums []album.Album
	result := db.Get().Find(&albums)
	fmt.Println("hey:")
	fmt.Printf("%v", result)

	c.IndentedJSON(http.StatusOK, as)
}
func Get(c *gin.Context) {
	//id := c.Param("id")
	c.IndentedJSON(http.StatusOK, as[0])
}
