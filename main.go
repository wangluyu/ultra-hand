package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"ultra-hand/cmd/demo"
	"ultra-hand/cmd/grpcdemo"
	"ultra-hand/pkg/config"
)

func main() {
	var g errgroup.Group
	demoApp, err := demo.InitGrpcDemo(config.Path("D:\\开发\\ultra-hand\\config\\config.yaml"))
	if err != nil {
		fmt.Println(err)
		return
	}
	g.Go(demoApp.Start())
	grpcDemoApp, err := grpcdemo.InitGrpcDemo(config.Path("D:\\开发\\ultra-hand\\config\\config.yaml"))
	if err != nil {
		fmt.Println(err)
		return
	}
	g.Go(grpcDemoApp.Start())
	fmt.Println("START!")
	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with err:%v\n", err)
	}
}
