// Code generated by hertz generator.

package main

import (
	handler "zz/hertz-server/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.GET("/gcf", handler.ExampleMethod)

	// your code ...
}