/*
A contrived SQL DB example.
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// DB Open
	db, err := sql.Open("sqlite3", "./foo.db")
	testError("DBOPEN", err)

	// Insert
	stmt, err := db.Prepare("INSERT INTO memos(text, priority) values(?,?)")
	testError("Prepare", err)
	res, err := stmt.Exec("memo data", 15)
	testError("Exec", err)
	fmt.Println(res)

	// Query
	rows, err := db.Query("SELECT * FROM memos")
	testError("Query", err)
	var text string
	var priority int
	// Scan all rows returned
	for rows.Next() {
		err = rows.Scan(&text, &priority)
		testError("Scan", err)
		fmt.Println(text)
		fmt.Println(priority)
	}
	rows.Close()

	// DB Close
	db.Close()

}

func testError(id string, err error) {
	if err != nil {
		panic(id + " " + err.Error())
	}
}
