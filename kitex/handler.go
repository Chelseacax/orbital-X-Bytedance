package main

import (
	"context"
	orbital "test4/kitex/kitex_gen/orbital"
)

// CalculatorServiceImpl implements the last service interface defined in the IDL.
type CalculatorServiceImpl struct{}

// Add implements the CalculatorServiceImpl interface.
func (s *CalculatorServiceImpl) Add(ctx context.Context, inputs *orbital.Variable) (resp *orbital.Result_, err error) {
	// TODO: Your code here...
	summation := inputs.FirstInt + inputs.SecondInt
	return &orbital.Result_{Output: summation}, nil
}

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *orbital.HelloRequest) (resp *orbital.HelloResponse, err error) {
	// TODO: Your code here...
	return &orbital.HelloResponse{ResponseBody: "Hello! :) Have a nice day"}, nil
}
