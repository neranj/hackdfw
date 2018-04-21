package models



import (
	"database/sql"
	_ "github.com/lib/pq"
)
var connection *sql.DB

func GetDb() *sql.DB {
	if connection == nil {
		var err error
		connection, err = sql.Open("postgres","user=﻿postgres password=﻿admin dbname=hackdfw sslmode=disable")
		if err != nil {
			panic(`DB connection could not be established`)
		}
	}
	return connection
}