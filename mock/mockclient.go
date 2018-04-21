package main

import (
	"net"
	"fmt"
	p "mir/proto"
	cp "mir/proto/client"
	"time"
)

//var host = "192.168.0.111"
var host = "127.0.0.1"

func send(conn net.Conn, pkg *p.Packet) {
	bytes := pkg.ToBytes(false)
	// fmt.Println("raw bytes: ", bytes, "len", len(bytes))
	data := p.Pack(bytes)
	conn.Write(data)
	// fmt.Println("after pack: ", data, "len", len(data))
}

func main() {
	conn, err := net.Dial("tcp", host+":7000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	pkg := &p.Packet{}

	fmt.Println("client version")
	pkg = &p.Packet{cp.CLIENT_VERSION, &cp.ClientVersion{}}
	send(conn, pkg)
	time.Sleep(time.Second)
	
	fmt.Println("login")
	pkg = &p.Packet{cp.LOGIN, &cp.Login{"222", "222222"}}
	send(conn, pkg)
	time.Sleep(time.Second)

	// TODO new character

	fmt.Println("start game")
	pkg = &p.Packet{cp.START_GAME, &cp.StartGame{}}
	send(conn, pkg)
	time.Sleep(time.Second)
	//
	//// refine cancel ??
	//
	////// in game
	//
	//// walk
	fmt.Println("walk")
	dir := []cp.Direction{cp.Up, cp.Right, cp.Down, cp.Left}
	for _, d := range dir {
		//pkg = &p.Packet{cp.WALK, &cp.Walk{cp.Up}}
		pkg = &p.Packet{cp.WALK, &cp.Walk{d}}
		send(conn, pkg)
		time.Sleep(time.Second)
	}

	// direction

	// chat
	fmt.Println("chat")
	pkg = &p.Packet{cp.CHAT, &cp.Chat{"测试消息1"}}
	//send(conn, pkg)
	pkg = &p.Packet{cp.CHAT, &cp.Chat{"^测试消息2abc"}}
	send(conn, pkg)
	time.Sleep(time.Second)

	fmt.Println(pkg)
}
