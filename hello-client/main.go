package main

import (
	"context"
	"fmt"
	service "goTest/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "zhangziyuan",
		"appKey": "123123",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	// creds, _ := credentials.NewClientTLSFromFile("D:\\go_project\\goTest\\grpcTest\\key\\test.pem",
	// 	"*.zhaoziyuan.com")

	// 连接到server端，此处禁用安全传输，没有加密和验证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	// 连接到server端，此处禁用安全传输，没有加密和验证
	// 不开启安全传输代码 grpc.WithTransportCredentials(insecure.NewCredentials())
	// 开启安全传输代码 grpc.WithTransportCredentials(creds)
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("did not connet: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := service.NewSayHelloClient(conn)
	// 执行rpc调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &service.HelloRequest{RequestName: "关鑫鹏"})
	fmt.Println(resp.GetResponseMsg())
}
