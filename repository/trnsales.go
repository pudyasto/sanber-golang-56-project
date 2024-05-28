package repository

import (
	"database/sql"
	"sanber-golang-56-paw/structs"
)

func GetAllTrnSales(db *sql.DB) (results []structs.TrnSales, err error) {
	sql := "SELECT id, customer_id, canvasser_id, code, date_sales, description, subtotal, discount, total FROM trn_sales"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var trn_sales = structs.TrnSales{}

		err = rows.Scan(&trn_sales.Id, &trn_sales.CustomerId, &trn_sales.CanvasserId, &trn_sales.Code, &trn_sales.DateSales, &trn_sales.Description, &trn_sales.SubTotal, &trn_sales.Discount, &trn_sales.Total)
		if err != nil {
			panic(err)
		}
		results = append(results, trn_sales)
	}

	return
}

func InsertTrnSales(db *sql.DB, trn_sales structs.TrnSales) (err error) {
	sql := "INSERT INTO public.trn_sales(customer_id, canvasser_id, code, date_sales, description, subtotal, discount, total) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	errs := db.QueryRow(sql, trn_sales.CustomerId, trn_sales.CanvasserId, trn_sales.Code, trn_sales.DateSales, trn_sales.Description, trn_sales.SubTotal, trn_sales.Discount, trn_sales.Total)
	return errs.Err()
}

func UpdateTrnSales(db *sql.DB, trn_sales structs.TrnSales) (err error) {
	sql := "UPDATE public.trn_sales SET customer_id=$1, canvasser_id=$2, date_sales=$3, description=$4, subtotal=$5, discount=$6, total=$7 WHERE id=$8"
	errs := db.QueryRow(sql, trn_sales.CustomerId, trn_sales.CanvasserId, trn_sales.DateSales, trn_sales.Description, trn_sales.SubTotal, trn_sales.Discount, trn_sales.Total, trn_sales.Id)
	return errs.Err()
}

func DeleteTrnSales(db *sql.DB, trn_sales structs.TrnSales) (err error) {
	sql := "DELETE FROM public.trn_sales WHERE id=$1"
	errs := db.QueryRow(sql, trn_sales.Id)
	return errs.Err()
}
