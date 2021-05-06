package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username  string = "jkc"
	password  string = "jkc"
	ipaddress string = "192.168.134.14"
	database  string = "test"
)

var (
	dsn = fmt.Sprintf("%v:%v@%v/%v", username, password, ipaddress, database)
)

func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
