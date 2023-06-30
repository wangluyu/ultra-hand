package grpcdemo

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"ultra-hand/service/grpcdemo/pb"
)

func encodeReplyRequest(_ context.Context, response interface{}) (request interface{}, err error) {
	req := grpcRequest.(*pb.ReplyRequest)
	return pb.ReplyRequest{Name: req.Name}, nil
}

func encodeGRPCCallResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	resp := grpcResponse.(pb.ReplyResponse)
	return pb.ReplyResponse{Reply: resp.Reply}, nil
}

type grpcDemoServer struct {
	pb.UnimplementedRpcDemoServer

	reply grpctransport.Handler
}

func NewGRPCServer(svc Service) pb.RpcDemoServer {
	return &grpcDemoServer{
		reply: grpctransport.NewServer(
			makeReplyEndpoint(svc),
			decodeGRPCCallRequest,
			encodeGRPCCallResponse,
		),
	}
}
