package db

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	dsn := "landmark:landmark@csmsu@tcp(202.28.34.197:3306)/landmark"
	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)

	log.Println("Database connected")
}
