package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "migrations.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// ########################
	// DDL
	// stmt := "CREATE TABLE demo (name text, ts_created timestamp)"

	// _, err = db.Exec(stmt)
	// if err != nil {
	// 	panic(err)
	// }

	// #########################
	// DML
	// tx, err := db.Begin()
	// if err != nil {
	// 	panic(err)
	// }

	// stmt, err := tx.Prepare("INSERT INTO demo (name) VALUES (?)")
	// if err != nil {
	// 	panic(err)
	// }

	// defer stmt.Close()

	// _, err = stmt.Exec("foo")
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = stmt.Exec("bar")
	// if err != nil {
	// 	panic(err)
	// }

	// err = tx.Commit()
	// if err != nil {
	// 	panic(err)
	// }

	// #########################
	// FETCHING DATA

	stmt, err := db.Prepare("SELECT rowid, name FROM demo WHERE rowid = ?")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	var id int
	var name string

	err = stmt.QueryRow("1").Scan(
		&id,
		&name,
	)

	if err != nil {
		panic(err)
	}

	log.Println(id, name)
}
