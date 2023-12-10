package main

import (

	"helloworld/handler"
	pb "helloworld/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

)

var (
	service = "helloworld"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
