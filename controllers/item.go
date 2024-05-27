package controllers

import (
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
			"code":   500,
			"result": err,
		}
	} else {
		result = gin.H{
			"code":   200,
			"result": item,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertItem(c *gin.Context) {
	var item structs.Item

	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = repository.InsertItem(database.DbConnection, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Insert Item",
	})
}

func UpdateItem(c *gin.Context) {
	var item structs.Item

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	item.Id = int64(id)

	err = repository.UpdateItem(database.DbConnection, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Update Item",
	})
}

func DeleteItem(c *gin.Context) {
	var item structs.Item

	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	item.Id = int64(id)

	err = repository.DeleteItem(database.DbConnection, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Delete Item",
	})
}
