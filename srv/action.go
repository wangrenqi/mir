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
	bytes := p.Pack(pkg.ToBytes())
	log.Println("send to client bytes:", bytes)
	conn.Write(bytes)
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
	if c.status != LOGIN {
		return nil
	}
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
	return nil
}

func (c *client) ChangePassword(packet *p.Packet) error {

	return nil
}

func (c *client) Login(pkg *p.Packet) error {
	if c.status != LOGIN {
		return nil
	}
	// check username and password
	username := pkg.Data.(*cp.Login).AccountId
	password := pkg.Data.(*cp.Login).Password
	var account orm.Account
	c.env.Db.Where(&orm.Account{UserName: username, Password: password}).First(&account)
	if account.UserName == "" {
		// login failed
		SendTo(c.conn, &sp.Login{Result: byte(4)})
		return nil
	}
	c.status = SELECT
	c.info["username"] = username
	c.info["accountId"] = account.ID
	// query characters
	var characters []orm.SelectInfo
	c.env.Db.Model(&account).Related(&characters)
	SendTo(c.conn, &sp.LoginSuccess{Characters: characters})
	return nil
}

func (c *client) NewCharacter(pkg *p.Packet) error {
	if c.status != SELECT {
		return nil
	}
	name := pkg.Data.(*cp.NewCharacter).Name
	gender := pkg.Data.(*cp.NewCharacter).Gender
	class := pkg.Data.(*cp.NewCharacter).Class
	var character orm.SelectInfo
	c.env.Db.First(&character, "name = ?", name)
	if character.Name != "" {
		// 已经存在角色名name
		SendTo(c.conn, &sp.NewCharacter{Result: 5})
		return nil
	}
	// TODO check gender class max...
	characterInfo := &orm.SelectInfo{
		Name:   name,
		Level:  1,
		Class:  class,
		Gender: gender,
		//LastAccess int64
		AccountID: c.info["accountId"].(uint),
	}
	c.env.Db.Create(characterInfo)
	SendTo(c.conn, &sp.NewCharacterSuccess{CharInfo: *characterInfo})
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
