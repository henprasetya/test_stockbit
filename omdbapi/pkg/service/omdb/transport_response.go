package omdb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/henprasetya/omdbapi/pkg/gen/proto/v1"
	"github.com/henprasetya/omdbapi/pkg/model"
)

func restEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func restEncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func grpcEncodeResponseSearchMovie(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*model.Response)
	var data []*proto.SearchMovie
	for _, s := range res.Search {
		data = append(data, &proto.SearchMovie{
			Title:  s.Title,
			ImdbId: s.ImdbID,
			Year:   s.Year,
			Type:   s.Type,
			Poster: s.Poster,
		})
	}
	return &proto.SearchMovieResponse{
		Search:       data,
		TotalResults: res.TotalResults,
		Response:     res.Response,
	}, nil
}

func grpcEncodeResponseMovieDetail(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*model.Movie)

	return &proto.MovieDetailResponse{
		Title:    res.Title,
		Year:     res.Year,
		Rated:    res.Rated,
		Released: res.Released,
		Runtime:  res.Runtime,
		Genre:    res.Genre,
		Director: res.Director,
		Writer:   res.Writer,
		Actors:   res.Actors,
		Plot:     res.Plot,
		Language: res.Language,
		Country:  res.Country,
		Awards:   res.Awards,
		Poster:   res.Poster,
		Response: res.Response,
	}, nil
}
