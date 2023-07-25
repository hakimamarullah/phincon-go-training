package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func main() {
	dbUser := "root"
	dbPass := "root"
	dbName := "phincon"
	dbHost := "127.0.0.1"
	dbPort := "3307"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO person(first_name, last_name) VALUES(?,?)", "Hakim"+strconv.FormatInt(time.Now().UnixNano(), 10), "Amarullah")
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var firstName string
		var lastName string
		rows.Scan(&firstName, &lastName)
		fmt.Println(firstName, lastName)
	}

}
