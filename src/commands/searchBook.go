//package with commands
package commands

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// SearchBook выводит список книг по поиску
// К сожалению, пока не реализован регистронезависимый поиск
func SearchBook(database string, query string) string {

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

	// запрос к базе данных с поиском книг
	rows, err := db.Query("select id,title from main.books where title LIKE '%" + query + "%'")
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
	str2 := strings.Replace(str1, "}]", "\nНажми на номер для получения описания и скачивания", -1)
	str3 := strings.Replace(str2, "[]", "По данному выражению ничего не найдено.", -1)
	return str3
}
