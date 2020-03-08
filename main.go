package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


// TODO download Sqlite3 driver for JetBrain to navigate using Database panel.

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}

	queryString := `SELECT * FROM customers;`
	_, err = db.Exec(queryString)
	if err != nil {
		panic (err)
	}


	fmt.Print("Hello World")
}