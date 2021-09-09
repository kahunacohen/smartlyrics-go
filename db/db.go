package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=db user=smartlyrics password=Vp~sJT46 dbname=smartlyrics port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return con
}
