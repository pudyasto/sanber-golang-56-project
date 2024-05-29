package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	result := gin.H{
		"success": true,
		"message": "Selamat datang di project akhir sanbercode golang batch 56 - paw",
		"data":    []string{},
	}

	c.JSON(http.StatusOK, result)
}
