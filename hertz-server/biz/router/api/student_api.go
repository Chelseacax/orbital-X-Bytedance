// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "zz/hertz-server/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_student := root.Group("/student", _studentMw()...)
		_student.POST("/insert", append(_insertstudentMw(), api.InsertStudent)...)
		_student.GET("/query", append(_querystudentMw(), api.QueryStudent)...)
	}
}
