package main

import (
	"net"
	"log"
	"os"
	"mir-go/env"
	"mir-go/proto"
)

var addr = "127.0.0.1:7000"
var id int32 = 0
var packetChan = make(chan *proto.Packet)

func main() {

	env := env.InitEnviron()

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

		go proto.BytesToPacket(conn, packetChan)
		go handleClient(conn, env)
	}
}
