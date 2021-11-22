package repo

import (
	"log"

	client "github.com/henprasetya/omdbapi/pkg/lib/http"
	"github.com/henprasetya/omdbapi/pkg/lib/mysql"
	reposql "github.com/henprasetya/omdbapi/pkg/repo/mysql"
	"github.com/henprasetya/omdbapi/pkg/repo/omdb"
)

type Repository struct {
	MovieRepo omdb.MovieRepo
	SqlRepo   reposql.OmdbMysql
}

func CreateRepository() *Repository {
	log.Print("create repo")
	var repo = mysql.NewMySql()
	return &Repository{
		MovieRepo: omdb.NewMovieRepo(client.NewDefaultHttpClient()),
		SqlRepo:   reposql.NewQueryMysql(repo),
	}
}
