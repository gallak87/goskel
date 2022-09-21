package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gallak87/goskel/pkg/server"
	"github.com/golang/glog"
)

var (
	addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	fmt.Printf("goskel: Starting server %s\n\n", *addr)
	ctx := context.Background()
	if err := server.Run(ctx, *network, *addr); err != nil {
		glog.Fatal(err)
	}
}
