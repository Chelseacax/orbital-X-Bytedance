package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	orbital "test4/kitex/kitex_gen/orbital"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
)

func main() {
	// Parameter: idl / thrift file path
	p, err := generic.NewThriftFileProvider("../idl/orbital.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	// Server0
	// Listening on IP address 0.0.0.0 & port number 8888
	helloAddr0, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:43000")
	helloServer0 := genericserver.NewServer(new(GenericServiceImpl0), g, server.WithServiceAddr(helloAddr0))

	helloAddr1, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:43001")
	helloServer1 := genericserver.NewServer(new(GenericServiceImpl1), g, server.WithServiceAddr(helloAddr1))

	calculatorAddr0, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:44000")
	calculatorServer0 := genericserver.NewServer(new(GenericServiceImpl2), g, server.WithServiceAddr(calculatorAddr0))

	calculatorAddr1, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:44001")
	calculatorServer1 := genericserver.NewServer(new(GenericServiceImpl3), g, server.WithServiceAddr(calculatorAddr1))

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

	klog.Info("Using helloServer0")
	// resp is a JSON string
	resp, err := (*HelloServiceImpl).HelloMethod(new(HelloServiceImpl), ctx, &orbital.HelloRequest{Name: "me"})
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
	klog.Info("Kitex: Request completed")
	return string(jsonData), nil
}

type GenericServiceImpl1 struct {
}

func (g *GenericServiceImpl1) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("Using helloServer1")
	// resp is a JSON string
	resp, err := (*HelloServiceImpl).HelloMethod(new(HelloServiceImpl), ctx, &orbital.HelloRequest{Name: "me"})
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
	klog.Info("Kitex: Request completed")
	return string(jsonData), nil
}

type GenericServiceImpl2 struct {
}

func (g *GenericServiceImpl2) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("Using calculator 0")
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

	resp, err := (*CalculatorServiceImpl).Add(new(CalculatorServiceImpl), ctx,
		&orbital.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
	if err != nil {
		log.Fatal(err)
	}

	// constructing Result struct to be converted to JSON
	sum := Result{
		Result: resp.Output,
	}

	// jsonData has a type of byte[]
	jsonData, err := json.Marshal(sum)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// converting jsonData into JSON string (byte[] --> interface{} / JSON string)
	// and returning it to Kitex client, which is from Hertz
	klog.Info("Kitex: Calculation completed")
	return string(jsonData), nil

}

type GenericServiceImpl3 struct {
}

func (g *GenericServiceImpl3) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("Using calculator 1")
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

	resp, err := (*CalculatorServiceImpl).Add(new(CalculatorServiceImpl), ctx,
		&orbital.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
	if err != nil {
		log.Fatal(err)
	}

	// constructing Result struct to be converted to JSON
	sum := Result{
		Result: resp.Output,
	}

	// jsonData has a type of byte[]
	jsonData, err := json.Marshal(sum)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// converting jsonData into JSON string (byte[] --> interface{} / JSON string)
	// and returning it to Kitex client, which is from Hertz
	klog.Info("Kitex: Calculation completed")
	return string(jsonData), nil

}

type Variable struct {
	FirstInt  int64 `json:"FirstInt"`
	SecondInt int64 `json:"SecondInt"`
}

type Result struct {
	Result int64 `json:"Output"`
}

type Response struct {
	Response string `json:"ResponseBody"`
}
