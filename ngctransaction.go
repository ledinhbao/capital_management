package main

import "time"

type NGCTransaction struct {
	Id uint
	CustomerId uint
	Date time.Time
	Amount int
	Price float32
	Type int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
