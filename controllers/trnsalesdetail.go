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

func GetAllTrnSalesDetail(c *gin.Context) {
	var (
		result gin.H
	)

	trnsalesdetail, err := repository.GetAllTrnSalesDetail(database.DbConnection)

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
			"message": "Berhasil mengambil seluruh data transaksi sales detail",
			"data":    trnsalesdetail,
		}

		c.JSON(http.StatusOK, result)
	}
}

func InsertTrnSalesDetail(c *gin.Context) {
	var (
		result gin.H
	)
	var trnsalesdetail structs.TrnSalesDetail

	err := c.ShouldBindJSON(&trnsalesdetail)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	currStock := checkStockMin(database.DbConnection, trnsalesdetail.TrnSalesId, trnsalesdetail.ItemId)

	if currStock < trnsalesdetail.Qty {
		result = gin.H{
			"success": true,
			"message": "Stock tidak mencukupi!",
			"data":    []string{},
		}
		c.JSON(http.StatusForbidden, result)
		return
	}

	trnsalesdetail.Subtotal = float64(trnsalesdetail.Qty) * (trnsalesdetail.Price)

	err = repository.InsertTrnSalesDetail(database.DbConnection, trnsalesdetail)
	if err != nil {
		result = gin.H{
			"success": true,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	updatedQty := currStock - trnsalesdetail.Qty

	updateStock(database.DbConnection, trnsalesdetail.TrnSalesId, trnsalesdetail.ItemId, updatedQty)

	result = gin.H{
		"success": true,
		"message": "Berhasil menambahkan data transaksi sales detail",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func UpdateTrnSalesDetail(c *gin.Context) {
	var (
		result gin.H
	)
	var trnsalesdetail structs.TrnSalesDetail

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&trnsalesdetail)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Internal Server Error",
			"data":    []string{},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	trnsalesdetail.Id = int64(id)

	trnsalesdetail.Subtotal = float64(trnsalesdetail.Qty) * (trnsalesdetail.Price)

	err = repository.UpdateTrnSalesDetail(database.DbConnection, trnsalesdetail)
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
		"message": "Berhasil mengubah data transaksi sales detail",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func DeleteTrnSalesDetail(c *gin.Context) {
	var (
		result gin.H
	)
	var trnsalesdetail structs.TrnSalesDetail

	id, _ := strconv.Atoi(c.Param("id"))

	trnsalesdetail.Id = int64(id)

	restoreStock(database.DbConnection, int64(id))

	err := repository.DeleteTrnSalesDetail(database.DbConnection, trnsalesdetail)

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
		"message": "Berhasil menghapus data transaksi sales detail",
		"data":    []string{},
	}
	c.JSON(http.StatusOK, result)
}

func checkStockMin(db *sql.DB, trn_sales_id int64, item_id int64) int64 {
	// Get DataHeader
	var canvasserId int64

	queryHeader := `SELECT canvasser_id::int FROM trn_sales WHERE id=$1`
	rows, errHeader := db.Query(queryHeader, trn_sales_id)

	if errHeader != nil {
		return 0
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&canvasserId)
	}

	// Get Data Stock
	var qty int64

	query := `SELECT qty::int qty FROM stock WHERE item_id=$1 and canvasser_id=$2`

	rows, err := db.Query(query, item_id, canvasserId)

	if err != nil {
		return 0
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&qty)
	}

	return qty
}

func updateStock(db *sql.DB, trn_sales_id int64, item_id int64, updatedQty int64) (err error) {
	var canvasserId int64

	queryHeader := `SELECT canvasser_id::int FROM trn_sales WHERE id=$1`
	rows, errHeader := db.Query(queryHeader, trn_sales_id)

	if errHeader != nil {
		return errHeader
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&canvasserId)
	}

	query := `UPDATE stock SET qty=$1 WHERE item_id=$2 and canvasser_id=$3`
	errs := db.QueryRow(query, updatedQty, item_id, canvasserId)
	return errs.Err()
}

func restoreStock(db *sql.DB, trn_sales_detail_id int64) (err error) {

	var trnsalesdetail structs.TrnSalesDetail

	queryDetail := `SELECT trn_sales_id, item_id, qty FROM trn_sales_detail WHERE id=$1`
	rows, errDetail := db.Query(queryDetail, trn_sales_detail_id)

	if errDetail != nil {
		return errDetail
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&trnsalesdetail.TrnSalesId, &trnsalesdetail.ItemId, &trnsalesdetail.Qty)
	}

	var canvasserId int64

	queryHeader := `SELECT canvasser_id::int FROM trn_sales WHERE id=$1`
	rows, errHeader := db.Query(queryHeader, trnsalesdetail.TrnSalesId)

	if errHeader != nil {
		return errHeader
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&canvasserId)
	}

	resStock := trnsalesdetail.Qty + checkStockMin(db, trnsalesdetail.TrnSalesId, trnsalesdetail.ItemId)
	fmt.Println(checkStockMin(db, trnsalesdetail.TrnSalesId, trnsalesdetail.ItemId))
	fmt.Println(trnsalesdetail.Qty)
	// return

	query := `UPDATE stock SET qty=$1 WHERE item_id=$2 and canvasser_id=$3`
	errs := db.QueryRow(query, resStock, trnsalesdetail.ItemId, canvasserId)
	return errs.Err()
}
