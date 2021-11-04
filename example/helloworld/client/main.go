package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/miaogaolin/gofirst/example/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 连接服务端
	// grpc.WithInsecure() 跳过安全验证，即明文传输
	// grpc.WithBlock() 等待与服务端握手成功
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 客户端实例化
	c := pb.NewGreeterClient(conn)

	// 参数
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 连接超时设置
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 远程调用
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 返回消息
	log.Printf("Greeting: %s", r.GetMessage())
}
