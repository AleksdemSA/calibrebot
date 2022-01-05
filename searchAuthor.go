package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// SearchAuthor
func SearchAuthor(database string, query string) string {

	type Author struct {
		id         int
		name       string
		authorSort string
	}

	var id int
	var name string
	var author_sort string

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select id,title,author_sort from main.books where author_sort LIKE '%" + query + "%'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var authors []Author

	for rows.Next() {
		err := rows.Scan(&id, &name, &author_sort)
		if err != nil {
			log.Fatal(err)
		}
		authors = append(authors, Author{id, name, "\n" + author_sort + "\n------------------------\n"})
	}
	str := strings.Replace(fmt.Sprint(authors), "} {", "\n/", -1)
	str = strings.Replace(str, "[{", "Founded:\n/", -1)
	str = strings.Replace(str, "}]", "\nPress to number for list author's bookd", -1)
	str = strings.Replace(str, "[]", "Not found.", -1)
	if len(str) > 4096 {
		str = "Answer too big, messenger can's sent it"
	}
	return str
}
