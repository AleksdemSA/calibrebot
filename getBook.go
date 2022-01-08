package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// GetBook return path to book
func GetBook(database string, query string, path string) string {

	var path1 string
	var path2 string
	var path3 string

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

	rows2, err := db.Query("select name from main.data where book=" + query)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows2.Next() {
			err := rows2.Scan(&path2)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	defer rows.Close()

	rows3, err := db.Query("select format from main.data where book=" + query)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows3.Next() {
			err := rows3.Scan(&path3)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	defer rows.Close()

	return path + path1 + "/" + path2 + "." + strings.ToLower(path3)

}
