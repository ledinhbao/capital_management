package main

import "time"

type NGCTransaction struct {
	Id         uint
	CustomerId uint
	Date       time.Time
	Amount     int
	Price      float32
	Type       int8
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

var schemaNGCtransactionCreateTable = `
CREATE TABLE IF NOT EXISTS ngctransactions (
    id          INTEGER PRIMARY KEY,
    customer_id INTEGER,
    date        DATE,
    amount      INTEGER,
    price       DOUBLE,
    created_at  DATETIME DEFAULT (DATETIME('now')),
    updated_at  DATETIME DEFAULT (DATETIME('now'))
);
`
