package album

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"arist"`
}

var as = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan"},
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, as)
}
func Get(c *gin.Context) {
	//id := c.Param("id")
	c.IndentedJSON(http.StatusOK, as[0])
}
