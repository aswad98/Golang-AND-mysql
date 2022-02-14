package controller

import (
	"database/sql"
	"fmt"
)

type DBConnect struct {
	conn *sql.DB
}

func DbConnect() *DBConnect {
	DB, err := sql.Open("mysql", "database:Aswad_database@123@tcp(127.0.0.1:3306)/calculatorDB")
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connected")
	}

	return &DBConnect{
		conn: DB,
	}

}
