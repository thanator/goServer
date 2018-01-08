package db

import (
	"database/sql"
	"fmt"
	"../consts"
)

func CreateConnection(query string) (*sql.Rows, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", consts.DB_USER, consts.DB_PASSWORD, consts.DB_NAME)
	dbase, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, err
	}

	err = dbase.Ping()
	if err != nil {
		fmt.Println("Ping error, %s", err)
		return nil, err
	} else {
		defer dbase.Close()
		return dbase.Query(query)
	}
}
