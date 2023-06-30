package grpcdemo

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"ultra-hand/service/grpcdemo"
	"ultra-hand/service/grpcdemo/pb"
)

var (
	grpcAddr = flag.String("grpc-addr", ":9008", "gRPC listen address")
)

func Start() func() error {
	svc := grpcdemo.NewService()
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
