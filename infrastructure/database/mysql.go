package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var user = "root"
var pass = "root"
var url = "mysql"
var port = "3306"
var name = "crud"

func Connect() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", user, pass, url, port, name)

	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(-1)
	}

	err = conn.Ping()

	if err != nil {
		conn.Close()
		panic(-1)
	}

	fmt.Printf("Database %s:%s/%s opened\n", url, port, name)

	return conn
}
