//package with commands
package commands

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Statistic
func Statistic(database string) string {

	var rowsNum string

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select ROW_NUMBER() OVER(ORDER BY Id) FROM main.books ORDER BY Id DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&rowsNum)
		if err != nil {
			log.Fatal(err)
		}
	}
	textOfResp := "Books: " + rowsNum

	return textOfResp
}
