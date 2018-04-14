package main

import (
	"net"
	"fmt"
	p "mir-go/proto"
	cp "mir-go/proto/client"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.111:7000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// client version
	pkg := &p.Packet{false, cp.CLIENT_VERSION, &cp.ClientVersion{}}
	bytes := pkg.ToBytes()
	conn.Write(bytes)

	// login

	// TODO new character

	// start game

	// refine cancel ??

	//// in game

	// walk

	// direction

}
