package srv

import (
	"mir/orm"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	"log"
)

func (c *client) ClientVersion(pkg *p.Packet) error {
	// TODO check client version
	c.conn.Write(p.Pack((&sp.ClientVersion{byte(1)}).ToBytes()))
	return nil
}

func (c *client) Disconnect(pkg *p.Packet) error {

	return nil
}
func (c *client) Keepalive(pkg *p.Packet) error {

	return nil
}
func (c *client) NewAccount(pkg *p.Packet) error {
	// TODO check duplicate username
	// if duplicate: return error

	c.env.Db.Create(&orm.Account{
		UserName: pkg.Data.(*cp.NewAccount).UserName,
		Password: pkg.Data.(*cp.NewAccount).Password,
	})

	return nil
}

func (c *client) ChangePassword(packet *p.Packet) error {

	return nil
}

func (c *client) Login(packet *p.Packet) error {
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
