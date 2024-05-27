package repository

import (
	"database/sql"
	"sanber-golang-56-paw/structs"
)

func GetAllCustomer(db *sql.DB) (results []structs.Customer, err error) {
	sql := "SELECT id, code, name, address, email FROM customer"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var customer = structs.Customer{}

		err = rows.Scan(&customer.Id, &customer.Code, &customer.Name, &customer.Address, &customer.Email)
		if err != nil {
			panic(err)
		}
		results = append(results, customer)
	}

	return
}

func InsertCustomer(db *sql.DB, customer structs.Customer) (err error) {
	sql := "INSERT INTO customer(code, name, address, email) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, customer.Code, customer.Name, customer.Address, customer.Email)
	return errs.Err()
}

func UpdateCustomer(db *sql.DB, customer structs.Customer) (err error) {
	sql := "UPDATE public.customer SET code=$1, name=$2, address=$3, email=$4 WHERE id=$5"
	errs := db.QueryRow(sql, customer.Code, customer.Name, customer.Address, customer.Email, customer.Id)
	return errs.Err()
}

func DeleteCustomer(db *sql.DB, customer structs.Customer) (err error) {
	sql := "DELETE FROM public.customer WHERE id=$1"
	errs := db.QueryRow(sql, customer.Id)
	return errs.Err()
}
