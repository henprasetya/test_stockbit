package mysql

import (
	"database/sql"
	"log"
)

type Mysql struct {
	DB *sql.DB
}

func NewMySql() *Mysql {
	s, err := dbConnection()
	if err != nil {
		panic(err.Error())
	}
	return s
}

func dbConnection() (*Mysql, error) {
	log.Print("Connect to DB")
	return &Mysql{}, nil
}
