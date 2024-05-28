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
			"code":   500,
			"result": err,
		}
	} else {
		result = gin.H{
			"code":   200,
			"result": trnsales,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetFormattedSales(c *gin.Context) {
	var (
		result gin.H
	)

	stock, err := repository.GetFormattedSales(database.DbConnection)

	if err != nil {
		result = gin.H{
			"code":   500,
			"result": err,
		}
	} else {
		result = gin.H{
			"code":   200,
			"result": stock,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertTrnSales(c *gin.Context) {
	var trnsales structs.TrnSales

	err := c.ShouldBindJSON(&trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Generate code canvasser
	code := generateCodeTrnSal(database.DbConnection)
	fmt.Println(code)

	trnsales.Code = code

	err = repository.InsertTrnSales(database.DbConnection, trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Insert TrnSales",
	})
}

func UpdateTrnSales(c *gin.Context) {
	var trnsales structs.TrnSales

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	trnsales.Id = int64(id)
	err = repository.UpdateTrnSales(database.DbConnection, trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Update TrnSales",
	})
}

func DeleteTrnSales(c *gin.Context) {
	var trnsales structs.TrnSales

	id, _ := strconv.Atoi(c.Param("id"))
	trnsales.Id = int64(id)

	err := repository.DeleteTrnSales(database.DbConnection, trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Delete TrnSales",
	})
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
