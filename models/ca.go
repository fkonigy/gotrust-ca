package models

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

type Ca struct {
	ID   int
	Name string
}

func DB() (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", "user=farid password=farid database=gotrust-ca sslmode=disable")
	checkErr(err, "sql.Open failed.")

	dbmap := &gorp.DbMap{
		Db: db, Dialect: gorp.PostgresDialect{}}

	// define tables
	dbmap.AddTableWithName(Ca{}, "cas").SetKeys(true, "ID")

	// create tables
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "dbmap.CreateTablesIfNotExists() failed.")

	return dbmap, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
