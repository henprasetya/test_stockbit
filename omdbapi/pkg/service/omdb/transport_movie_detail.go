package omdb

import (
	"context"
	"net/http"

	"github.com/henprasetya/omdbapi/pkg/gen/proto/v1"
	"github.com/henprasetya/omdbapi/pkg/lib/router"
)

type requestDetail struct {
	omdbID string
}

func restDecodeRequestDetail(_ context.Context, r *http.Request) (interface{}, error) {
	id := r.URL.Query().Get("omdbID")
	return requestDetail{omdbID: id}, nil
}

func grpcDecodeRequestDetailMovie(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.MovieDetailRequest)
	return requestDetail{omdbID: req.OmdbId}, nil
}

type grpcTransportDetailMovie struct {
	proto.UnimplementedMovieDetailServiceServer
	handler router.Handler
}

func (g *grpcTransportDetailMovie) MovieDetail(ctx context.Context, request *proto.MovieDetailRequest) (*proto.MovieDetailResponse, error) {
	_, res, err := g.handler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*proto.MovieDetailResponse), nil
}
