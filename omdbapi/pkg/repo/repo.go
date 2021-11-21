package repo

import (
	"log"

	client "github.com/henprasetya/omdbapi/pkg/lib/http"
	"github.com/henprasetya/omdbapi/pkg/lib/mysql"
	repomysql "github.com/henprasetya/omdbapi/pkg/repo/mysql"
	"github.com/henprasetya/omdbapi/pkg/repo/omdb"
)

type Repository struct {
	MovieRepo omdb.MovieRepo
}

func CreateRepository() *Repository {
	log.Print("create repo")
	var mysql = mysql.NewMySql()
	var repo = repomysql.NewQueryMysql(mysql)
	return &Repository{
		MovieRepo: omdb.NewMovieRepo(client.NewDefaultHttpClient(), repo),
	}
}
