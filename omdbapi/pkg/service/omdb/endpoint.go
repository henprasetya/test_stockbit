package omdb

import (
	"context"

	"github.com/henprasetya/omdbapi/pkg/gen/proto/v1"
	"github.com/henprasetya/omdbapi/pkg/lib/router"
	"google.golang.org/grpc"
)

func endpoint(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(request)
		m, err := s.SearchMovie(r.search, r.page)
		return m, err
	}
}

func endpointDetail(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requestDetail)
		m, err := s.MovieDetail(r.omdbID)
		return m, err
	}
}

func RestHandler(svc Service, r router.Router) {
	handler := router.NewHandler(endpoint(svc), restDecodeRequest, restEncodeResponse, restEncodeError)
	handlerDetail := router.NewHandler(endpointDetail(svc), restDecodeRequestDetail, restEncodeResponse, restEncodeError)
	r.GET("/api/search_movie", handler)
	r.GET("/api/movie_detail", handlerDetail)
}

func GrpcHandler(svc Service, grpcServer *grpc.Server) {
	t := &grpcTransportSearchMovie{handler: router.NewGrpcHandler(endpoint(svc), grpcDecodeRequestSearchMovie, grpcEncodeResponseSearchMovie)}
	d := &grpcTransportDetailMovie{handler: router.NewGrpcHandler(endpointDetail(svc), grpcDecodeRequestDetailMovie, grpcEncodeResponseMovieDetail)}
	proto.RegisterSearchMovieServiceServer(grpcServer, t)
	proto.RegisterMovieDetailServiceServer(grpcServer, d)
}
