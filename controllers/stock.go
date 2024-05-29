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

func GetAllStock(c *gin.Context) {
	var (
		result gin.H
	)

	stock, err := repository.GetAllStock(database.DbConnection)

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
			"message": "Berhasil mengambil seluruh data stock",
			"data":    stock,
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetFormattedStock(c *gin.Context) {
	var (
		result gin.H
	)

	stock, err := repository.GetFormattedStock(database.DbConnection)

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
			"message": "Berhasil mengambil seluruh data stock",
			"data":    stock,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertStock(c *gin.Context) {
	var (
		result gin.H
	)
	var stock structs.Stock
	var msg string

	err := c.ShouldBindJSON(&stock)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	existStock := checkExistStock(database.DbConnection, stock.ItemId, stock.CanvasserId)
	fmt.Sprintln(existStock)

	// Kondisi jika stok sudah ada maka akan di update data qty nya
	if existStock {
		err = repository.UpdateStock(database.DbConnection, stock)
		msg = "Success Update Stock"
	} else {
		err = repository.InsertStock(database.DbConnection, stock)
		msg = "Success Insert Stock"
	}

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
		"message": msg,
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func UpdateStock(c *gin.Context) {
	var (
		result gin.H
	)
	var stock structs.Stock

	item_id, _ := strconv.Atoi(c.Param("item_id"))
	canvasser_id, _ := strconv.Atoi(c.Param("canvasser_id"))

	err := c.ShouldBindJSON(&stock)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	stock.ItemId = int64(item_id)
	stock.CanvasserId = int64(canvasser_id)

	err = repository.UpdateStock(database.DbConnection, stock)
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
		"message": "Berhasil mengubah data stock",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func DeleteStock(c *gin.Context) {
	var (
		result gin.H
	)
	var stock structs.Stock

	item_id, _ := strconv.Atoi(c.Param("item_id"))
	canvasser_id, _ := strconv.Atoi(c.Param("canvasser_id"))

	stock.ItemId = int64(item_id)
	stock.CanvasserId = int64(canvasser_id)

	err := repository.DeleteStock(database.DbConnection, stock)
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
		"message": "Berhasil menghapus data stock",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func checkExistStock(db *sql.DB, item_id int64, canvasser_id int64) bool {

	var number int

	query := `SELECT id::int FROM stock WHERE item_id=$1 AND canvasser_id=$2`
	rows, err := db.Query(query, item_id, canvasser_id)
	if err != nil {
		return false
	}
	for rows.Next() {
		rows.Scan(&number)
	}

	defer rows.Close()

	if number > 0 {
		return true
	} else {
		return false
	}
}
