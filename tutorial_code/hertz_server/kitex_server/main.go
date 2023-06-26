package main

import (
	management "github.com/cloudwego/hertz/hertz_server/kitex_client/kitex_gen/user/management/usermanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(UserManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
