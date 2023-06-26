package main

import (
	server0 "github.com/cloudwego/kitex/gcf/idl/kitex_gen/kitex/test/server/echoservice"
	"log"
)

func main() {
	svr := server0.NewServer(new(EchoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
