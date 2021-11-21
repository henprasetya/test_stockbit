package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/henprasetya/omdbapi/pkg/repo"
	"github.com/henprasetya/omdbapi/pkg/server"
	"github.com/henprasetya/omdbapi/pkg/service"
)

func main() {

	repo := repo.CreateRepository()
	service := service.CreateServices(repo)

	httpServer := server.NewHttpServer(service, 8080)
	grpcServer := server.NewGrpcServer(service, 9090)

	/*
	* Goroutine
	 */
	runtime.GOMAXPROCS(10) //jumlah core yang digunakan
	go httpServer.Run()
	go grpcServer.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case o := <-term:
		log.Printf("exiting gracefully %s", o.String())
	case er := <-httpServer.ListenError():
		log.Printf("error starting http server, exiting gracefully %s", er.Error())
	case er := <-grpcServer.ListenError():
		log.Printf("error starting grpc server, exiting gracefully %s", er.Error())
	}

}
