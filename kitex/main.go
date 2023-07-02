package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	example "orbital/kitex/kitex_gen/example"
	"sync"

	// Required for kitex generic call feature server
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server/genericserver"

	// Required for LoadBalancing
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// Parameter: idl / thrift file path
	p, err := generic.NewThriftFileProvider("../idl/example.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	// LoadBalancing (with 2 Kitex servers)

	// Server0
	// Listening on IP address 0.0.0.0 & port number 8888
	addr0, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8888")
	server0 := genericserver.NewServer(new(GenericServiceImpl0), g, server.WithServiceAddr(addr0))

	// Server1
	// Listening on IP address 0.0.0.0 & port number 8888\9
	addr1, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8889")
	server1 := genericserver.NewServer(new(GenericServiceImpl1), g, server.WithServiceAddr(addr1))

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := server0.Run(); err != nil {
			log.Println("server0 stopped with error:", err)
		} else {
			log.Println("server0 stopped")
		}
	}()
	go func() {
		defer wg.Done()
		if err := server1.Run(); err != nil {
			log.Println("server1 stopped with error:", err)
		} else {
			log.Println("server1 stopped")
		}
	}()
	wg.Wait()
}

type GenericServiceImpl0 struct {
}

func (g *GenericServiceImpl0) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("Using server0")

	if method == "HelloMethod" {
		// request is a JSON string
		// converting it to a regular string
		m := request.(string)

		// Print the value of m
		fmt.Printf("Recv: %v\n", m)

		// resp is a JSON string
		resp, err := (*ExampleServiceImpl).HelloMethod(new(ExampleServiceImpl), context.Background(), &example.HelloReq{Name: "me"})
		if err != nil {
			log.Fatal(err)
		}

		// return type of resp.RespBody is string
		// constructing a struct to marshal the string into JSON string
		respp := Response{
			Response: resp.RespBody,
		}

		// jsonData has a type of byte[]
		jsonData, err := json.Marshal(respp)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// converting jsonData into JSON string (byte[] --> interface{})
		return string(jsonData), nil

	} else if method == "Add" {
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

		resp, err := (*ExampleServiceImpl).Add(new(ExampleServiceImpl), context.Background(),
			&example.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
		if err != nil {
			log.Fatal(err)
		}

		// constructing summation struct to be converted to JSON
		sum := Summation{
			Summation: resp.Sum,
		}

		jsonData, err := json.Marshal(sum)
		if err != nil {
			fmt.Println("Error:", err)
		}

		return string(jsonData), nil
	}
	return
}

type GenericServiceImpl1 struct {
}

func (g *GenericServiceImpl1) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	klog.Info("Using server1")

	if method == "HelloMethod" {
		// request is a JSON string
		// converting it to a regular string
		m := request.(string)

		// Print the value of m
		fmt.Printf("Recv: %v\n", m)

		// resp is a JSON string
		resp, err := (*ExampleServiceImpl).HelloMethod(new(ExampleServiceImpl), context.Background(), &example.HelloReq{Name: "me"})
		if err != nil {
			log.Fatal(err)
		}

		// return type of resp.RespBody is string
		// constructing a struct to marshal the string into JSON string
		respp := Response{
			Response: resp.RespBody,
		}

		// jsonData has a type of byte[]
		jsonData, err := json.Marshal(respp)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// converting jsonData into JSON string (byte[] --> interface{})
		return string(jsonData), nil

	} else if method == "Add" {
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

		resp, err := (*ExampleServiceImpl).Add(new(ExampleServiceImpl), context.Background(),
			&example.Variable{FirstInt: inputs.FirstInt, SecondInt: inputs.SecondInt})
		if err != nil {
			log.Fatal(err)
		}

		// constructing summation struct to be converted to JSON
		sum := Summation{
			Summation: resp.Sum,
		}

		jsonData, err := json.Marshal(sum)
		if err != nil {
			fmt.Println("Error:", err)
		}

		return string(jsonData), nil
	}
	return
}

type Response struct {
	Response string `json:"RespBody"`
}

type Variable struct {
	FirstInt  int64 `json:"FirstInt"`
	SecondInt int64 `json:"SecondInt"`
}

type Summation struct {
	Summation int64 `json:"Sum"`
}
