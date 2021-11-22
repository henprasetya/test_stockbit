package service

import (
	"github.com/henprasetya/omdbapi/pkg/repo"
	service "github.com/henprasetya/omdbapi/pkg/service/omdb"
)

type Services struct {
	MovieService service.Service
}

func CreateServices(r *repo.Repository) *Services {

	m := service.NewMovieService(r.MovieRepo, r.SqlRepo)
	return &Services{
		MovieService: m,
	}
}
