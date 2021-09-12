package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _con *gorm.DB = nil

func Get() *gorm.DB {
	if _con == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			os.Getenv("SMARTLYRICS_POSTGRES_INTERNAL_HOST"),
			os.Getenv("SMARTLYRICS_POSTGRES_USER"),
			os.Getenv("SMARTLYRICS_POSTGRES_PASSWORD"),
			os.Getenv("SMARTLYRICS_POSTGRES_DATABASE"),
			os.Getenv("SMARTLYRICS_POSTGRES_PORT"),
			os.Getenv("SMARTLYRICS_POSTGRES_SSL_MODE"),
			os.Getenv("SMARTLYRICS_POSTGRES_TZ"),
		)
		con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		_con = con
	}
	return _con
}

// func Truncate() {
// 	con := Get("localhost")
// 	con.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&album.Album{})
// }
