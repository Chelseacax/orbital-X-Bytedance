// Code generated by Kitex v0.6.1. DO NOT EDIT.
package studentmanagement

import (
	server "github.com/cloudwego/kitex/server"
	management "zz/kitex-server/kitex_gen/student/management"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler management.StudentManagement, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
