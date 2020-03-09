package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func saveNGCCustomerToDatabase(db *sqlx.DB, entity NGCCustomer) {
	tx := db.MustBegin()
	_, b := tx.NamedExec(entity.UpdateStringForNamedQuery(), entity)
	if b != nil {
		fmt.Println(b)
	}
	_ = tx.Commit()
}

func main() {
	//db, err := sql.Open("sqlite3", "./database.db")
	db, err := sqlx.Connect("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}

	//sampleCustomer := NGCCustomer{
	//	Id:        6,
	//	Name:      "Le Thanh Tra",
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//}
	//saveNGCCustomerToDatabase(db, sampleCustomer)

	var customers []NGCCustomer
	err = db.Select(&customers, "SELECT * FROM customers")
	if err != nil {
		panic(err)
	}
	for _, item := range customers {
		fmt.Println(item)
	}

	fmt.Print("Hello World")
}
