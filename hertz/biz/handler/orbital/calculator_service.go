// Code generated by hertz generator.

package orbital

import (
	"context"
	"encoding/json"
	"fmt"

	orbital "test4/hertz/biz/model/orbital"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

// Load Balancing Options
var lb = loadbalance.NewWeightedRandomBalancer()
var calculatorServer0 = "127.0.0.1:44000"
var calculatorServer1 = "127.0.0.1:44001"

// Add .
// @router /add [POST]
func Add(ctx context.Context, c *app.RequestContext) {

	var err error

	// validating the inputs given
	var req orbital.Variable
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	klog.Info("Passing through Hertz to Kitex")

	// kitex's generic call feature
	path := "../idl/orbital.thrift"
	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("New thrift file provider failed: %v", err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		klog.Fatalf("New map thrift generic failed: %v", err)
	}

	// "calculator" is the service name
	cli, err := genericclient.NewClient("calculator", g,
		client.WithHostPorts(calculatorServer0, calculatorServer1),
		client.WithLoadBalancer(lb))
	if err != nil {
		klog.Fatalf("New HTTP generic client failed: %v", err)
	}

	// constructing struct for the JSON data (retrived as integers)
	var variables Variable
	variables = Variable{
		FirstInt:  req.FirstInt,
		SecondInt: req.SecondInt,
	}

	// Marhsalling the integers into jsonData (variable struct{int64, int64} --> []byte)
	jsonData, err := json.Marshal(variables)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// "Add" is the method name
	// the string method changes []byte into inteface{} based on ASCII
	resp, err := cli.GenericCall(context.Background(), "Add", string(jsonData))
	klog.Info(resp)
	// Unmarshalling of JSON string into Summer struct (string --> Result struct)
	s := resp.(string)
	var answer Result
	err = json.Unmarshal([]byte(s), &answer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ans := utils.H{"The asnwer is": answer.Answer}
	c.JSON(consts.StatusOK, ans)

	klog.Info("Hertz: Calculation completed")
}

type Variable struct {
	FirstInt  int64 `json:"FirstInt"`
	SecondInt int64 `json:"SecondInt"`
}

type Result struct {
	Answer int64 `json:"Output"`
}
