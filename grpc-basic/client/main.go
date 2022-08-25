package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc_proto "grpc-my-test/grpc-test/service"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:5500", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	defer conn.Close()


	c := hello_grpc_proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// 虽然SayHello由远程服务端实现，但调用起来就像一个本地函数一样
	r, err := c.SayHello(ctx, &hello_grpc_proto.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("call service failed: %v", err)
	}
	fmt.Println("call service success: ", r.Message)

}