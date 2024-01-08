package main

import (
	"context"
	"errors"
	"fmt"
	service "goTest/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net"
)

// hello server
type server struct {
	service.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "zhangziyuan" || appKey != "123123" {
		return nil, errors.New("token 不正确")
	}

	return &service.HelloResponse{ResponseMsg: "你好" + req.RequestName}, nil
}

func main() {
	// TLS认证
	// 两个参数分别是 cretFile，keyFile
	// 自签名证书文件和私钥文件
	// creds, _ := credentials.NewServerTLSFromFile("D:\\go_project\\goTest\\grpcTest\\key\\test.pem",
	// 	"D:\\go_project\\goTest\\grpcTest\\key\\test.key")
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	// 不启动安全认证 grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// 在grpc服务端中去注册我们自己编写的服务
	service.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
