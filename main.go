package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//db, err := sql.Open("sqlite3", "./database.db")
	db, err := sqlx.Connect("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}

	var customers []NGCCustomer
	err = db.Select(&customers, "SELECT * FROM customers")
	if err != nil {
		panic(err)
	}
	fmt.Println(customers)

	fmt.Print("Hello World")
}
