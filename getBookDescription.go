package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// GetBookDescription return book description
func GetBookDescription(database string, query string) string {

	var description string = "description not found"
	var text, format, title string

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows0, err := db.Query("select title from main.books where id=" + query)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows0.Next() {
			err := rows0.Scan(&title)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	defer rows0.Close()

	rows1, err := db.Query("select text from main.comments where book=" + query)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows1.Next() {
			err := rows1.Scan(&text)
			if err != nil {
				log.Fatal(err)
			}
			if len(text) < 2 {
				description = "description not found"
			} else {
				description = text
			}
		}
	}
	defer rows1.Close()

	rows2, err := db.Query("select format from main.data where book=" + query)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows2.Next() {
			err := rows2.Scan(&format)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	defer rows2.Close()

	return title + "\n------------------------\n" + description + "\n------------------------\n Format: " + format
}
