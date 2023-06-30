package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"ultra-hand/cmd/demo"
	"ultra-hand/cmd/grpcdemo"
)

func main() {
	var g errgroup.Group
	g.Go(demo.Start())
	g.Go(grpcdemo.Start())
	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with err:%v\n", err)
	}
}
