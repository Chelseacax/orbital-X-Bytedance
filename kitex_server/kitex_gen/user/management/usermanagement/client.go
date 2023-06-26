// Code generated by Kitex v0.6.1. DO NOT EDIT.

package usermanagement

import (
	"context"
	management "github.com/cloudwego/hertz/hertz_server/kitex_client/kitex_gen/user/management"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	QueryUser(ctx context.Context, req *management.QueryUserRequest, callOptions ...callopt.Option) (r *management.QueryUserResponse, err error)
	AddUser(ctx context.Context, req *management.AddUserRequest, callOptions ...callopt.Option) (r *management.AddUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserManagementClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserManagementClient struct {
	*kClient
}

func (p *kUserManagementClient) QueryUser(ctx context.Context, req *management.QueryUserRequest, callOptions ...callopt.Option) (r *management.QueryUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUser(ctx, req)
}

func (p *kUserManagementClient) AddUser(ctx context.Context, req *management.AddUserRequest, callOptions ...callopt.Option) (r *management.AddUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddUser(ctx, req)
}