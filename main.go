package main

import (
	"net"
	"log"
	"os"
	"mir/env"
	"mir/srv"
)

//var addr = "127.0.0.1:7000"
var addr = "192.168.0.100:7000"

func main() {

	env := env.InitEnviron()
	defer env.Db.Close()

	listener, err := net.Listen("tcp", addr)
	defer listener.Close()
	if err != nil {
		log.Fatalln("start server error")
		os.Exit(-1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("ERROR: %s", err)
			continue
		}
		log.Println("new connection: ", conn.RemoteAddr())

		go srv.HandleClient(conn, env)
	}
}
