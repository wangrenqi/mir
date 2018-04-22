package srv

import (
	"mir/orm"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	"log"
	"net"
)

const (
	NONE         = iota
	LOGIN
	SELECT
	GAME
	DISCONNECTED
)

type Packet interface {
	ToBytes() []byte
}

func SendTo(conn net.Conn, pkg Packet) {
	conn.Write(p.Pack(pkg.ToBytes()))
}

func (c *client) ClientVersion(pkg *p.Packet) error {
	// TODO check client version
	SendTo(c.conn, &sp.ClientVersion{Result: byte(1)})
	c.status = LOGIN
	return nil
}

func (c *client) Disconnect(pkg *p.Packet) error {

	return nil
}
func (c *client) Keepalive(pkg *p.Packet) error {

	return nil
}
func (c *client) NewAccount(pkg *p.Packet) error {
	if c.status == LOGIN {
		username := pkg.Data.(*cp.NewAccount).UserName
		password := pkg.Data.(*cp.NewAccount).Password

		var account orm.Account
		c.env.Db.First(&account, "user_name = ?", username)
		if account.UserName == username {
			SendTo(c.conn, &sp.NewAccount{Result: byte(7)})
			return nil
		}
		c.env.Db.Create(&orm.Account{
			UserName: username,
			Password: password,
		})
		SendTo(c.conn, &sp.NewAccount{Result: byte(8)})
	}
	return nil
}

func (c *client) ChangePassword(packet *p.Packet) error {

	return nil
}

func (c *client) Login(packet *p.Packet) error {
	if c.status == LOGIN {
		// check username and password
		c.status = SELECT
	}
	return nil
}

func (c *client) NewCharacter(packet *p.Packet) error {

	return nil
}
func (c *client) DeleteCharacter(packet *p.Packet) error {

	return nil
}
func (c *client) StartGame(packet *p.Packet) error {

	return nil
}
func (c *client) Logout(packet *p.Packet) error {

	return nil
}
func (c *client) Turn(packet *p.Packet) error {

	return nil
}
func (c *client) Walk(packet *p.Packet) error {

	return nil
}
func (c *client) Run(packet *p.Packet) error {
	return nil

}
func (c *client) Chat(packet *p.Packet) error {
	msg := packet.Data.(*cp.Chat).Message
	log.Println("received client message:", msg)
	return nil
}
