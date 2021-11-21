package omdb

import (
	"github.com/henprasetya/omdbapi/pkg/model"
	repo "github.com/henprasetya/omdbapi/pkg/repo/omdb"
)

type Service interface {
	SearchMovie(s string, page string) (*model.Response, error)
	MovieDetail(omdbID string) (*model.Movie, error)
}

type service struct {
	api repo.MovieRepo
}

func NewMovieService(api repo.MovieRepo) Service {
	var svc Service
	svc = &service{
		api: api,
	}
	return svc
}

func (s *service) SearchMovie(search string, page string) (*model.Response, error) {

	movie, err := s.api.SearchMovie(search, page)
	if err != nil {
		return nil, err
	}

	return &*movie, err
}

func (s *service) MovieDetail(omdbID string) (*model.Movie, error) {

	movie, err := s.api.MovieDetail(omdbID)
	if err != nil {
		return nil, err
	}

	return &*movie, err
}
