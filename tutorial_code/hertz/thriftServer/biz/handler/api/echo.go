// Code generated by hertz generator.

package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HelloMethod .
// @router /hello [GET]
func EchoMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.EchoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(api.EchoResp)

	// You can modify the logic of the entire function, not just the current template
	resp.message = req.Name +"completed!" // added logic

	c.JSON(200, resp)
}
