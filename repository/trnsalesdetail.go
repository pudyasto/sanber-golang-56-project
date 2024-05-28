package repository

import (
	"database/sql"
	"sanber-golang-56-paw/structs"
)

func GetAllTrnSalesDetail(db *sql.DB) (results []structs.TrnSalesDetail, err error) {
	sql := "SELECT id, trn_sales_id, item_id, qty, subtotal, total FROM trn_sales_detail"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var trn_sales_detail = structs.TrnSalesDetail{}

		err = rows.Scan(&trn_sales_detail.Id, &trn_sales_detail.TrnSalesId, &trn_sales_detail.ItemId, &trn_sales_detail.Qty, &trn_sales_detail.Subtotal, &trn_sales_detail.Total)
		if err != nil {
			panic(err)
		}
		results = append(results, trn_sales_detail)
	}

	return
}

func InsertTrnSalesDetail(db *sql.DB, trn_sales_detail structs.TrnSalesDetail) (err error) {
	sql := "INSERT INTO public.trn_sales_detail(trn_sales_id, item_id, qty, subtotal, total) VALUES ($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, trn_sales_detail.TrnSalesId, trn_sales_detail.ItemId, trn_sales_detail.Qty, trn_sales_detail.Subtotal, trn_sales_detail.Total)
	return errs.Err()
}

func UpdateTrnSalesDetail(db *sql.DB, trn_sales_detail structs.TrnSalesDetail) (err error) {
	sql := "UPDATE public.trn_sales_detail SET trn_sales_id=$1, item_id=$2, qty=$3, subtotal=$4, total=$5 WHERE id=$6"
	errs := db.QueryRow(sql, trn_sales_detail.TrnSalesId, trn_sales_detail.ItemId, trn_sales_detail.Qty, trn_sales_detail.Subtotal, trn_sales_detail.Total, trn_sales_detail.Id)
	return errs.Err()
}

func DeleteTrnSalesDetail(db *sql.DB, trn_sales_detail structs.TrnSalesDetail) (err error) {
	sql := "DELETE FROM public.trn_sales_detail WHERE id=$1"
	errs := db.QueryRow(sql, trn_sales_detail.Id)
	return errs.Err()
}
