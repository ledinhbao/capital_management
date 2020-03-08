package main

import "time"

type Customer struct {
	Id   uint
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}
