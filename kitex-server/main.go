package main

import (
	"log"
	// management "zz/kitex-server/kitex_gen/student/management/studentmanagement"

	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server/genericserver"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
    "github.com/cloudwego/kitex/server"
    etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	/*
		svr := management.NewServer(new(StudentManagementImpl))

		err := svr.Run()

		if err != nil {
			log.Println(err.Error())
		}
	*/

	// Parse IDL with Local Files
	// YOUR_IDL_PATH thrift file path,eg: ./idl/example.thrift
	p, err := generic.NewThriftFileProvider("../idl/example_service.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
    if err != nil {
        panic(err)
    }

	svr := genericserver.NewServer(new(GenericServiceImpl), g,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "example"}), server.WithRegistry(r))
	if err != nil {
		panic(err)
	}




	
	// svr := server0.NewServer(new(ExampleServiceImpl))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

type GenericServiceImpl struct {
}

func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	// use jsoniter or other json parse sdk to assert request
	m := request.(string)

	// print the value of m
	fmt.Printf("Recv: %v\n", m)

	// return to the client (in JSON string)
	return "{\"Msg\": \"same to you\"}", nil
}
