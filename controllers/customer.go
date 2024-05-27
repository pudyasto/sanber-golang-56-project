package controllers

import (
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
			"code":   500,
			"result": err,
		}
	} else {
		result = gin.H{
			"code":   200,
			"result": customer,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCustomer(c *gin.Context) {
	var customer structs.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = repository.InsertCustomer(database.DbConnection, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Insert Customer",
	})
}

func UpdateCustomer(c *gin.Context) {
	var customer structs.Customer

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	customer.Id = int64(id)

	err = repository.UpdateCustomer(database.DbConnection, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Update Customer",
	})
}

func DeleteCustomer(c *gin.Context) {
	var customer structs.Customer

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	customer.Id = int64(id)

	err = repository.DeleteCustomer(database.DbConnection, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Delete Customer",
	})
}
