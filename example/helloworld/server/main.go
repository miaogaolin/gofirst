package main

import (
	"context"
	"log"
	"net"

	pb "github.com/miaogaolin/gofirst/example/helloworld"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// 实现服务端 GreeterServer 接口
type server struct {
	// 必须嵌套
	// 包含了 GreeterServer 接口其它方法实现
	pb.UnimplementedGreeterServer
}

// 远程调用方法的具体逻辑
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 监听端口，用于接受客户端的请求
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 服务端实例化
	s := grpc.NewServer()

	// 将具体的实现注册到服务端
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	// 阻塞等待
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
