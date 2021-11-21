package mysql

import (
	"database/sql"
	"log"

	"github.com/henprasetya/omdbapi/pkg/lib/mysql"
)

type OmdbMysql interface {
	SelectFromDb()
}

type omdbmysql struct {
	db *sql.DB
}

func NewQueryMysql(db *mysql.Mysql) OmdbMysql {
	return &omdbmysql{
		db: db.DB,
	}
}

func (q *omdbmysql) SelectFromDb() {
	log.Print("Select From DB")
}
