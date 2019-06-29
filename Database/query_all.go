package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Todos struct {
	ID     int
	Title  string
	Status string
}

var dataList = []Todos{}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	stmt, _ := db.Prepare("select id, title ,status from todos")
	rows, _ := stmt.Query()
	for rows.Next() {
		todo := Todos{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		if err != nil {
			log.Fatal("error", err.Error())
		}
		dataList = append(dataList, todo)
		fmt.Println("data :", todo)
	}
}
