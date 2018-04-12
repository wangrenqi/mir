package main

import (
	"net"
	"sync/atomic"
	"log"
	"errors"
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

func processPacket(conn net.Conn, packetChan chan *p.Packet) {
	// TODO
	// get data from conn

	// convert data to package
	// pkg := p.ToPacket(data)
	// packetChan <- pkg
}

func handleClient(conn net.Conn, env *env.Environ) {
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
	return errors.New("invalid package")
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
