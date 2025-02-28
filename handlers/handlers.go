package handlers

import (
	"go-sql/db"
	"go-sql/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ดึงข้อมูลสถานที่ทั้งหมด
func GetAllcountrys(c *gin.Context) {
	var countrys []models.Country
	query := "SELECT * FROM country"
	err := db.DB.Select(&countrys, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, countrys)
}

// ดึงข้อมูลตาม ID
func GetcountryById(c *gin.Context) {
	id := c.Param("id")
	var country models.Country
	query := "SELECT * FROM country WHERE idx = ?"
	err := db.DB.Get(&country, query, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "country not found"})
		return
	}
	c.JSON(http.StatusOK, country)
}

// เพิ่มข้อมูล
func Addcountry(c *gin.Context) {
	var country models.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO country (name) VALUES (?)"
	result, err := db.DB.Exec(query, country.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	country.Idx = int(id)
	c.JSON(http.StatusCreated, country)
}

// อัปเดตข้อมูล
func Updatecountry(c *gin.Context) {
	id := c.Param("id")
	var country models.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE country SET name = ? WHERE idx = ?"
	_, err := db.DB.Exec(query, country.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "country updated successfully"})
}

// ลบข้อมูล
func Deletecountry(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM country WHERE idx = ?"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "country deleted successfully"})
}
