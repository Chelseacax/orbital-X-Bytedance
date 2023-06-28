// Code generated by hertz generator.

package api

import (
	"context"
	"fmt"
	"zz/hertz-server/kitex_gen/student/management"
	"zz/hertz-server/kitex_gen/student/management/studentmanagement"

	client2 "github.com/cloudwego/kitex/client"

	api "zz/hertz-server/biz/model/api"

	"github.com/cloudwego/hertz/pkg/app"
)

// QueryStudent .
// @router student/query [GET]
func QueryStudent(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.QueryStudentRequest
	// Parameter binding and validation capabilities provided by hertz
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	// todo: do some other business processing

	// Use kitex client to make rpc calls and request back-end services
	client, err := studentmanagement.NewClient("student", client2.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	reqRpc := &management.QueryStudentRequest{
		Num: fmt.Sprintf("%d", req.Num),
	}

	respRpc, err := client.QueryStudent(ctx, reqRpc)
	if err != nil {
		panic(err)
	}

	// todo: Finish the rpc call and go do some other business logic

	if !respRpc.Exist {
		resp := &api.QueryStudentResponse{
			Msg: fmt.Sprintf("don't have the num: %d", req.Num),
		}
		c.JSON(200, resp)
		return
	}

	resp := &api.QueryStudentResponse{
		Num:    respRpc.Num,
		Name:   respRpc.Name,
		Gender: respRpc.Gender,
	}

	c.JSON(200, resp)
}

// InsertStudent .
// @router student/insert [POST]
func InsertStudent(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.InsertStudentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	client, err := studentmanagement.NewClient("student", client2.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	reqRpc := &management.InsertStudentRequest{
		Num:    req.Num,
		Name:   req.Name,
		Gender: req.Gender,
	}

	respRpc, err := client.InsertStudent(ctx, reqRpc)
	if err != nil {
		panic(err)
	}

	if !respRpc.Ok {
		resp := api.InsertStudentResponse{
			Ok:  false,
			Msg: respRpc.Msg,
		}
		c.JSON(200, resp)
		return
	}

	resp := api.InsertStudentResponse{
		Ok:  true,
		Msg: respRpc.Msg,
	}

	c.JSON(200, resp)
}
