package main

import "time"

type NGCCustomer struct {
	Id        uint      `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// UpdateStringForNamedQuery return insert query string if NGCCustomer.id != 0
func (x *NGCCustomer) UpdateStringForNamedQuery() string {
	queryString := ""
	if x.Id != 0 {
		queryString = "UPDATE customers SET name = :name, updated_at = date(\"now\") WHERE id = :id"
	} else {
		queryString = "INSERT INTO customers (name, created_at, updated_at) " +
			"VALUES (:name, :created_at, :updated_at)"
	}
	return queryString
}
