package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _db *gorm.DB = nil

func Get() *gorm.DB {
	if _db == nil {
		dsn := "host=db user=smartlyrics password=Vp~sJT46 dbname=smartlyrics port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		fmt.Println(dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		_db = db
		return db
	}
	return _db
}
