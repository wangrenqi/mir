package main

import (
	"net"
	"os"
	"mir/srv"
	"mir/com"
)

//var addr = "127.0.0.1:7000"
var addr = "192.168.0.110:7000"

func main() {

	e := com.InitEnviron()
	defer e.DB.Close()

	listener, err := net.Listen("tcp", addr)
	defer listener.Close()
	if err != nil {
		os.Exit(-1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go srv.HandleClient(conn, e)
	}
}
