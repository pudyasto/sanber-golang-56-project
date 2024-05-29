package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sanber-golang-56-paw/database"
	"sanber-golang-56-paw/jwt"
	"sanber-golang-56-paw/repository"
	"sanber-golang-56-paw/structs"

	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllCanvasser(c *gin.Context) {
	var (
		result gin.H
	)

	canvasser, err := repository.GetAllCanvasser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"code":   500,
			"result": err,
		}
	} else {
		result = gin.H{
			"code":   200,
			"result": canvasser,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetLoginCanvasser(c *gin.Context) {
	var canvasser structs.Canvasser

	err := c.ShouldBindJSON(&canvasser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data Username dan Password Wajib Di Isi"})
		return
	}

	var (
		result gin.H
	)

	username := canvasser.Username
	password := canvasser.Password

	dataCanvasser, err := repository.GetLoginCanvasser(database.DbConnection, username)

	if len(dataCanvasser) > 0 && checkPasswordHash(password, dataCanvasser[0].Password) {
		if err != nil {
			result = gin.H{
				"code":   500,
				"result": err,
			}
		} else {
			result = gin.H{
				"code":   200,
				"result": dataCanvasser[0],
				"token":  jwt.GenerateToken(),
			}
		}
		c.JSON(http.StatusOK, result)
	} else {
		result = gin.H{
			"code":   401,
			"result": "Username atau Password anda Salah",
		}
		c.JSON(http.StatusUnauthorized, result)
	}
}

func InsertCanvasser(c *gin.Context) {
	var canvasser structs.Canvasser

	err := c.ShouldBindJSON(&canvasser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Generate code
	code := generateCodeCnv(database.DbConnection)
	canvasser.Code = code

	// Hash Password bycript
	canvasser.Password = hashPassword(canvasser.Password)

	err = repository.InsertCanvasser(database.DbConnection, canvasser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Insert Canvasser",
	})
}

func UpdateCanvasser(c *gin.Context) {
	var canvasser structs.Canvasser

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&canvasser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	canvasser.Id = int64(id)
	canvasser.Password = hashPassword(canvasser.Password)
	err = repository.UpdateCanvasser(database.DbConnection, canvasser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Update Canvasser",
	})
}

func DeleteCanvasser(c *gin.Context) {
	var canvasser structs.Canvasser

	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)

	canvasser.Id = int64(id)

	err := repository.DeleteCanvasser(database.DbConnection, canvasser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": "Success Delete Canvasser",
	})
}

func generateCodeCnv(db *sql.DB) string {
	var number int
	var prefix = "SL"

	query := `SELECT MAX(substr(code,3,10))::int as number FROM canvasser`
	rows, err := db.Query(query)
	if err != nil {
		return "SL00001"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&number)
	}
	number++
	code := fmt.Sprintf("%s%05d", prefix, number)
	return code
}

// Fungsi untuk mengenkripsi password
func hashPassword(password string) string {
	// Menggunakan bcrypt untuk mengenkripsi password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// Fungsi untuk memverifikasi password
func checkPasswordHash(password, hash string) bool {
	// Membandingkan password dengan hash yang tersimpan
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
