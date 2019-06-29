package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	fmt.Println("url: ", url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("faltal", err.Error())
	}
	title := "HomeWork"
	status := "Inactive"
	row := db.QueryRow("insert into todos (title ,status) values ($1,$2) returning id", title, status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't scan id", id)
	}
	fmt.Println("result id : ", id)
	defer db.Close()
}
