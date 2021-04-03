//package with commands
package commands

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// SearchAuthor выводит список авторов. Опять же как и поиск книг, эта операция регистрозависима.
func SearchAuthor(database string, query string) string {

	type Author struct {
		id         int
		name       string
		authorSort string
	}

	var id int
	var name string
	var authorSort string

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// запрос к базе данных с поиском книг
	rows, err := db.Query("select id,title,authorSort from main.books where authorSort LIKE '%" + query + "%'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var authors []Author

	for rows.Next() {
		err := rows.Scan(&id, &name, &authorSort)
		if err != nil {
			log.Fatal(err)
		}
		authors = append(authors, Author{id, name, "\n" + authorSort + "\n------------------------\n"})
	}
	str := strings.Replace(fmt.Sprint(authors), "} {", "\n/", -1)
	str1 := strings.Replace(str, "[{", "Найдено:\n/", -1)
	str2 := strings.Replace(str1, "}]", "\nНажми на номер для получения списка книг автора", -1)
	str3 := strings.Replace(str2, "[]", "По данному выражению ничего не найдено.", -1)
	return str3
}
