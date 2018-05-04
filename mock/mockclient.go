package main

import (
	"net"
	"fmt"
	p "mir/proto"
	cp "mir/proto/client"
	cm "mir/common"
	"time"
)

//var host = "192.168.0.105"
var host = "127.0.0.1"

func send(conn net.Conn, pkg *p.Packet) {
	bytes := pkg.ToBytes()
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
	pkg = &p.Packet{Index: cp.CLIENT_VERSION, Data: &cp.ClientVersion{}}
	send(conn, pkg)
	time.Sleep(time.Second)

	// new account
	//fmt.Println("new account")
	//pkg = &p.Packet{Index: cp.NEW_ACCOUNT, Data: &cp.NewAccount{
	//	AccountID: "111",
	//	Password: "111111",
	//}}
	//send(conn, pkg)
	//time.Sleep(time.Second)

	fmt.Println("login")
	pkg = &p.Packet{Index: cp.LOGIN, Data: &cp.Login{
		AccountID: "111",
		Password:  "111111",
	}}
	send(conn, pkg)
	time.Sleep(time.Second)

	// TODO new character
	fmt.Println("new character")
	pkg = &p.Packet{Index: cp.NEW_CHARACTER, Data: &cp.NewCharacter{
		Name:   "hello",
		Gender: cm.MirGender(1),
		Class:  cm.MirClass(2),
	}}
	send(conn, pkg)
	time.Sleep(time.Second)
	//
	//fmt.Println("start game")
	//pkg = &p.Packet{Index: cp.START_GAME, Data: &cp.StartGame{}}
	//send(conn, pkg)
	//time.Sleep(time.Second)
	//
	//// refine cancel ??
	//
	////// in game
	//
	//// walk
	//fmt.Println("walk")
	//dir := []cm.MirDirection{cm.UP, cm.RIGHT, cm.DOWN, cm.LEFT}
	//for _, d := range dir {
	//	//pkg = &p.Packet{cp.WALK, &cp.Walk{cp.Up}}
	//	pkg = &p.Packet{Index: cp.WALK, Data: &cp.Walk{d}}
	//	send(conn, pkg)
	//	time.Sleep(time.Second)
	//}

	// direction

	// chat
	//fmt.Println("chat")
	//pkg = &p.Packet{Index: cp.CHAT, Data: &cp.Chat{"测试消息1"}}
	////send(conn, pkg)
	//pkg = &p.Packet{Index: cp.CHAT, Data: &cp.Chat{"^测试消息2abc"}}
	//send(conn, pkg)
	//time.Sleep(time.Second)
	//
	//fmt.Println(pkg)
}
