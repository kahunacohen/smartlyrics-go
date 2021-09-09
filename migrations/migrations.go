package migrations

import (
	"fmt"
	"time"

	"github.com/kahunacohen/smartlyrics/db"
	"github.com/kahunacohen/smartlyrics/models/album"
)

func Run() {
	con := db.Get()
	con.AutoMigrate(album.Album{})
	con.Create(&album.Album{Title: "Abbey Road", Artist: "The Beatles", CreatedAt: time.Now(), UpdatedAt: time.Now()})
	var a album.Album
	con.First(&a, 1) // find product with integer primary key
	fmt.Printf("%v", a.Artist)
}
