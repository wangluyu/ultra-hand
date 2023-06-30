package demo

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"ultra-hand/service/demo"
)

var (
	httpAddr = flag.String("http-addr", ":9009", "HTTP listen address")
)

func Start() func() error {
	svc := demo.NewService()

	return func() error {
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			fmt.Printf("http: net.Listen(tcp, %s) failed, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		httpHandler := demo.NewHTTPServer(svc)
		return http.Serve(httpListener, httpHandler)
	}
}
