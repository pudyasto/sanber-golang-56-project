package controllers

import (
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

func InsertTrnSales(c *gin.Context) {
	var trnsales structs.TrnSales

	err := c.ShouldBindJSON(&trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

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

	err := c.ShouldBindJSON(&trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	trnsales.Id = int64(id)

	err = repository.DeleteTrnSales(database.DbConnection, trnsales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Delete TrnSales",
	})
}
