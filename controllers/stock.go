package controllers

import (
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

func InsertStock(c *gin.Context) {
	var stock structs.Stock

	err := c.ShouldBindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = repository.InsertStock(database.DbConnection, stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Insert Stock",
	})
}

func UpdateStock(c *gin.Context) {
	var stock structs.Stock

	item_id, _ := strconv.Atoi(c.Param("item_id"))
	canvasser_id, _ := strconv.Atoi(c.Param("canvasser_id"))

	err := c.ShouldBindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	stock.ItemId = int64(item_id)
	stock.CanvasserId = int64(canvasser_id)

	err = repository.UpdateStock(database.DbConnection, stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Update Stock",
	})
}

func DeleteStock(c *gin.Context) {
	var stock structs.Stock

	item_id, _ := strconv.Atoi(c.Param("item_id"))
	canvasser_id, _ := strconv.Atoi(c.Param("canvasser_id"))

	err := c.ShouldBindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	stock.ItemId = int64(item_id)
	stock.CanvasserId = int64(canvasser_id)

	err = repository.DeleteStock(database.DbConnection, stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Delete Stock",
	})
}
