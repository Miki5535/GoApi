// package main

// import (
// 	"go-sql/db"
// 	"go-sql/handlers"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	db.InitDB()

// 	r := gin.Default()

// 	// กำหนดเส้นทาง API
// 	r.GET("/country", handlers.GetAllcountrys)
// 	r.GET("/country/:id", handlers.GetcountryById)
// 	r.POST("/country", handlers.Addcountry)
// 	r.PUT("/country/:id", handlers.Updatecountry)
// 	r.DELETE("/country/:id", handlers.Deletecountry)

// 	r.Run(":8080") // API จะรันที่ http://localhost:8080
// }

package handler

import (
	"go-sql/db"
	"go-sql/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db.InitDB()
	router := gin.Default()

	// กำหนดเส้นทาง API
	router.GET("/country", handlers.GetAllcountrys)
	router.GET("/country/:id", handlers.GetcountryById)
	router.POST("/country", handlers.Addcountry)
	router.PUT("/country/:id", handlers.Updatecountry)
	router.DELETE("/country/:id", handlers.Deletecountry)

	// ใช้ `ServeHTTP()` เพื่อให้ Vercel ใช้ Gin เป็น Handler
	router.ServeHTTP(w, r)
}
