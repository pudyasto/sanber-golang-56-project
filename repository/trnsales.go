package repository

import (
	"database/sql"
	"sanber-golang-56-paw/structs"
)

func GetAllTrnSales(db *sql.DB) (results []structs.TrnSales, err error) {
	sql := "SELECT id, customer_id, canvasser_id, code, date_sales, description FROM trn_sales"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var trn_sales = structs.TrnSales{}

		err = rows.Scan(&trn_sales.Id, &trn_sales.CustomerId, &trn_sales.CanvasserId, &trn_sales.Code, &trn_sales.DateSales, &trn_sales.Description)
		if err != nil {
			panic(err)
		}
		results = append(results, trn_sales)
	}

	return
}

func InsertTrnSales(db *sql.DB, trn_sales structs.TrnSales) (err error) {
	sql := "INSERT INTO public.trn_sales(customer_id, canvasser_id, code, date_sales, description) VALUES ($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, trn_sales.CustomerId, trn_sales.CanvasserId, trn_sales.Code, trn_sales.DateSales, trn_sales.Description)
	return errs.Err()
}

func UpdateTrnSales(db *sql.DB, trn_sales structs.TrnSales) (err error) {
	sql := "UPDATE public.trn_sales SET customer_id=$1, canvasser_id=$2, date_sales=$3, description=$4 WHERE id=$8"
	errs := db.QueryRow(sql, trn_sales.CustomerId, trn_sales.CanvasserId, trn_sales.DateSales, trn_sales.Description, trn_sales.Id)
	return errs.Err()
}

func DeleteTrnSales(db *sql.DB, trn_sales structs.TrnSales) (err error) {
	sql := "DELETE FROM public.trn_sales WHERE id=$1"
	errs := db.QueryRow(sql, trn_sales.Id)
	return errs.Err()
}

func GetFormattedSales(db *sql.DB) (results []structs.ReportSales, err error) {
	sql := `SELECT trn_sales.code, 
				trn_sales.date_sales, 
				trn_sales.description,
				canvasser.code as canvasser_code,
				canvasser.name as canvasser_name,
				customer.code as customer_code,
				customer.name as customer_name,
				SUM(trn_sales_detail.subtotal) as total
			FROM public.trn_sales
			JOIN customer on customer.id = trn_sales.customer_id
			JOIN canvasser on canvasser.id = trn_sales.canvasser_id
			JOIN trn_sales_detail ON trn_sales_detail.trn_sales_id = trn_sales.id
			GROUP BY trn_sales.code, 
				trn_sales.date_sales, 
				trn_sales.description,
				canvasser.code,
				canvasser.name,
				customer.code,
				customer.name`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var rptSales = structs.ReportSales{}

		err = rows.Scan(&rptSales.Code, &rptSales.DateSales, &rptSales.Description, &rptSales.CanvasserCode, &rptSales.CanvasserName, &rptSales.CustomerCode, &rptSales.CustomerName, &rptSales.Total)
		if err != nil {
			panic(err)
		}
		results = append(results, rptSales)
	}

	return
}
