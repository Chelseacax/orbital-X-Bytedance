package main

import (
	"context"
	orbital "test1/kitex/kitex_gen/orbital"
)

type CombineServiceImpl struct{}

// Add implements the CalculatorServiceImpl interface.
func (s *CombineServiceImpl) Add(ctx context.Context, inputs *orbital.Variable) (resp *orbital.Answer, err error) {
	// TODO: Your code here...
	summation := inputs.FirstInt + inputs.SecondInt
	return &orbital.Answer{Output: summation}, nil
}

// Subtract implements the CalculatorServiceImpl interface.
func (s *CombineServiceImpl) Subtract(ctx context.Context, inputs *orbital.Variable) (resp *orbital.Answer, err error) {
	// TODO: Your code here...
	subtraction := inputs.FirstInt - inputs.SecondInt
	return &orbital.Answer{Output: subtraction}, nil
}

// HelloMethod implements the HelloServiceImpl interface.
func (s *CombineServiceImpl) HelloMethod(ctx context.Context, request *orbital.HelloRequest) (resp *orbital.HelloResponse, err error) {
	// TODO: Your code here...
	return &orbital.HelloResponse{ResponseBody: "Hello! :) Have a nice day"}, nil
}
