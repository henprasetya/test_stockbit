package omdb

import (
	"github.com/henprasetya/omdbapi/pkg/model"
	sqlrepo "github.com/henprasetya/omdbapi/pkg/repo/mysql"
	repo "github.com/henprasetya/omdbapi/pkg/repo/omdb"
)

type Service interface {
	SearchMovie(s string, page string) (*model.Response, error)
	MovieDetail(omdbID string) (*model.Movie, error)
}

type service struct {
	api     repo.MovieRepo
	sqlrepo sqlrepo.OmdbMysql
}

func NewMovieService(api repo.MovieRepo, sql sqlrepo.OmdbMysql) Service {

	svc := &service{
		api:     api,
		sqlrepo: sql,
	}
	return svc
}

func (s *service) SearchMovie(search string, page string) (*model.Response, error) {
	s.sqlrepo.SelectFromDb()
	movie, err := s.api.SearchMovie(search, page)
	if err != nil {
		return nil, err
	}

	return &*movie, err
}

func (s *service) MovieDetail(omdbID string) (*model.Movie, error) {
	s.sqlrepo.SelectFromDb()
	movie, err := s.api.MovieDetail(omdbID)
	if err != nil {
		return nil, err
	}

	return &*movie, err
}
