// Code generated by Kitex v0.6.1. DO NOT EDIT.

package combineservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	orbital "test4/kitex/kitex_gen/orbital"
)

type CombineService interface {
	orbital.CalculatorService
	orbital.HelloService
}

func serviceInfo() *kitex.ServiceInfo {
	return combineServiceServiceInfo
}

var combineServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CombineService"
	handlerType := (*CombineService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Add":         kitex.NewMethodInfo(addHandler, newCalculatorServiceAddArgs, newCalculatorServiceAddResult, false),
		"HelloMethod": kitex.NewMethodInfo(helloMethodHandler, newHelloServiceHelloMethodArgs, newHelloServiceHelloMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "orbital",
	}
	extra["combine_service"] = true
	extra["combined_service_list"] = []string{"CalculatorService", "HelloService"}
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

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*orbital.CalculatorServiceAddArgs)
	realResult := result.(*orbital.CalculatorServiceAddResult)
	success, err := handler.(orbital.CalculatorService).Add(ctx, realArg.Inputs)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCalculatorServiceAddArgs() interface{} {
	return orbital.NewCalculatorServiceAddArgs()
}

func newCalculatorServiceAddResult() interface{} {
	return orbital.NewCalculatorServiceAddResult()
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

func (p *kClient) Add(ctx context.Context, inputs *orbital.Variable) (r *orbital.Result_, err error) {
	var _args orbital.CalculatorServiceAddArgs
	_args.Inputs = inputs
	var _result orbital.CalculatorServiceAddResult
	if err = p.c.Call(ctx, "Add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
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