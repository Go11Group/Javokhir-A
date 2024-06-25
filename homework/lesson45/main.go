package main

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson45/service"
	"google.golang.org/grpc"
)

func main() {
	t := new(service.TranslatorService)

	grpcServer := grpc.NewServer()

	grpcServer.RegisterService()
}
