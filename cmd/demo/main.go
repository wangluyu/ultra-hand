package demo

import (
	"flag"
	"fmt"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"ultra-hand/app/demo"
	"ultra-hand/pkg/log"
)

var (
	httpAddr    = flag.String("http-addr", ":9009", "HTTP listen address")
	rpcDemoAddr = flag.String("rpcDemo-addr", "127.0.0.1:9008", "rpc demo address")
)

type Demo struct {
	logger log.Logger
}

func New(logger log.Logger) Demo {
	return Demo{logger: logger}
}

func (d Demo) Start() func() error {
	return func() error {
		flag.Parse()

		svc := demo.NewService()
		svc = demo.NewLogMiddleware(d.logger, svc)

		rpcConn, err := grpc.Dial(*rpcDemoAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("connect rpc %s failed, err: %v", *rpcDemoAddr, err)
			return err
		}
		defer func(rpcConn *grpc.ClientConn) {
			err := rpcConn.Close()
			if err != nil {
				return
			}
		}(rpcConn)

		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			fmt.Printf("http: net.Listen(tcp, %s) failed, err:%v\n", *httpAddr, err)
			return err
		}
		defer func(httpListener net.Listener) {
			err := httpListener.Close()
			if err != nil {
				return
			}
		}(httpListener)
		httpHandler := demo.NewHTTPServer(svc, d.logger, rpcConn)
		return http.Serve(httpListener, httpHandler)
	}
}

var DemoProvideSet = wire.NewSet(New)
