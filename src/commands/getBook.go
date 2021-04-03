//package with commands
package commands

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// GetBook возвращает путь к книге
func GetBook(database string, query string) string {

	var path1 string
	var path2 string
	var path3 string

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// получаем папку автора книги
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

	// получаем путь к файлу без формата
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

	// получаем формат книги
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

	// объединяем всё это
	return path1 + "/" + path2 + "." + strings.ToLower(path3)

}
