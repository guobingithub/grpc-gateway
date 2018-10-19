package main

import (
	pb "grpc-gateway/proto" // 引入proto包
	"moa-user-manager/logger"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:70059"
)

func main() {
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("C:/My/gopath/src/grpc-gateway/keys/server.pem", "server name")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloHttpClient(conn)

	// 调用方法
	reqBody := new(pb.HelloHttpRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)

	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Println(r.Message)
	logger.Error("client get msg, rsp:",r)
}

// OR: curl -X POST -k https://localhost:50052/example/echo -d '{"name": "gRPC-HTTP is working!"}'