package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// Hertz server will be listening on IP address 127.0.0.1
	// and the port number is 42000
	h := server.Default(server.WithHostPorts("127.0.0.1:42000"))

	register(h)
	h.Spin()
}
