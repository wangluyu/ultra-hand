package demo

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"ultra-hand/service/demo"
)

var (
	httpAddr    = flag.String("http-addr", ":9009", "HTTP listen address")
	rpcDemoAddr = flag.String("rpcDemo-addr", "127.0.0.1:9008", "rpc demo address")
)

func Start() func() error {
	return func() error {
		flag.Parse()

		logger := log.NewLogfmtLogger(os.Stderr)
		svc := demo.NewService()
		svc = demo.NewLogMiddleware(logger, svc)

		rpcConn, err := grpc.Dial(*rpcDemoAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("connect rpc %s failed, err: %v", *rpcDemoAddr, err)
			return err
		}
		defer rpcConn.Close()

		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			fmt.Printf("http: net.Listen(tcp, %s) failed, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		httpHandler := demo.NewHTTPServer(svc, logger, rpcConn)
		return http.Serve(httpListener, httpHandler)
	}
}
