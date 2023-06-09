// Code generated by Kitex v0.6.1. DO NOT EDIT.

package exampleservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	example "orbital/hertz/kitex_gen/example"
)

func serviceInfo() *kitex.ServiceInfo {
	return exampleServiceServiceInfo
}

var exampleServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ExampleService"
	handlerType := (*example.ExampleService)(nil)
	methods := map[string]kitex.MethodInfo{
		"HelloMethod": kitex.NewMethodInfo(helloMethodHandler, newExampleServiceHelloMethodArgs, newExampleServiceHelloMethodResult, false),
		"Add":         kitex.NewMethodInfo(addHandler, newExampleServiceAddArgs, newExampleServiceAddResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "example",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func helloMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ExampleServiceHelloMethodArgs)
	realResult := result.(*example.ExampleServiceHelloMethodResult)
	success, err := handler.(example.ExampleService).HelloMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newExampleServiceHelloMethodArgs() interface{} {
	return example.NewExampleServiceHelloMethodArgs()
}

func newExampleServiceHelloMethodResult() interface{} {
	return example.NewExampleServiceHelloMethodResult()
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.ExampleServiceAddArgs)
	realResult := result.(*example.ExampleServiceAddResult)
	success, err := handler.(example.ExampleService).Add(ctx, realArg.Inputs)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newExampleServiceAddArgs() interface{} {
	return example.NewExampleServiceAddArgs()
}

func newExampleServiceAddResult() interface{} {
	return example.NewExampleServiceAddResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) HelloMethod(ctx context.Context, request *example.HelloReq) (r *example.HelloResp, err error) {
	var _args example.ExampleServiceHelloMethodArgs
	_args.Request = request
	var _result example.ExampleServiceHelloMethodResult
	if err = p.c.Call(ctx, "HelloMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Add(ctx context.Context, inputs *example.Variable) (r *example.Summer, err error) {
	var _args example.ExampleServiceAddArgs
	_args.Inputs = inputs
	var _result example.ExampleServiceAddResult
	if err = p.c.Call(ctx, "Add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
