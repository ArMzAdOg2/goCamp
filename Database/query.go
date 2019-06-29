package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int
	Title  string
	Status string
}

func main() {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	stmt, _ := db.Prepare("SELECT id, title, status from todos where id = $1;")
	row := stmt.QueryRow(1)

	var todo = Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		log.Fatal("error", err.Error())
	}
	fmt.Println("one row", todo)
}
