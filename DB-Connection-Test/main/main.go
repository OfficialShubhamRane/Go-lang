package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/gorm")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	//createProductTable(db)

	//product := Product{"pen", 2.2, "black pen", true}
	//insertProductTable(db, product)

	//newProduct := getProductById(db, 1) // Assuming 123 is the product ID
	//if err != nil {
	//	// Handle the error (e.g., log, return to the user)
	//} else if newProduct == nil {
	//	// Handle the case where the product was not found
	//} else {
	//	// Use the product data
	//	fmt.Println("Product Name:", newProduct.Name)
	//	// ...
	//}

	//getProduct(db)
}

func getProduct(db *sql.DB) *[]Product {

	query := `SELECT name, price, description, is_available FROM gorm.product`

	var name string
	var price float64
	var desc string
	var isAvailable bool
	var productSlice []Product

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&name, &price, &desc, &isAvailable)
		if err != nil {
			log.Fatal(err)
		}
		productSlice = append(productSlice, Product{name, price, desc, isAvailable})
	}
	fmt.Println(productSlice)
	return &productSlice

}

type Product struct {
	Name        string
	Price       float64
	Desc        string
	IsAvailable bool
}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE gorm.product (
			  seq INT NOT NULL AUTO_INCREMENT,
			  name VARCHAR(45) NULL,
			  price INT NULL,
			  description VARCHAR(45) NULL,
			  is_available TINYINT NULL,
			  createdStamp TIMESTAMP NULL DEFAULT now(),
			  PRIMARY KEY (seq))`

	_, err := db.Exec(query)
	if err != nil {
		return
	}

}

func insertProductTable(db *sql.DB, product Product) {
	query := `INSERT INTO gorm.product(
			name, price, description, is_available)
			VALUES
			(?, ?, ?, ?);`

	err := db.QueryRow(query, product.Name, product.Price, product.Desc, product.IsAvailable)
	if err != nil {
		return
	}

}

func getProductById(db *sql.DB, seq int) *Product {
	query := `SELECT name, price, description, is_available FROM gorm.product where seq = ?`

	var name string
	var price float64
	var desc string
	var isAvailable bool

	err := db.QueryRow(query, seq).Scan(&name, &price, &desc, &isAvailable)
	if err != nil {
		log.Fatal(err)
	}
	return &Product{name, price, desc, isAvailable}
}
