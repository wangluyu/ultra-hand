package grpcdemo

import (
	"context"
	"ultra-hand/service/grpcdemo/pb"
)

type Service interface {
	pb.RpcDemoServer
}

type service struct {
	pb.UnimplementedRpcDemoServer
}

func (s service) Reply(_ context.Context, req *pb.ReplyRequest) (resp *pb.ReplyResponse, err error) {
	name := req.Name
	return &pb.ReplyResponse{Reply: "Hello " + name}, nil
}

func NewService() Service {
	return &service{}
}
