package repository

import (
	"database/sql"
	"sanber-golang-56-paw/structs"
)

func GetAllItem(db *sql.DB) (results []structs.Item, err error) {
	sql := "SELECT id, code, name, price FROM item"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item = structs.Item{}

		err = rows.Scan(&item.Id, &item.Code, &item.Name, &item.Price)
		if err != nil {
			panic(err)
		}
		results = append(results, item)
	}

	return
}

func InsertItem(db *sql.DB, item structs.Item) (err error) {
	sql := "INSERT INTO public.item(code, name, price) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, item.Code, item.Name, item.Price)
	return errs.Err()
}

func UpdateItem(db *sql.DB, item structs.Item) (err error) {
	sql := "UPDATE public.item SET code=$1, name=$2, price=$3 WHERE id=$4"
	errs := db.QueryRow(sql, item.Code, item.Name, item.Price, item.Id)
	return errs.Err()
}

func DeleteItem(db *sql.DB, item structs.Item) (err error) {
	sql := "DELETE FROM public.item WHERE id=$1"
	errs := db.QueryRow(sql, item.Id)
	return errs.Err()
}
