package main

import (
	"github.com/kahunacohen/smartlyrics/db"
	"github.com/kahunacohen/smartlyrics/routers"
)

func main() {
	r := routers.Get()
	routers.InitRoutes()
	db.Connect()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// 	type Product struct {
	//   gorm.Model
	//   Code  string
	//   Price uint
	// }
	// 	db.AutoMigrate(&Product{})
	// 	db.Create(&Product{Code: "D42", Price: 100})
}
