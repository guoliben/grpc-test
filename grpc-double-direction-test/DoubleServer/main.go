package main

import (
	"fmt"
	"google.golang.org/grpc"
	hello_grpc_proto "grpc-my-test/grpc-double-direction-test/service"
	"net"
	"time"
)

// 实现RouteGuideServer接口
type routeGuideServer struct {
	hello_grpc_proto.UnimplementedRouteGuidServer
}


//func (s *routeGuideServer) RouteChat(ctx context.Context, in *hello_grpc_proto.RouteNoteRequest)(resp *hello_grpc_proto.RouteNoteResponse, err error){
//	log.Printf("Received: %v", in.GetSlice())
//	return &hello_grpc_proto.RouteNoteResponse{Ack: "Hello " + in.GetSlice()}, nil
//}


func main() {
	// 监听,返回listener
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5500))
	fmt.Println(err)

	grpcServer := grpc.NewServer()
	// 注册服务
	hello_grpc_proto.RegisterRouteGuidServer(grpcServer, &routeGuideServer{})
	// 使用服务  ,grpcServer.Serve(lis)内部调用 rawConn, err := lis.Accept() ->   s.handleRawConn(rawConn)
	grpcServer.Serve(lis)





}


func (s *routeGuideServer) RouteChat(stream hello_grpc_proto.RouteGuid_RouteChatServer) error {
	////发送客户端心跳
	//go func() {
	//	fmt.Println("==============发送客户端消息=============")
	//
	//	for {
	//		time.Sleep(time.Second * 1)
	//
	//
	//	}
	//}()


	for {
		fmt.Println("==============读取客户端消息=============")
		// 接收
		in, err := stream.Recv()
		fmt.Println(in)

		fmt.Println("==============发送客户端消息=============")

		// 发送
		err = stream.Send(&hello_grpc_proto.RouteNoteResponse{
			Ack: "World",
		} )
		fmt.Println(err)
		time.Sleep(time.Second * 1)
	}

}


