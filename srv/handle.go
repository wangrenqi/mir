package srv

import (
	"net"
	"sync/atomic"
	"log"
	"mir/engine"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	"fmt"
)

type client struct {
	id      int32
	conn    net.Conn
	reqChan <-chan []byte
	env     *engine.Environ
	status  int
	info    map[string]interface{}
	player  *engine.PlayerObject
}

var id int32 = 0
var clients = make(map[int32]*client)
// TODO map 换成atomic map

func HandleClient(conn net.Conn, env *engine.Environ) {
	reqChan := make(chan []byte, 1024)
	client := &client{
		id:      id,
		conn:    conn,
		reqChan: reqChan,
		env:     env,
		status:  0,
		info:    make(map[string]interface{}),
		player:  nil,
	}
	atomic.AddInt32(&id, 1)
	clients[id] = client
	go client.run()

	conn.Write(p.Pack((&sp.Connected{}).ToBytes()))

	tmpBuffer := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		tmpBuffer = p.UnPack(append(tmpBuffer, buffer[:n]...), reqChan)
	}
}

func GetClients() map[int32]*client {
	return clients
}

func (c *client) run() {
	for {
		select {
		case bytes := <-c.reqChan:
			index, structData := p.BytesToStruct(bytes)

			err := c.process(&p.Packet{index, structData})
			if err != nil {
				log.Printf("client process packet return err: %v\n", err)
			}
		}
	}
}

func (c *client) process(pkg *p.Packet) (err error) {
	if pkg == nil || pkg.Index == -1 {
		return nil
	}
	switch pkg.Index {
	case cp.CLIENT_VERSION:
		return c.ClientVersion(pkg)
	case cp.DISCONNECT:
		return c.Disconnect(pkg)
	case cp.KEEPALIVE:
		return c.Keepalive(pkg)
	case cp.NEW_ACCOUNT:
		return c.NewAccount(pkg)
	case cp.CHANGE_PASSWORD:
		return c.ChangePassword(pkg)
	case cp.LOGIN:
		return c.Login(pkg)
	case cp.NEW_CHARACTER:
		return c.NewCharacter(pkg)
	case cp.DELETE_CHARACTER:
		return c.DeleteCharacter(pkg)
	case cp.START_GAME:
		return c.StartGame(pkg)
	case cp.LOGOUT:
		return c.Logout(pkg)
	case cp.TURN:
		return c.Turn(pkg)
	case cp.WALK:
		return c.Walk(pkg)
	case cp.RUN:
		return c.Run(pkg)
	case cp.CHAT:
		return c.Chat(pkg)
	}
	return fmt.Errorf("unknow pkg index: %d", pkg.Index)
}
