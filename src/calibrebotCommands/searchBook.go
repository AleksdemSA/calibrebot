package calibrebotCommands

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// SearchBook
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
	str = strings.Replace(str, "[{", "Founded:\n/", -1)
	str = strings.Replace(str, "}]", "\nPress to number for read description and download", -1)
	str = strings.Replace(str, "[]", "Not found.", -1)
	if len(str) > 4096 {
		str = "Answer too big, messenger can't sent it"
	}
	return str

}
