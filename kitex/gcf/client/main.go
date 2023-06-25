package main

import (
	"context"
    "fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	path := "../idl/echo_service.thrift" // depends on current directory
	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new map thrift generic failed: %v", err)
	}
	cli, err := genericclient.NewClient("echo", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}
    resp, err := cli.GenericCall(context.Background() , "EchoMethod", "{\"Msg\": \"hello\"}")
    // resp is a JSON string
	s := resp.(string)

	fmt.Printf(s)
}
