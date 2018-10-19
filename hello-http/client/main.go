package main

import (
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"im-cs/logger"
	"fmt"
	pb "grpc-gateway/proto"
)

const (
	ip = "127.0.0.1"
	grpc_port = ":60059"
	restful_port = ":8888"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(ip + grpc_port, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewHelloHttpClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloHttpRequest{
		Name: "gRPC-grpc is working!",
	})
	if err != nil {
		logger.Fatal(fmt.Sprintf("Create error: %v", err))
	}
	logger.Info(fmt.Sprintf("Create response: %s", r))
}
