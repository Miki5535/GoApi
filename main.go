package main

import (
	"go-sql/db"
	"go-sql/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	// กำหนดเส้นทาง API
	r.GET("/country", handlers.GetAllcountrys)
	r.GET("/country/:id", handlers.GetcountryById)
	r.POST("/country", handlers.Addcountry)
	r.PUT("/country/:id", handlers.Updatecountry)
	r.DELETE("/country/:id", handlers.Deletecountry)

	r.Run(":8080") // API จะรันที่ http://localhost:8080
}
