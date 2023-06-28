package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ExampleMethod(ctx context.Context, c *app.RequestContext) {
	path := "../idl/example_service.thrift" // depends on current directory
	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new map thrift generic failed: %v", err)
	}
	cli, err := genericclient.NewClient("example", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "ExampleMethod", "{\"Msg\": \"Your mama so fat\"}")
	// resp is a JSON string
	s := resp.(string)

	fmt.Printf(s)
}
