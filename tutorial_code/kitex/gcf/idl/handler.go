package main

import (
	"context"
	server0 "github.com/cloudwego/kitex/gcf/idl/kitex_gen/kitex/test/server"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// EchoMethod implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) EchoMethod(ctx context.Context, req *server0.EchoReq) (resp *server0.EchoResp, err error) {
	// TODO: Your code here...
	return
}
