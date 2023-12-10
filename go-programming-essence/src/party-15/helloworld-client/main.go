package main

import (
	"context"
	"time"

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

	srv := micro.NewService()

	srv.Init()

	// Create client
	c := pb.NewHelloworldService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.Call(context.Background(), &pb.CallRequest{Name: "John"})
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
