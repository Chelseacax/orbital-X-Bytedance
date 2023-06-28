package main

import (
	management "orbital-X-Bytedance/kitex_server/kitex_gen/user/management/usermanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(UserManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
