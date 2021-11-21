package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/henprasetya/omdbapi/pkg/lib/router"
	"github.com/henprasetya/omdbapi/pkg/service"
	moviedetail "github.com/henprasetya/omdbapi/pkg/service/omdb"
)

type httpServer struct {
	svc   *service.Services
	port  int
	errCh chan error
}

func NewHttpServer(svc *service.Services, port int) Server {
	return &httpServer{
		svc:  svc,
		port: port,
	}
}

func (h *httpServer) Run() {
	log.Printf("start running http server on port %d\n", h.port)

	route := router.NewDefaultRouter()

	moviedetail.RestHandler(h.svc.MovieService, route)
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), route.Handler())

	if err != nil {
		h.errCh <- err
	}
}

func (h *httpServer) ListenError() <-chan error {
	return h.errCh
}
