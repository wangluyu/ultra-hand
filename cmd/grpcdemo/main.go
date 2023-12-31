package grpcdemo

import (
	"flag"
	"fmt"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"net"
	"ultra-hand/app/grpcdemo"
	"ultra-hand/app/grpcdemo/pb"
	"ultra-hand/pkg/log"
)

var (
	grpcAddr = flag.String("grpc-addr", ":9008", "gRPC listen address")
)

type GrpcDemo struct {
	logger log.Logger
}

func New(logger log.Logger) GrpcDemo {
	return GrpcDemo{logger: logger}
}

func (gd GrpcDemo) Start() func() error {
	svc := grpcdemo.NewService()
	gd.logger.Info("GrpcDemo")
	return func() error {
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			fmt.Printf("grpc: net.Listen(tcp, %s) faield, err:%v\n", *grpcAddr, err)
			return err
		}
		s := grpc.NewServer()
		pb.RegisterRpcDemoServer(s, svc)
		err = s.Serve(grpcListener)
		if err != nil {
			fmt.Printf("failed to grpc server: %v", err)
			return err
		}
		return nil
	}
}

var GrpcProvideSet = wire.NewSet(New)
