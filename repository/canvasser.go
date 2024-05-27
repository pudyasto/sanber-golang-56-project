package repository

import (
	"database/sql"
	"fmt"
	"sanber-golang-56-paw/structs"
)

func GetAllCanvasser(db *sql.DB) (err error, results []structs.Canvasser) {
	sql := "SELECT * FROM canvasser"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var canvasser = structs.Canvasser{}

		err = rows.Scan(&canvasser.Id, &canvasser.Code, &canvasser.Name, &canvasser.Phone, &canvasser.Username, &canvasser.Password)
		if err != nil {
			panic(err)
		}
		results = append(results, canvasser)
	}

	return
}

func GetLoginCanvasser(db *sql.DB, username string) (results []structs.Canvasser, err error) {
	query := "SELECT * FROM canvasser WHERE username = $1"

	rows, err := db.Query(query, username)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var canvasser structs.Canvasser

		err = rows.Scan(&canvasser.Id, &canvasser.Code, &canvasser.Name, &canvasser.Phone, &canvasser.Username, &canvasser.Password)

		results = append(results, canvasser)
		fmt.Println(results)

	}
	return
}

func InsertCanvasser(db *sql.DB, canvasser structs.Canvasser) (err error) {
	sql := "INSERT INTO canvasser(code, name, phone, username, password) VALUES ($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, canvasser.Code, canvasser.Name, canvasser.Phone, canvasser.Username, canvasser.Password)
	return errs.Err()
}

func UpdateCanvasser(db *sql.DB, canvasser structs.Canvasser) (err error) {
	sql := "UPDATE public.canvasser SET code=$1, name=$2, phone=$3, username=$4, password=$5 WHERE id=$6"
	errs := db.QueryRow(sql, canvasser.Code, canvasser.Name, canvasser.Phone, canvasser.Username, canvasser.Password, canvasser.Id)
	return errs.Err()
}

func DeleteCanvasser(db *sql.DB, canvasser structs.Canvasser) (err error) {
	sql := "DELETE FROM public.canvasser WHERE id=$1"
	errs := db.QueryRow(sql, canvasser.Id)
	return errs.Err()
}
