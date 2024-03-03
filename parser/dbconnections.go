package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql(username string, password string, host string, port int, DBName string) (*sql.DB, error) {
	var db *sql.DB
	dsn := "root:19990902Aa@@tcp(127.0.0.1:3306)/ETHERX"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
