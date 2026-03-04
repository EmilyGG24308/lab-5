package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "series.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name FROM series")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Conexión exitosa 🔥")
}

