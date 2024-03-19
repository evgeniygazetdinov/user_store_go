package reports

import (
	"database/sql"
	"fmt"
	"log"
)

func DoConnection() *sql.DB {

	db, err := sql.Open("mysql", "docker:password@tcp(0.0.0.0:3306)/golang")

	if err != nil {
		log.Println(err)
	}
	database := db
	return database
}

func IndexReportService() []Report {
	database := DoConnection() // REFACTOR ON BASE METHOD input query and Model and Model Params
	rows, err := database.Query("select * from golang.reports")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
	reports := []Report{}
	for rows.Next() {
		r := Report{}
		err := rows.Scan(&r.Id, &r.Name, &r.Rows, &r.Columns)
		if err != nil {
			fmt.Println(err)

		}
		reports = append(reports, r)
	}
	return reports
}
