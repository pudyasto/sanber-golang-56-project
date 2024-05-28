package controllers

import (
	"net/http"
	"sanber-golang-56-paw/database"
	"sanber-golang-56-paw/repository"
	"sanber-golang-56-paw/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTrnSalesDetail(c *gin.Context) {
	var (
		result gin.H
	)

	trnsalesdetail, err := repository.GetAllTrnSalesDetail(database.DbConnection)

	if err != nil {
		result = gin.H{
			"code":   500,
			"result": err,
		}
	} else {
		result = gin.H{
			"code":   200,
			"result": trnsalesdetail,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertTrnSalesDetail(c *gin.Context) {
	var trnsalesdetail structs.TrnSalesDetail

	err := c.ShouldBindJSON(&trnsalesdetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	trnsalesdetail.Subtotal = float64(trnsalesdetail.Qty) * (trnsalesdetail.Price)

	err = repository.InsertTrnSalesDetail(database.DbConnection, trnsalesdetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Insert TrnSalesDetail",
	})
}

func UpdateTrnSalesDetail(c *gin.Context) {
	var trnsalesdetail structs.TrnSalesDetail

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&trnsalesdetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	trnsalesdetail.Id = int64(id)
	err = repository.UpdateTrnSalesDetail(database.DbConnection, trnsalesdetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Update TrnSalesDetail",
	})
}

func DeleteTrnSalesDetail(c *gin.Context) {
	var trnsalesdetail structs.TrnSalesDetail

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&trnsalesdetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	trnsalesdetail.Id = int64(id)

	err = repository.DeleteTrnSalesDetail(database.DbConnection, trnsalesdetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Delete TrnSalesDetail",
	})
}
