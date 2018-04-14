package server

import (
	"net"
	"sync/atomic"
	"log"
	"mir-go/env"
	"mir-go/orm"
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
		return c.clientVersion(pkg)
	case cp.DISCONNECT:
		return c.disconnect(pkg)
	case cp.KEEPALIVE:
		return c.keepalive(pkg)
	case cp.NEW_ACCOUNT:
		return c.newAccount(pkg)
	}
	return nil
}

func (c *client) clientVersion(pkg *p.Packet) error {

	return nil
}

func (c *client) disconnect(pkg *p.Packet) error {

	return nil
}
func (c *client) keepalive(pkg *p.Packet) error {

	return nil
}
func (c *client) newAccount(pkg *p.Packet) error {
	// TODO check duplicate username
	// if duplicate: return error

	c.env.Db.Create(&orm.Account{
		UserName: pkg.Data.(*cp.NewAccount).UserName,
		Password: pkg.Data.(*cp.NewAccount).Password,
	})

	return nil
}
