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

func GetAllTrnSales(c *gin.Context) {
	var (
		result gin.H
	)

	trnsales, err := repository.GetAllTrnSales(database.DbConnection)

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
			"message": "Berhasil mengambil seluruh data transaksi sales",
			"data":    trnsales,
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetFormattedSales(c *gin.Context) {
	var (
		result gin.H
	)

	formattedSales, err := repository.GetFormattedSales(database.DbConnection)

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
			"message": "Berhasil mengambil seluruh data transaksi sales",
			"data":    formattedSales,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertTrnSales(c *gin.Context) {

	var (
		result gin.H
	)
	var trnsales structs.TrnSales

	err := c.ShouldBindJSON(&trnsales)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	// Generate code canvasser
	code := generateCodeTrnSal(database.DbConnection)
	fmt.Println(code)

	trnsales.Code = code

	err = repository.InsertTrnSales(database.DbConnection, trnsales)
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
		"message": "Berhasil menambahkan data transaksi sales",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func UpdateTrnSales(c *gin.Context) {
	var (
		result gin.H
	)
	var trnsales structs.TrnSales

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&trnsales)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	trnsales.Id = int64(id)
	err = repository.UpdateTrnSales(database.DbConnection, trnsales)
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
		"message": "Berhasil mengubah data transaksi sales",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func DeleteTrnSales(c *gin.Context) {
	var (
		result gin.H
	)
	var trnsales structs.TrnSales

	id, _ := strconv.Atoi(c.Param("id"))
	trnsales.Id = int64(id)

	err := repository.DeleteTrnSales(database.DbConnection, trnsales)
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
		"message": "Berhasil menghapus data transaksi sales",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func generateCodeTrnSal(db *sql.DB) string {
	var number int
	var prefix = "TSL"

	query := `SELECT MAX(substr(code,4,10))::int as number FROM trn_sales`
	rows, err := db.Query(query)
	if err != nil {
		return "TSL00001"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&number)
	}
	number++
	code := fmt.Sprintf("%s%05d", prefix, number)
	return code
}
