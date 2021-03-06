package main

import (
	"mir/com"
	"mir/srv"
	"net"
	"os"
)

func main() {

	e := com.InitEnviron()
	defer e.DB.Close()

	e.Cron = srv.InitScheduler()

	listener, err := net.Listen("tcp", com.Addr)
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
