// Code generated by Kitex v0.6.1. DO NOT EDIT.

package helloservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	orbital "test4/kitex/kitex_gen/orbital"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceServiceInfo
}

var helloServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "HelloService"
	handlerType := (*orbital.HelloService)(nil)
	methods := map[string]kitex.MethodInfo{
		"HelloMethod": kitex.NewMethodInfo(helloMethodHandler, newHelloServiceHelloMethodArgs, newHelloServiceHelloMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "orbital",
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
	realArg := arg.(*orbital.HelloServiceHelloMethodArgs)
	realResult := result.(*orbital.HelloServiceHelloMethodResult)
	success, err := handler.(orbital.HelloService).HelloMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServiceHelloMethodArgs() interface{} {
	return orbital.NewHelloServiceHelloMethodArgs()
}

func newHelloServiceHelloMethodResult() interface{} {
	return orbital.NewHelloServiceHelloMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) HelloMethod(ctx context.Context, request *orbital.HelloRequest) (r *orbital.HelloResponse, err error) {
	var _args orbital.HelloServiceHelloMethodArgs
	_args.Request = request
	var _result orbital.HelloServiceHelloMethodResult
	if err = p.c.Call(ctx, "HelloMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}