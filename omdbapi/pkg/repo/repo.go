package repo

import (
	"log"

	client "github.com/henprasetya/omdbapi/pkg/lib/http"
	"github.com/henprasetya/omdbapi/pkg/repo/omdb"
)

type Repository struct {
	MovieRepo omdb.MovieRepo
}

func CreateRepository() *Repository {
	log.Print("create repo")
	return &Repository{
		MovieRepo: omdb.NewMovieRepo(client.NewDefaultHttpClient()),
	}
}
