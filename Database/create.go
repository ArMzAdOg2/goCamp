package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	const url = "postgres://ienwwjmt:UjeG5IsHwsTTg8jiHH2rimwo4sOFH_dX@satao.db.elephantsql.com:5432/ienwwjmt"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("faltal", err.Error())
	}
	createTb := `
	CREATE TABLE todos(
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create database", err.Error())
	}
	defer db.Close()
}
