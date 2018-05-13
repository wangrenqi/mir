package srv

import (
	"mir/orm"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	cm "mir/common"
	"log"
	"net"
	"mir/object"
	"mir/env"
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

func sendToAll(this client, pkg Packet) {
	clients := GetClients()
	for id, client := range clients {
		if this.id == id {
			continue
		}
		SendTo(client.conn, pkg)
	}
}

func sendToNearly(this client, pkg Packet) {
	connections := this.aoiEntity.GetNearlyPlayerConnections()
	for _, conn := range connections {
		if this.conn == conn {
			continue
		}
		SendTo(conn, pkg)
	}
}

func Broadcast(this *client, pkg Packet) {
	// TODO
	//根据消息类型 如果是全局sendToAll
	//否则sendToNearly
}

func (c *client) ClientVersion(pkg *p.Packet) {
	// TODO check client version
	SendTo(c.conn, &sp.ClientVersion{Result: byte(1)})
	c.status = LOGIN
}

func (c *client) Disconnect(pkg *p.Packet) {

}
func (c *client) Keepalive(pkg *p.Packet) {

}
func (c *client) NewAccount(pkg *p.Packet) {
	if c.status != LOGIN {
		return
	}
	accountId := pkg.Data.(*cp.NewAccount).AccountID
	password := pkg.Data.(*cp.NewAccount).Password
	var account orm.AccountInfo
	c.env.DB.First(&account, "account_id = ?", accountId)
	if account.AccountID == accountId {
		SendTo(c.conn, &sp.NewAccount{Result: byte(7)})
	}
	c.env.DB.Create(&orm.AccountInfo{
		AccountID: accountId,
		Password:  password,
	})
	SendTo(c.conn, &sp.NewAccount{Result: byte(8)})
}

func (c *client) ChangePassword(packet *p.Packet) {

}

func toSelectInfos(infos []orm.CharacterInfo) []sp.SelectInfo {
	result := make([]sp.SelectInfo, 0)
	for _, i := range infos {
		result = append(result, sp.SelectInfo{
			Index:  i.Index,
			Name:   i.Name,
			Level:  i.Level,
			Class:  cm.MirClass(i.Class),
			Gender: cm.MirGender(i.Gender),
			//LastAccess:i,
		})
	}
	return result
}

func (c *client) Login(pkg *p.Packet) {
	if c.status != LOGIN {
		return
	}
	// check accountId and password
	accountId := pkg.Data.(*cp.Login).AccountID
	password := pkg.Data.(*cp.Login).Password
	var account orm.AccountInfo
	c.env.DB.Where(&orm.AccountInfo{AccountID: accountId, Password: password}).First(&account)
	if account.AccountID == "" {
		// login failed
		SendTo(c.conn, &sp.Login{Result: byte(4)})
		return
	}
	c.status = SELECT
	c.info["AccountID"] = accountId
	c.info["AccountInfoID"] = account.Index
	// query characters
	var characters []orm.CharacterInfo
	c.env.DB.Model(&account).Related(&characters)
	selectInfos := toSelectInfos(characters)
	SendTo(c.conn, &sp.LoginSuccess{Characters: selectInfos})
}

func (c *client) NewCharacter(pkg *p.Packet) {
	if c.status != SELECT {
		return
	}
	name := pkg.Data.(*cp.NewCharacter).Name
	gender := pkg.Data.(*cp.NewCharacter).Gender
	class := pkg.Data.(*cp.NewCharacter).Class
	var character orm.CharacterInfo
	c.env.DB.First(&character, "name = ?", name)
	if character.Name != "" {
		// 已经存在角色名name
		SendTo(c.conn, &sp.NewCharacter{Result: 5})
		return
	}
	// TODO check gender class max...
	startPoint := env.GetStartPoint()
	characterInfo := &orm.CharacterInfo{
		AccountInfoID: c.info["AccountInfoID"].(uint32),
		Name:          name,
		Level:         1,
		Class:         byte(class),
		Gender:        byte(gender),
		Hair:          1,
		//GuildIndex: 1,
		//CreationIP: 1,
		CurrentMapIndex:  0,
		CurrentLocationX: startPoint.X,
		CurrentLocationY: startPoint.Y,
		Direction:        cm.DOWN,
		HP:               1,
		MP:               1,
		Experience:       0,
	}
	c.env.DB.Create(characterInfo)
	SendTo(c.conn, &sp.NewCharacterSuccess{CharInfo: sp.SelectInfo{
		Name:   name,
		Level:  1,
		Class:  class,
		Gender: gender,
		//AccountID: c.info["accountId"].(uint),
	}})
}

func (c *client) DeleteCharacter(pkg *p.Packet) {

}

func (c *client) StartGame(pkg *p.Packet) {
	if c.status != SELECT {
		return
	}

	index := pkg.Data.(*cp.StartGame).CharacterIndex
	accountInfoId := c.info["AccountInfoID"].(uint32)
	var character orm.CharacterInfo
	c.env.DB.Where(&orm.CharacterInfo{Index: uint32(index), AccountInfoID: accountInfoId}).First(&character)
	if character.AccountInfoID == 0 || character.Index == 0 {
		return
	}
	// TODO
	//aoiEntity := env.GetAOIEntity(character.CurrentMapIndex, cm.Point{X: character.CurrentLocationX, Y: character.CurrentLocationY})
	//aoiEntity.Connections[c.id] = c.conn
	//c.aoiEntity = aoiEntity
	c.player = &object.PlayerObject{
		MapObject: object.MapObject{
			ObjectID:        character.Index, // TODO 应该取map object id，随地图object 数量递增
			Name:            character.Name,
			CurrentMapIndex: character.CurrentMapIndex,
			CurrentLocation: cm.Point{X: character.CurrentLocationX, Y: character.CurrentLocationY},
			Level:           character.Level,
			Direction:       character.Direction,
			// TODO
			// ...
		},
		HP: character.HP,
		MP: character.MP,
	}

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
	// TODO
	SendTo(c.conn, &sp.UserInformation{
		ObjectID:                  1,                                                                      //uint32
		RealId:                    character.Index,                                                        //uint32
		Name:                      character.Name,                                                         //string
		GuildName:                 "测试工会名字",                                                               //string
		GuildRank:                 "测试工会Rank",                                                             //string
		NameColour:                1,                                                                      //uint32
		Class:                     cm.MirClass(character.Class),                                           //cm.MirClass
		Gender:                    cm.MirGender(character.Gender),                                         //cm.MirGender
		Level:                     character.Level,                                                        //uint16
		Location:                  cm.Point{X: character.CurrentLocationX, Y: character.CurrentLocationY}, //Point
		Direction:                 1,                                                                      //cm.MirDirection
		Hair:                      1,                                                                      //byte
		HP:                        character.HP,                                                           //uint16
		MP:                        character.MP,                                                           //uint16
		Experience:                1,                                                                      //uint64
		MaxExperience:             1,                                                                      //uint64
		LevelEffect:               1,                                                                      //LevelEffects
		Inventory:                 1,                                                                      //interface{} // []UserItem
		Equipment:                 1,                                                                      //interface{} // []UserItem
		QuestInventory:            1,                                                                      //interface{} // []UserItem
		Gold:                      1,                                                                      //uint32
		Credit:                    1,                                                                      //uint32
		HasExpandedStorage:        false,                                                                  //bool
		ExpandedStorageExpiryTime: 1,                                                                      //uint64      // DateTime
		Magics:                    1,                                                                      //interface{} // []ClientMagic
		IntelligentCreatures:      1,                                                                      //interface{} // []ClientIntelligentCreature
		IntelligentCreatureType:   1,                                                                      //cm.IntelligentCreatureType
		CreatureSummoned:          false,                                                                  //bool
	})
	c.status = GAME
}

func (c *client) Logout(pkg *p.Packet) {

}

func (c *client) Turn(pkg *p.Packet) {
	if c.status != GAME {
		return
	}
	if byte(pkg.Data.(*cp.Turn).Direction) == 100 {
		return
	}
	Broadcast(c, &sp.ObjectTurn{ObjectID: c.player.ObjectID, Direction: pkg.Data.(*cp.Turn).Direction, Location: c.player.CurrentLocation})
	SendTo(c.conn, &sp.UserLocation{Direction: pkg.Data.(*cp.Turn).Direction, Location: c.player.CurrentLocation})
}

func (c *client) Walk(pkg *p.Packet) {
	if c.status != GAME {
		return
	}
	if !c.player.CanWalk() || !c.player.CanMove() {
		SendTo(c.conn, &sp.UserLocation{c.player.CurrentLocation, c.player.Direction})
	}
	playerMap := (*c.env.Maps)[c.player.CurrentMapIndex]
	targetDirection := pkg.Data.(*cp.Walk).Direction
	targetPoint := c.player.CurrentLocation.Move(targetDirection, 1)
	if !playerMap.ValidPoint(targetPoint) {
		SendTo(c.conn, &sp.UserLocation{c.player.CurrentLocation, c.player.Direction})
	}
	// TODO ...剩下的各种判断
	c.player.CurrentLocation = targetPoint
	c.player.Direction = targetDirection
	// 广播给附近玩家，在其他client player视角里，本client player 就是object player
	SendTo(c.conn, &sp.UserLocation{targetPoint, targetDirection})
	Broadcast(c, &sp.ObjectWalk{ObjectID: c.player.ObjectID, Direction: targetDirection, Location: targetPoint})
}

func (c *client) Run(pkg *p.Packet) {
	if c.status != GAME {
		return
	}
	if !c.player.CanMove() || !c.player.CanMove() || !c.player.CanRun() {
		SendTo(c.conn, &sp.UserLocation{c.player.CurrentLocation, c.player.Direction})
	}
	playerMap := (*c.env.Maps)[c.player.CurrentMapIndex]
	playerLocation := c.player.CurrentLocation
	targetDirection := pkg.Data.(*cp.Run).Direction
	targetPoint := c.player.CurrentLocation.Move(targetDirection, 1)
	steps := 2
	for i := 1; i <= steps; i ++ {
		// TODO check point
		targetPoint = c.player.CurrentLocation.Move(targetDirection, 1)
		if !playerMap.ValidPoint(targetPoint) {
			targetPoint = c.player.CurrentLocation
			break
		}
		c.player.CurrentLocation = targetPoint
		c.player.Direction = targetDirection
	}
	if playerLocation != targetPoint {
		SendTo(c.conn, &sp.UserLocation{targetPoint, targetDirection})
		Broadcast(c, &sp.ObjectRun{c.player.ObjectID, targetPoint, targetDirection})
	} else {
		SendTo(c.conn, &sp.UserLocation{playerLocation, targetDirection})
		Broadcast(c, &sp.UserLocation{c.player.CurrentLocation, targetDirection})
	}
}

func (c *client) Chat(pkg *p.Packet) {
	if c.status != GAME {
		return
	}
	msg := pkg.Data.(*cp.Chat).Message
	// TODO switch case msg type
	Broadcast(c, &sp.Chat{Message: msg, Type: cm.CT_NORMAL})
}

func (c *client) MoveItem(pkg *p.Packet)           {}
func (c *client) StoreItem(pkg *p.Packet)          {}
func (c *client) TakeBackItem(pkg *p.Packet)       {}
func (c *client) MergeItem(pkg *p.Packet)          {}
func (c *client) EquipItem(pkg *p.Packet)          {}
func (c *client) RemoveItem(pkg *p.Packet)         {}
func (c *client) RemoveSlotItem(pkg *p.Packet)     {}
func (c *client) SplitItem(pkg *p.Packet)          {}
func (c *client) UseItem(pkg *p.Packet)            {}
func (c *client) DropItem(pkg *p.Packet)           {}
func (c *client) DepositRefineItem(pkg *p.Packet)  {}
func (c *client) RetrieveRefineItem(pkg *p.Packet) {}
func (c *client) RefineCancel(pkg *p.Packet)       {}
func (c *client) RefineItem(pkg *p.Packet)         {}
func (c *client) CheckRefine(pkg *p.Packet)        {}
func (c *client) ReplaceWedRing(pkg *p.Packet)     {}
func (c *client) DepositTradeItem(pkg *p.Packet)   {}
func (c *client) RetrieveTradeItem(pkg *p.Packet)  {}
func (c *client) DropGold(pkg *p.Packet)           {}
func (c *client) PickUp(pkg *p.Packet)             {}
func (c *client) Inspect(pkg *p.Packet)            {}
func (c *client) ChangeAMode(pkg *p.Packet)        {}
func (c *client) ChangePMode(pkg *p.Packet)        {}
func (c *client) ChangeTrade(pkg *p.Packet)        {}
func (c *client) Attack(pkg *p.Packet)             {}
func (c *client) RangeAttack(pkg *p.Packet)        {}
func (c *client) Harvest(pkg *p.Packet)            {}
func (c *client) CallNPC(pkg *p.Packet)            {}
func (c *client) TalkMonsterNPC(pkg *p.Packet)     {}
func (c *client) BuyItem(pkg *p.Packet)            {}
func (c *client) SellItem(pkg *p.Packet)           {}
func (c *client) CraftItem(pkg *p.Packet)          {}
func (c *client) RepairItem(pkg *p.Packet)         {}
func (c *client) BuyItemBack(pkg *p.Packet)        {}
func (c *client) SRepairItem(pkg *p.Packet)        {}
func (c *client) MagicKey(pkg *p.Packet)           {}
func (c *client) Magic(pkg *p.Packet)              {}
