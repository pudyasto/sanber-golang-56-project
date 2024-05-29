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

func GetAllCustomer(c *gin.Context) {

	var (
		result gin.H
	)
	customer, err := repository.GetAllCustomer(database.DbConnection)

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
			"message": "Berhasil mengambil seluruh data customer",
			"data":    customer,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertCustomer(c *gin.Context) {
	var (
		result gin.H
	)
	var customer structs.Customer

	err := c.ShouldBindJSON(&customer)
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
	code := generateCodeCust(database.DbConnection)
	customer.Code = code

	err = repository.InsertCustomer(database.DbConnection, customer)
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
		"message": "Berhasil menambahkan data customer",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func UpdateCustomer(c *gin.Context) {
	var (
		result gin.H
	)
	var customer structs.Customer

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	customer.Id = int64(id)

	err = repository.UpdateCustomer(database.DbConnection, customer)
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
		"message": "Berhasil mengubah data customer",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func DeleteCustomer(c *gin.Context) {
	var (
		result gin.H
	)
	var customer structs.Customer

	id, _ := strconv.Atoi(c.Param("id"))

	customer.Id = int64(id)

	err := repository.DeleteCustomer(database.DbConnection, customer)
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
		"message": "Berhasil menghapus data customer",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func generateCodeCust(db *sql.DB) string {
	var number int
	var prefix = "CST"

	query := `SELECT MAX(substr(code,4,10))::int as number FROM customer`
	rows, err := db.Query(query)
	if err != nil {
		return "CST00001"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&number)
	}
	number++
	code := fmt.Sprintf("%s%05d", prefix, number)
	return code
}
