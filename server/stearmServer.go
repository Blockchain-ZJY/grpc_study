// package main

// import (
// 	"fmt"
// 	common "grpc_study/common"
// 	"net"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/grpclog"
// )

// type Common struct {
// 	common.UnimplementedCommonServiceServer
// }

// func main() {
// 	// 监听端口
// 	listen, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		grpclog.Fatalf("Failed to listen: %v", err)
// 	}

// 	// 创建一个gRPC服务器实例。
// 	s := grpc.NewServer()
// 	CommonServer := Common{}

// 	// 将server结构体注册为gRPC服务。
// 	common.RegisterCommonServiceServer(s, CommonServer)

// 	fmt.Println("grpc server running :8080")
// 	// 开始处理客户端请求。
// 	s.Serve(listen)
// }

package main

import (
	"fmt"
	common "grpc_study/common"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Common struct {
	common.UnimplementedCommonServiceServer
}

func (c Common) GetUserInfo(req *common.UserId, stream common.CommonService_GetUserInfoServer) error {
	fmt.Println("入参：", req.Id)

	// 发送多条用户信息
	users := []*common.UserInfo{
		{Name: "user1", Email: "user1@example.com"},
		{Name: "user2", Email: "user2@example.com"},
		{Name: "user23", Email: "user2@example.com"},
		{Name: "user42", Email: "user2@example.com"},
	}

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	CommonServer := Common{}

	// 将server结构体注册为gRPC服务。
	common.RegisterCommonServiceServer(s, CommonServer)

	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	s.Serve(listen)
}
