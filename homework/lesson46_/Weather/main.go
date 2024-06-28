package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Go11Group/Javokhir-A/homework/lesson46_/Weather/genproto/weather"
	"github.com/Go11Group/Javokhir-A/homework/lesson46_/Weather/service"
	"github.com/Go11Group/Javokhir-A/homework/lesson46_/Weather/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db := postgres.NewDatabaseConnection()
	defer db.Close()

	weatherService := service.NewWeatherService(db)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	weather.RegisterWeatherServiceServer(s, weatherService)

	go func() {
		log.Println("Starting gRPC server on localhost:50054")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	log.Println("Stopping gRPC server...")
	s.GracefulStop()
	log.Println("Server stopped gracefully")

}
