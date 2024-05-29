package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sanber-golang-56-paw/database"
	"sanber-golang-56-paw/repository"
	"sanber-golang-56-paw/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllItem(c *gin.Context) {
	var (
		result gin.H
	)

	item, err := repository.GetAllItem(database.DbConnection)

	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"success": true,
			"message": "Berhasil mengambil seluruh data item",
			"data":    item,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertItem(c *gin.Context) {
	var (
		result gin.H
	)
	var item structs.Item

	err := c.ShouldBindJSON(&item)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	// Generate code
	code := generateCodeItem(database.DbConnection)
	item.Code = code

	err = repository.InsertItem(database.DbConnection, item)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"success": true,
		"message": "Berhasil menambahkan data item",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func UpdateItem(c *gin.Context) {
	var (
		result gin.H
	)
	var item structs.Item

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&item)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	item.Id = int64(id)

	err = repository.UpdateItem(database.DbConnection, item)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"success": true,
		"message": "Berhasil mengubah data item",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func DeleteItem(c *gin.Context) {
	var (
		result gin.H
	)
	var item structs.Item

	id, _ := strconv.Atoi(c.Param("id"))

	item.Id = int64(id)

	err := repository.DeleteItem(database.DbConnection, item)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"success": true,
		"message": "Berhasil menghapus data item",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func generateCodeItem(db *sql.DB) string {
	var number int
	var prefix = "IT"

	query := `SELECT MAX(substr(code,3,10))::int as number FROM item`
	rows, err := db.Query(query)
	if err != nil {
		return "IT00001"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&number)
	}
	number++
	code := fmt.Sprintf("%s%05d", prefix, number)
	return code
}
