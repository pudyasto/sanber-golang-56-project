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
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"success": true,
			"message": "Berhasil mengambil seluruh data canvasser",
			"data":    canvasser,
		}

		c.JSON(http.StatusOK, result)
	}

}

func GetLoginCanvasser(c *gin.Context) {

	var (
		result gin.H
	)
	var canvasser structs.Canvasser

	err := c.ShouldBindJSON(&canvasser)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Data Username dan Password Wajib Di Isi",
			"data":    []string{},
		}
		c.JSON(http.StatusUnauthorized, result)
		return
	}

	username := canvasser.Username
	password := canvasser.Password

	dataCanvasser, err := repository.GetLoginCanvasser(database.DbConnection, username)

	if len(dataCanvasser) > 0 && checkPasswordHash(password, dataCanvasser[0].Password) {
		if err != nil {
			result = gin.H{
				"success": false,
				"message": "Internal Server Error",
				"data":    []string{},
			}
			c.JSON(http.StatusInternalServerError, result)
		} else {
			user := make(map[string]string)

			// Menambahkan elemen ke map
			user["id"] = strconv.Itoa(int(dataCanvasser[0].Id))
			user["nama"] = dataCanvasser[0].Name
			user["phone"] = dataCanvasser[0].Phone

			result = gin.H{
				"success": true,
				"message": "Login Berhasil",
				"data":    user,
				"token":   jwt.GenerateToken(),
			}
			c.JSON(http.StatusOK, result)
		}
	} else {
		result = gin.H{
			"success": false,
			"message": "Username atau Password anda Salah",
			"data":    []string{},
		}
		c.JSON(http.StatusUnauthorized, result)
	}
}

func InsertCanvasser(c *gin.Context) {
	var (
		result gin.H
	)
	var canvasser structs.Canvasser

	err := c.ShouldBindJSON(&canvasser)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	// Generate code
	code := generateCodeCnv(database.DbConnection)
	canvasser.Code = code

	// Hash Password bycript
	canvasser.Password = hashPassword(canvasser.Password)

	err = repository.InsertCanvasser(database.DbConnection, canvasser)
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
		"message": "Berhasil menambahkan data canvasser",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func UpdateCanvasser(c *gin.Context) {
	var (
		result gin.H
	)
	var canvasser structs.Canvasser

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&canvasser)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	canvasser.Id = int64(id)
	canvasser.Password = hashPassword(canvasser.Password)
	err = repository.UpdateCanvasser(database.DbConnection, canvasser)
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
		"message": "Berhasil mengubah data canvasser",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func DeleteCanvasser(c *gin.Context) {
	var (
		result gin.H
	)
	var canvasser structs.Canvasser

	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)

	canvasser.Id = int64(id)

	err := repository.DeleteCanvasser(database.DbConnection, canvasser)
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
		"message": "Berhasil menghapus data canvasser",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
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
