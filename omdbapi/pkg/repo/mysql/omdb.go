package mysql

import (
	"database/sql"
	"log"

	"github.com/henprasetya/omdbapi/pkg/lib/mysql"
)

//go:generate mockgen -destination=../../mock/repo/mysql/mock_omdb.go -package=mock_repository github.com/henprasetya/omdbapi/pkg/repo/mysql OmdbMysql

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
