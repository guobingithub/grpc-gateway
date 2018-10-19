package main

import (
	"golang.org/x/net/context"
	pb "grpc-gateway/proto"
	"net"
	"net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"im-cs/logger"
	"fmt"
)

const (
	ip = "127.0.0.1"
	grpc_port = ":60059"
	restful_port = ":8888"
)

// 定义helloHTTPService并实现约定的接口
type helloHTTPService struct{}

// HelloHTTPService Hello HTTP服务
var HelloHTTPService = helloHTTPService{}

// SayHello 实现Hello服务接口
func (h helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHttpRequest) (*pb.HelloHttpReply, error) {
	resp := new(pb.HelloHttpReply)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func StartGrpcServer() {
	//ip, err := util.GetLocalIp("以太网")
	//ip := "localhost"
	//grpc_port := ":50059"

	lis, err := net.Listen("tcp", ip + grpc_port)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterHelloHttpServer(s, &helloHTTPService{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logger.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}

func StartRESTFulServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//ip, err := util.GetLocalIp("以太网")
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHelloHttpHandlerFromEndpoint(ctx, mux, ip + grpc_port, opts)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to RegisterHelloHttpHandlerFromEndpoint: %v", err))
	}

	if err := http.ListenAndServe(ip + restful_port, mux); err != nil {
		logger.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}

func main() {
	go StartGrpcServer()
	StartRESTFulServer()
}
