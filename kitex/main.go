package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	orbital "test1/kitex/kitex_gen/orbital"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/generic/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	// path is the pathway to the IDL file
	path := "../idl/orbital.thrift"

	// This function call will change the default IDL parsing
	// from only last service to all services in the IDL file.
	// By default (i.e. w/o this func call),
	// the parser will only parse the last service declared in the IDL file.
	thrift.SetDefaultParseMode(thrift.CombineServices)

	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("New thrift file provider failed: %v", err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	// Service registry and discovery

	// This HTTP server (in Hertz) will run into an error
	// if the calculator microservice servers (in Kitex) are not up.
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	// Listening on IP address 127.0.0.1
	// Port 43000 & 43001 are used for hello microservice
	helloAddr0, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:43000")
	helloServer0 := genericserver.NewServer(new(GenericServiceImpl0), g,
		server.WithServiceAddr(helloAddr0),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello0"}),
		server.WithRegistry(r))

	helloAddr1, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:43001")
	helloServer1 := genericserver.NewServer(new(GenericServiceImpl1), g,
		server.WithServiceAddr(helloAddr1),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello1"}),
		server.WithRegistry(r))

	// Port 44000 & 44001 are used for calculator microservice
	calculatorAddr0, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:44000")
	calculatorServer0 := genericserver.NewServer(new(GenericServiceImpl2), g,
		server.WithServiceAddr(calculatorAddr0),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "calculator0"}),
		server.WithRegistry(r))

	calculatorAddr1, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:44001")
	calculatorServer1 := genericserver.NewServer(new(GenericServiceImpl3), g,
		server.WithServiceAddr(calculatorAddr1),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "calculator1"}),
		server.WithRegistry(r))

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		if err := helloServer0.Run(); err != nil {
			log.Println("server0 stopped with error:", err)
		} else {
			log.Println("server0 stopped")
		}
	}()
	go func() {
		defer wg.Done()
		if err := helloServer1.Run(); err != nil {
			log.Println("server1 stopped with error:", err)
		} else {
			log.Println("server1 stopped")
		}
	}()
	go func() {
		defer wg.Done()
		if err := calculatorServer0.Run(); err != nil {
			log.Println("calculator 0 stopped with error:", err)
		} else {
			log.Println("calculator 0 stopped")
		}
	}()
	go func() {
		defer wg.Done()
		if err := calculatorServer1.Run(); err != nil {
			log.Println("calculator 1 stopped with error:", err)
		} else {
			log.Println("calculator 1 stopped")
		}
	}()
	wg.Wait()
}

type GenericServiceImpl0 struct {
}

func (g *GenericServiceImpl0) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("KITEX: Using helloServer0")

	// resp is a JSON string
	resp, err := (*CombineServiceImpl).HelloMethod(new(CombineServiceImpl), ctx, &orbital.HelloRequest{Name: "me"})
	if err != nil {
		log.Fatal(err)
	}

	// return type of resp.ResponseBody is string
	// constructing a struct to marshal the string into JSON string
	respp := Response{
		Response: resp.ResponseBody,
	}

	// jsonData has a type of byte[]
	jsonData, err := json.Marshal(respp)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// converting jsonData into JSON string (byte[] --> interface{} / JSON string)
	// and returning it to Kitex client, which is from Hertz
	klog.Info("KITEX: Request completed")
	return string(jsonData), nil
}

type GenericServiceImpl1 struct {
}

func (g *GenericServiceImpl1) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("KITEX: Using helloServer1")

	// resp is a JSON string
	resp, err := (*CombineServiceImpl).HelloMethod(new(CombineServiceImpl), ctx, &orbital.HelloRequest{Name: "me"})
	if err != nil {
		log.Fatal(err)
	}

	// return type of resp.ResponseBody is string
	// constructing a struct to marshal the string into JSON string
	respp := Response{
		Response: resp.ResponseBody,
	}

	// jsonData has a type of byte[]
	jsonData, err := json.Marshal(respp)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// converting jsonData into JSON string (byte[] --> interface{} / JSON string)
	// and returning it to Kitex client, which is from Hertz
	klog.Info("KITEX: Request completed")
	return string(jsonData), nil
}

type GenericServiceImpl2 struct {
}

func (g *GenericServiceImpl2) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("KITEX: Using calculator 0")

	// request is a JSON string
	// converting it to a regular string
	m := request.(string)

	// converting JSON data (in the form of string) to Variable struct
	var inputs Variable
	err = json.Unmarshal([]byte(m), &inputs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var resp *orbital.Answer

	if method == "Add" {
		// Calling the function in the handler.go file (main logic of the function)
		resp, err = (*CombineServiceImpl).Add(new(CombineServiceImpl), ctx,
			&orbital.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
		if err != nil {
			log.Fatal(err)
		}

	} else if method == "Subtract" {
		// Calling the function in the handler.go file (main logic of the function)
		resp, err = (*CombineServiceImpl).Subtract(new(CombineServiceImpl), ctx,
			&orbital.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
		if err != nil {
			log.Fatal(err)
		}
	}

	// Constructing Answer struct to be converted to JSON data
	ans := Answer{
		Answer: resp.Output,
	}

	// jsonData has a type of byte[]
	jsonData, err := json.Marshal(ans)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Converting jsonData into JSON string (byte[] --> interface{} / JSON string)
	// and returning it to Kitex client, which is from Hertz
	klog.Info("KITEX: Calculation completed")
	return string(jsonData), nil
}

type GenericServiceImpl3 struct {
}

func (g *GenericServiceImpl3) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("KITEX: Using calculator 1")

	// request is a JSON string
	// converting it to a regular string
	m := request.(string)

	// converting JSON data (in the form of string) to Variable struct
	var inputs Variable
	err = json.Unmarshal([]byte(m), &inputs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var resp *orbital.Answer

	if method == "Add" {
		// Calling the function in the handler.go file (main logic of the function)
		resp, err = (*CombineServiceImpl).Add(new(CombineServiceImpl), ctx,
			&orbital.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
		if err != nil {
			log.Fatal(err)
		}

	} else if method == "Subtract" {
		// Calling the function in the handler.go file (main logic of the function)
		resp, err = (*CombineServiceImpl).Subtract(new(CombineServiceImpl), ctx,
			&orbital.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
		if err != nil {
			log.Fatal(err)
		}
	}

	// Constructing Answer struct to be converted to JSON data
	ans := Answer{
		Answer: resp.Output,
	}

	// jsonData has a type of byte[]
	jsonData, err := json.Marshal(ans)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Converting jsonData into JSON string (byte[] --> interface{} / JSON string)
	// and returning it to Kitex client, which is from Hertz
	klog.Info("KITEX: Calculation completed")
	return string(jsonData), nil
}

type Variable struct {
	FirstInt  int64 `json:"FirstInt"`
	SecondInt int64 `json:"SecondInt"`
}

type Answer struct {
	Answer int64 `json:"Output"`
}

type Response struct {
	Response string `json:"ResponseBody"`
}
