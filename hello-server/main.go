package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "xxb-grpc-study/hello-server/proto" //起了一个别名
)

type server struct {
	pb.UnimplementedSayHelloServer //嵌入了UnimplementedSayHelloServer这个接口
}

//重写sayHello方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Printf("hello:" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	//1.开启端口
	//net.Listen 方法的第二个参数应该是一个 IP 地址(:9098),9098只是一个端口号
	listen, _ := net.Listen("tcp", ":9098")
	//创建grpc
	grpcServer := grpc.NewServer()
	//&server{}是将server类型的实例化对象传递给pb.RegisterSayHelloServer()方法的参数，
	//该方法将该服务注册到grpc中，让grpc服务器能够接收该服务的调用请求并且找到实现该服务的方法。
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to erve:%v", err)
		return
	}
}
