package omdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	client "github.com/henprasetya/omdbapi/pkg/lib/http"
	"github.com/henprasetya/omdbapi/pkg/model"
	"github.com/henprasetya/omdbapi/pkg/repo/mysql"
	"github.com/joho/godotenv"
)

type MovieRepo interface {
	SearchMovie(search string, page string) (*model.Response, error)
	MovieDetail(imdbID string) (*model.Movie, error)
}

type movieRepo struct {
	client  client.HttpClient
	baseUrl string
	apikey  string

	db mysql.OmdbMysql
}

func NewMovieRepo(client client.HttpClient, db mysql.OmdbMysql) MovieRepo {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Cant Load Properties")
	}
	baseUrl := os.Getenv("movie_uri")
	omdbKey := os.Getenv("movie_key")
	return &movieRepo{
		client:  client,
		baseUrl: baseUrl,
		apikey:  omdbKey,
		db:      db,
	}
}

func (m *movieRepo) SearchMovie(search string, page string) (*model.Response, error) {

	m.db.SelectFromDb()
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%s", m.baseUrl, m.apikey, search, page)
	log.Print(url)
	data, err := m.client.GET(url)
	if err != nil {
		return nil, err
	}

	var resp model.Response
	err = json.Unmarshal(data, &resp)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *movieRepo) MovieDetail(omdbID string) (*model.Movie, error) {
	m.db.SelectFromDb()
	url := fmt.Sprintf("%s/?apikey=%s&i=%s", m.baseUrl, m.apikey, omdbID)
	log.Print(url)
	data, err := m.client.GET(url)
	if err != nil {
		return nil, err
	}

	var resp model.Movie
	err = json.Unmarshal(data, &resp)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}
