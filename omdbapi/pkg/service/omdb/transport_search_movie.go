package omdb

import (
	"context"
	"net/http"

	"github.com/henprasetya/omdbapi/pkg/gen/proto/v1"
	"github.com/henprasetya/omdbapi/pkg/lib/router"
)

type request struct {
	search string
	page   string
}

func restDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	search := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	return request{search: search, page: page}, nil
}

func grpcDecodeRequestSearchMovie(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.SearchMovieRequest)
	return request{search: req.Search, page: req.Page}, nil
}

type grpcTransportSearchMovie struct {
	proto.UnimplementedSearchMovieServiceServer
	handler router.Handler
}

func (g *grpcTransportSearchMovie) SearchMovie(ctx context.Context, request *proto.SearchMovieRequest) (*proto.SearchMovieResponse, error) {
	_, res, err := g.handler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*proto.SearchMovieResponse), nil
}
