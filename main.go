package main

type Customer struct {
	Id   uint   `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
