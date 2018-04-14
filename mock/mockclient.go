package main

import (
	"net"
	"fmt"
	p "mir-go/proto"
	cp "mir-go/proto/client"
	"time"
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
	time.Sleep(time.Second)

	// login
	pkg = &p.Packet{false, cp.LOGIN, &cp.Login{"222", "222222"}}
	bytes = pkg.ToBytes()
	conn.Write(bytes)
	time.Sleep(time.Second)

	// TODO new character

	// start game
	pkg = &p.Packet{false, cp.START_GAME, &cp.StartGame{}}
	bytes = pkg.ToBytes()
	conn.Write(bytes)
	time.Sleep(time.Second)

	// refine cancel ??

	//// in game

	// walk
	dir := []cp.Direction{cp.Up, cp.Right, cp.Down, cp.Left}
	for _, d := range dir {
		//pkg = &p.Packet{false, cp.WALK, &cp.Walk{cp.Up}}
		pkg = &p.Packet{false, cp.WALK, &cp.Walk{d}}
		bytes = pkg.ToBytes()
		conn.Write(bytes)
		time.Sleep(time.Second)
	}

	// direction

}
