package repository

import (
	"database/sql"
	"sanber-golang-56-paw/structs"
)

func GetAllStock(db *sql.DB) (results []structs.Stock, err error) {
	sql := "SELECT id, item_id, canvasser_id, qty FROM stock"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock = structs.Stock{}

		err = rows.Scan(&stock.Id, &stock.ItemId, &stock.CanvasserId, &stock.Qty)
		if err != nil {
			panic(err)
		}
		results = append(results, stock)
	}

	return
}

func InsertStock(db *sql.DB, stock structs.Stock) (err error) {
	sql := "INSERT INTO public.stock(item_id, canvasser_id, qty) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, stock.ItemId, stock.CanvasserId, stock.Qty)
	return errs.Err()
}

func UpdateStock(db *sql.DB, stock structs.Stock) (err error) {
	sql := "UPDATE public.stock SET qty=$1 WHERE item_id=$2 AND canvasser_id=$3"
	errs := db.QueryRow(sql, stock.Qty, stock.ItemId, stock.CanvasserId)
	return errs.Err()
}

func DeleteStock(db *sql.DB, stock structs.Stock) (err error) {
	sql := "DELETE FROM public.stock WHERE item_id=$1 AND canvasser_id=$2"
	errs := db.QueryRow(sql, stock.ItemId, stock.CanvasserId)
	return errs.Err()
}
