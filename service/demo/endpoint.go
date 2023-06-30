package demo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"ultra-hand/service/grpcdemo/pb"
)

// 一个 endpoint 表示对外提供的一个方法

type CallRequest struct {
	Name string `json:"name"`
}

type CallResponse struct {
	Reply string `json:"reply"`
}

// 借助 适配器 将 方法 -> endpoint
func makeCallEndpoint(svc Service, rpcConn *grpc.ClientConn) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		rpcReplyEndpoint := makeReplyEndpoint(rpcConn)
		req := request.(CallRequest)
		reply, err := svc.Call(ctx, req, rpcReplyEndpoint)
		if err != nil {
			return CallResponse{Reply: reply}, nil
		}
		return CallResponse{Reply: reply}, nil
	}
}

// 调用 rpc

type rpcReplyRequest struct {
	name string
}

type rpcReplyResponse struct {
	reply string
}

// 客户端 endpoint, 请求其他服务
func makeReplyEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn,
		"pb.RpcDemo",
		"Reply",
		encodeReplyRequest,
		decodeReplyResponse,
		pb.ReplyResponse{},
	).Endpoint()
}
