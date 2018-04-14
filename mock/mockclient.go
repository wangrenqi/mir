package main

import (
	"net"
	"fmt"
	p "mir-go/proto"
	cp "mir-go/proto/client"
	"time"
)

//var host = "192.168.0.111"
var host = "127.0.0.1"

func main() {
	conn, err := net.Dial("tcp", host+":7000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("client version")
	pkg := &p.Packet{false, cp.CLIENT_VERSION, &cp.ClientVersion{}}
	bytes := pkg.ToBytes()
	conn.Write(bytes)
	time.Sleep(time.Second)

	fmt.Println("login")
	pkg = &p.Packet{false, cp.LOGIN, &cp.Login{"222", "222222"}}
	bytes = pkg.ToBytes()
	conn.Write(bytes)
	time.Sleep(time.Second)

	// TODO new character

	fmt.Println("start game")
	pkg = &p.Packet{false, cp.START_GAME, &cp.StartGame{}}
	bytes = pkg.ToBytes()
	conn.Write(bytes)
	time.Sleep(time.Second)

	// refine cancel ??

	//// in game

	// walk
	fmt.Println("walk")
	dir := []cp.Direction{cp.Up, cp.Right, cp.Down, cp.Left}
	for _, d := range dir {
		//pkg = &p.Packet{false, cp.WALK, &cp.Walk{cp.Up}}
		pkg = &p.Packet{false, cp.WALK, &cp.Walk{d}}
		bytes = pkg.ToBytes()
		conn.Write(bytes)
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 3)
	pkg = &p.Packet{false, cp.WALK, &cp.Walk{cp.Up}}
	bytes = pkg.ToBytes()
	conn.Write(bytes)

	// direction

}
