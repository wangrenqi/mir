package server

import (
	"net"
	"sync/atomic"
	"log"
	"mir-go/env"
	p "mir-go/proto"
	cp "mir-go/proto/client"
)

type client struct {
	id      int32
	conn    net.Conn
	reqChan <-chan *p.Packet
	env     *env.Environ
}

var id int32 = 0
var packetChan = make(chan *p.Packet)
var buf = make([]byte, 2048)

func splite(bytes []byte) {
	length := p.GetBytesLength(bytes)
	if length < 2 {
		return
	}
	// TODO BUG
	pkg := p.ToPacket(bytes[:length])
	if pkg != nil {
		packetChan <- pkg
	}
	splite(bytes[length:])
}

func ProcessPacket(conn net.Conn) {
	for {
		n, err := conn.Read(buf)
		if n < 4 || err != nil {
			break
		}
		length := p.GetBytesLength(buf[:n])
		if length > len(buf) || length < 2 {
			break
		}
		splite(buf)
	}
}

func HandleClient(conn net.Conn, env *env.Environ) {
	client := &client{
		id:      id,
		conn:    conn,
		reqChan: packetChan,
		env:     env,
	}
	atomic.AddInt32(&id, 1)

	client.run()
}

func (c *client) run() {
	for {
		select {
		case pkg := <-c.reqChan:
			err := c.process(pkg)
			if err != nil {
				log.Printf("client process packet %v return err: %v\n", pkg, err)
			}
		}
	}
}

func (c *client) process(pkg *p.Packet) (err error) {
	if pkg == nil || pkg.Index == -1 {
		return nil
	}
	log.Printf("client packet index: %d", pkg.Index)
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
	return nil
}
