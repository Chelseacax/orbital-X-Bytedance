package main

import (
	"context"
	example "orbital/kitex/kitex_gen/example"
)

// ExampleServiceImpl implements the last service interface defined in the IDL.
type ExampleServiceImpl struct{}

// HelloMethod implements the ExampleServiceImpl interface.
func (s *ExampleServiceImpl) HelloMethod(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	// TODO: Your code here...
	return &example.HelloResp{RespBody: "Hello! :) Have a nice day"}, nil
}

// Add implements the ExampleServiceImpl interface.
func (s *ExampleServiceImpl) Add(ctx context.Context, inputs *example.Variable) (resp *example.Summer, err error) {
	// TODO: Your code here...
	summation := inputs.FirstInt + inputs.SecondInt
	return &example.Summer{Sum: summation}, nil
}
