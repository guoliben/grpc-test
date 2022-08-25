package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc_proto "grpc-my-test/grpc-test/service"
	"log"
	"net"
)

type server struct{
	hello_grpc_proto.UnimplementedGreeterServer
	hello_grpc_proto.UnimplementedWeaknessServer
}

func (s *server) SayHello(ctx context.Context, in *hello_grpc_proto.HelloRequest)(resp *hello_grpc_proto.HelloReply, err error){
	log.Printf("Received: %v", in.GetName())
	return &hello_grpc_proto.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) WeaknessUpload(ctx context.Context, in *hello_grpc_proto.WeaknessScanRequest)(resp *hello_grpc_proto.WeaknessScanResponse, err error){
	log.Printf("Received: %v", in.GetSlice())
	return &hello_grpc_proto.WeaknessScanResponse{Ack: "Hello " + in.GetSlice() +" ----"+in.Slice+" "+in.Time+" "+in.Detail+" "+in.Message+" " }, nil
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5500))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC 服务器
	s := grpc.NewServer()

	// 将服务器与处理器绑定
	hello_grpc_proto.RegisterGreeterServer(s, &server{})
	hello_grpc_proto.RegisterWeaknessServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}