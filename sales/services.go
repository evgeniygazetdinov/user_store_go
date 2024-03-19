package sales

import (
	"context"
	// "database/sql"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DoConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://michael:secret@localhost:27017/"),
	)

	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("mongodb disconnect error : %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("connection error :%v", err)

	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
	}
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
