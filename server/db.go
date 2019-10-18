package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type userTable struct {
	id int64
}

type bookTable struct {
	id int64
}

var db *sql.DB

func init() {
	var err error
	user := os.Getenv("PSQLUSER")
	dbName := os.Getenv("PSQLDB")
	pass := os.Getenv("PSQLPASS")
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbName, pass))
	if err != nil {
		log.Fatal(err)
	}
}

func (ut userTable) existUserData() bool {
	var id int64
	stmt := fmt.Sprintf("select userid from usertable where userid = %d", ut.id)
	err := db.QueryRow(stmt).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

func (bt bookTable) existBookData() bool {
	var id int64
	stmt := fmt.Sprintf("select bookid from booktable where bookid = %d", bt.id)
	err := db.QueryRow(stmt).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

func (bt bookTable) insertBookData() {
	stmt := fmt.Sprintf("insert into booktable (bookid) values (%d)", bt.id)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func insertReceiptData(bt bookTable, ut userTable) {
	stmt := fmt.Sprintf("insert into receipt values (%d, %d);", ut.id, bt.id)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteReceiptData(bt bookTable, ut userTable) {
	stmt := fmt.Sprintf("delete from receipt where userid = (%d) and bookid = (%d);", ut.id, bt.id)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}
