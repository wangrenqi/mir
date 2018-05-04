package srv

import (
	"mir/orm"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	cm "mir/common"
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

func Broadcast(this *client, pkg Packet) {
	// TODO
	//根据消息类型 如果是全局就走所有人广播
	//否则Area Of Interesting
	clients := GetClients()

	for id, client := range clients {
		if this.id == id {
			continue
		}
		SendTo(client.conn, pkg)
	}
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
	c.info["accountId"] = account.AccountID
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

func (c *client) DeleteCharacter(pkg *p.Packet) error {

	return nil
}

func (c *client) StartGame(pkg *p.Packet) error {
	if c.status != SELECT {
		return nil
	}

	// TODO get player by username and characterIndex
	// characterIndex := pkg.Data.(*cp.StartGame).CharacterIndex
	//c.player = &env.Player{}

	SendTo(c.conn, &sp.StartGame{})
	SendTo(c.conn, &sp.MapInformation{
		Index:        12289,   //uint16 // TODO
		Title:        "比奇省",   //string
		MiniMap:      101,     //uint16 // TODO why?
		BigMap:       135,     //uint16
		Music:        0,       //uint16
		Lightning:    false,   //bool
		Fire:         false,   //bool
		MapDarkLight: byte(0), //byte
	})
	SendTo(c.conn, &sp.UserInformation{
		ObjectID:                  1,                    //uint32
		RealId:                    1,                    //uint32
		Name:                      "测试名字",               //string
		GuildName:                 "测试工会名字",             //string
		GuildRank:                 "测试工会Rank",           //string
		NameColour:                1,                    //uint32
		Class:                     1,                    //cm.MirClass
		Gender:                    1,                    //cm.MirGender
		Level:                     1,                    //uint16
		Location:                  cm.Point{X: 1, Y: 1}, //Point
		Direction:                 1,                    //cm.MirDirection
		Hair:                      1,                    //byte
		HP:                        1,                    //uint16
		MP:                        1,                    //uint16
		Experience:                1,                    //uint64
		MaxExperience:             1,                    //uint64
		LevelEffect:               1,                    //LevelEffects
		Inventory:                 1,                    //interface{} // []UserItem
		Equipment:                 1,                    //interface{} // []UserItem
		QuestInventory:            1,                    //interface{} // []UserItem
		Gold:                      1,                    //uint32
		Credit:                    1,                    //uint32
		HasExpandedStorage:        false,                //bool
		ExpandedStorageExpiryTime: 1,                    //uint64      // DateTime
		Magics:                    1,                    //interface{} // []ClientMagic
		IntelligentCreatures:      1,                    //interface{} // []ClientIntelligentCreature
		IntelligentCreatureType:   1,                    //cm.IntelligentCreatureType
		CreatureSummoned:          false,                //bool
	})
	c.status = GAME
	return nil
}

func (c *client) Logout(pkg *p.Packet) error {

	return nil
}

func (c *client) Turn(pkg *p.Packet) error {
	if c.status != GAME {
		return nil
	}
	Broadcast(c, &sp.ObjectTurn{})
	SendTo(c.conn, &sp.UserLocation{})
	return nil
}

func (c *client) Walk(pkg *p.Packet) error {
	if c.status != GAME {
		return nil
	}
	if !c.player.CanWalk() || !c.player.CanMove() {
		SendTo(c.conn, &sp.UserLocation{c.player.CurrentLocation, c.player.Direction})
	}
	playerMap := (*c.env.Maps)[c.player.CurrentMapIndex]
	targetPoint := c.player.CurrentLocation.Move(c.player.Direction, 1)
	if !playerMap.ValidPoint(targetPoint) {
		SendTo(c.conn, &sp.UserLocation{c.player.CurrentLocation, c.player.Direction})
	}
	// TODO ...剩下的各种判断

	// 广播给附近玩家，在其他client player视角里，本client player 就是object player
	Broadcast(c, &sp.ObjectWalk{})
	return nil
}

func (c *client) Run(pkg *p.Packet) error {
	if c.status != GAME {
		return nil
	}
	if !c.player.CanMove() || !c.player.CanMove() || !c.player.CanRun() {
		SendTo(c.conn, &sp.UserLocation{c.player.CurrentLocation, c.player.Direction})
	}
	SendTo(c.conn, &sp.UserLocation{})
	Broadcast(c, &sp.ObjectRun{})
	return nil
}

func (c *client) Chat(pkg *p.Packet) error {
	if c.status != GAME {
		return nil
	}
	msg := pkg.Data.(*cp.Chat).Message
	log.Println("received client message:", msg)
	return nil
}
