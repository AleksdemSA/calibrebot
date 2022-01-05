package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetImage(database string, query string) string {

	var path1 string

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select path from main.books where id=" + query)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			err := rows.Scan(&path1)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	defer rows.Close()

	return path1 + "/cover.jpg"

}
