package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc_proto "grpc-my-test/grpc-double-direction-test/service"
	"log"
	"time"
)

func main() {

	var notes []string
	notes = append(notes, "dfafad")
	notes = append(notes, "dfafad3")
	notes = append(notes, "dfafada")
	notes = append(notes, "dfafad32")

	conn, err := grpc.Dial("127.0.0.1:5500", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	defer conn.Close()


	c := hello_grpc_proto.NewRouteGuidClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stream, err := c.RouteChat(ctx)

	go func() {
		for {
			fmt.Println("==============读取服务端消息=============")
			// 连续读取数据
			time.Sleep(time.Second * 1)

			in, err := stream.Recv()
			fmt.Println(in, err)
		}
	}()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("==============发送服务端消息=============")

		// 连续发送数据
		stream.Send(&hello_grpc_proto.RouteNoteRequest{
			Slice: "World",
			Time: "what's time",
			Block: "which block",
			Detail: "this is detail ",
			Message: "upload",
		} );
	}
	// 使用完记得关闭
	stream.CloseSend()



}