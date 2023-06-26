// Code generated by Kitex v0.6.1. DO NOT EDIT.

package usermanagement

import (
	"context"
	management "github.com/cloudwego/hertz/hertz_server/kitex_client/kitex_gen/user/management"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userManagementServiceInfo
}

var userManagementServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserManagement"
	handlerType := (*management.UserManagement)(nil)
	methods := map[string]kitex.MethodInfo{
		"queryUser": kitex.NewMethodInfo(queryUserHandler, newUserManagementQueryUserArgs, newUserManagementQueryUserResult, false),
		"AddUser":   kitex.NewMethodInfo(addUserHandler, newUserManagementAddUserArgs, newUserManagementAddUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "management",
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

func queryUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*management.UserManagementQueryUserArgs)
	realResult := result.(*management.UserManagementQueryUserResult)
	success, err := handler.(management.UserManagement).QueryUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserManagementQueryUserArgs() interface{} {
	return management.NewUserManagementQueryUserArgs()
}

func newUserManagementQueryUserResult() interface{} {
	return management.NewUserManagementQueryUserResult()
}

func addUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*management.UserManagementAddUserArgs)
	realResult := result.(*management.UserManagementAddUserResult)
	success, err := handler.(management.UserManagement).AddUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserManagementAddUserArgs() interface{} {
	return management.NewUserManagementAddUserArgs()
}

func newUserManagementAddUserResult() interface{} {
	return management.NewUserManagementAddUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) QueryUser(ctx context.Context, req *management.QueryUserRequest) (r *management.QueryUserResponse, err error) {
	var _args management.UserManagementQueryUserArgs
	_args.Req = req
	var _result management.UserManagementQueryUserResult
	if err = p.c.Call(ctx, "queryUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddUser(ctx context.Context, req *management.AddUserRequest) (r *management.AddUserResponse, err error) {
	var _args management.UserManagementAddUserArgs
	_args.Req = req
	var _result management.UserManagementAddUserResult
	if err = p.c.Call(ctx, "AddUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
