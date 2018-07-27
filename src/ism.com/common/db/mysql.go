package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func GetDatabase() *sql.DB {
	var err error
	if dbConn == nil {
		dbConn, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/go_ism")
		// if there is an error opening the connection, handle it
		if err != nil {
			panic(err.Error())
		}

		dbConn.Ping()
		dbConn.SetMaxIdleConns(5)
		dbConn.SetMaxOpenConns(100)
	}
	return dbConn
}
