//package with commands
package commands

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// RandBook выводит случайную книгу из библиотеки
func RandBook(database string) string {

	type Book struct {
		id    int
		title string
	}

	var id int
	var title string

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// выбираем случайную книгу из таблицы
	rows, err := db.Query("SELECT id,title FROM main.books ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, Book{id, title})
	}
	str := strings.Replace(fmt.Sprint(books), "} {", "\n/", -1)
	str1 := strings.Replace(str, "[{", "Найдено:\n/", -1)
	str2 := strings.Replace(str1, "}]", "\n\nНажми на номер для получения описания и скачивания", -1)
	return str2
}
