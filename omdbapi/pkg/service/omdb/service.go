package omdb

import (
	"github.com/henprasetya/omdbapi/pkg/model"
	reposql "github.com/henprasetya/omdbapi/pkg/repo/mysql"
	repo "github.com/henprasetya/omdbapi/pkg/repo/omdb"
)

type Service interface {
	SearchMovie(s string, page string) (*model.Response, error)
	MovieDetail(omdbID string) (*model.Movie, error)
}

type service struct {
	api repo.MovieRepo
	sql reposql.OmdbMysql
}

func NewMovieService(api repo.MovieRepo, sql reposql.OmdbMysql) Service {

	svc := &service{
		api: api,
		sql: sql,
	}
	return svc
}

func (s *service) SearchMovie(search string, page string) (*model.Response, error) {
	s.sql.SelectFromDb()
	movie, err := s.api.SearchMovie(search, page)
	if err != nil {
		return nil, err
	}

	return &*movie, err
}

func (s *service) MovieDetail(omdbID string) (*model.Movie, error) {
	s.sql.SelectFromDb()
	movie, err := s.api.MovieDetail(omdbID)
	if err != nil {
		return nil, err
	}

	return &*movie, err
}
