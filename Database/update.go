package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("faltal", err.Error())
	}
	stmt, err := db.Prepare("update todos set status=$2 where id =$1;")
	if err != nil {
		log.Fatal("prepare eror", err.Error())
	}
	if _, err := stmt.Exec(1, "active"); err != nil {
		log.Fatal("exec error", err.Error())
	}
	defer db.Close()
}
