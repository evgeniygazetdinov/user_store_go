package sales

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

func MainService() []Products {
	database := DoConnection()
	rows, err := database.Query("select * from golang.products")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
	products := []Products{}
	for rows.Next() {
		p := Products{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)

		}
		products = append(products, p)
	}
	return products
}

func DeleteProductService(idOFProduct string) {
	database := DoConnection()
	_, err := database.Exec("delete from golang.products where id = ?", idOFProduct)
	if err != nil {
		log.Println(err)
	}
}

func AddProductService(model string, company string, price string) {
	database := DoConnection()
	database.Exec("insert into golang.products(model, company, price) values (?, ?,?)", model, company, price)
}

func EditProductService(id string) (Products, error) {
	database := DoConnection()
	products := Products{}
	row := database.QueryRow("select * from golang.products where id = ?", id)
	err := row.Scan(&products.Id, &products.Model, &products.Company, &products.Price)
	return products, err
}

func EditProductSpecitcService(id string, model string, company string, price string) error {
	database := DoConnection()
	_, err := database.Exec("update golang.products set model = ?, company = ?, price = ? where id = ?", model, company, price, id)
	return err
}
