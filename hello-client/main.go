package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "xxb-grpc-study/hello-server/proto"
)

//在 grpc 中，一个服务通常由一个或多个方法组成
//在 gRPC 中，一个服务通常是一个有多个方法的接口，这些方法定义了客户端可以请求服务的不同功能。如果一个服务只有一个方法，那么它也可以被视为独立的 grpc 服务。
//因此，无论一个服务有一个方法还是多个方法，都可以被视为一个 grpc 服务。
func main() {
	//通过grpc连接到server端，此处没有加密和验证
	/*
			1、grpc.Dail使用的是http/2，这种协议使用了多路复用技术，
		在一个tcp连接上同时处理多个请求和响应
		2、这个方法会尝试和服务器连接，如果连接成功，会返回一个ClientConn对象，
	该对象代表了一个服务器得到了稳定的连接*/
	conn, err := grpc.Dial("127.0.0.1:9098", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("没有连接 %v", err)
	}
	defer conn.Close() //所有的连接用完一定要关闭

	//建立连接
	//创建一个客户端
	client := pb.NewSayHelloClient(conn)
	//context.Background():会找到上下文请求
	//SayHello不是本地的，是调用远程的方法，执行rpc的调用
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "：jinli"})
	//看一下resp是什么
	fmt.Println(resp.GetResponseMsg())

}
