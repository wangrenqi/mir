package srv

import (
	"log"
	"mir/com"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
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

func SendTo(conn net.Conn, pkgs ...interface{}) {
	bytes := make([]byte, 0)
	for _, pkg := range pkgs {
		bytes = append(bytes, p.Pack(pkg.(Packet).ToBytes())...)
	}
	log.Println("send to client bytes:", bytes)
	conn.Write(bytes)
}

func sendToAll(this com.Client, pkg Packet) {
	clients := GetClients()
	for id, client := range clients {
		if this.Id == id {
			continue
		}
		SendTo(client.Conn, pkg)
	}
}

func sendToNearly(this com.Client, pkg Packet) {
	connections := this.AOIEntity.GetNearlyPlayerConnections()
	for _, conn := range connections {
		if this.Conn == conn {
			continue
		}
		SendTo(conn, pkg)
	}
}

func Broadcast(this *com.Client, pkg Packet) {
	// TODO
	//根据消息类型 如果是全局sendToAll
	//否则sendToNearly
}

func ClientVersion(c *com.Client, pkg *p.Packet) {
	// TODO check client version
	SendTo(c.Conn, &sp.ClientVersion{Result: byte(1)})
	c.Status = LOGIN
}

func Disconnect(c *com.Client, pkg *p.Packet) {

}
func Keepalive(c *com.Client, pkg *p.Packet) {

}
func NewAccount(c *com.Client, pkg *p.Packet) {
	if c.Status != LOGIN {
		return
	}
	accountId := pkg.Data.(*cp.NewAccount).AccountID
	password := pkg.Data.(*cp.NewAccount).Password
	var account com.AccountInfo
	c.Env.DB.First(&account, "account_id = ?", accountId)
	if account.AccountID == accountId {
		SendTo(c.Conn, &sp.NewAccount{Result: byte(7)})
	}
	c.Env.DB.Create(&com.AccountInfo{
		AccountID: accountId,
		Password:  password,
	})
	SendTo(c.Conn, &sp.NewAccount{Result: byte(8)})
}

func ChangePassword(c *com.Client, packet *p.Packet) {

}

func toSelectInfos(infos []com.CharacterInfo) []sp.SelectInfo {
	result := make([]sp.SelectInfo, 0)
	for _, i := range infos {
		result = append(result, sp.SelectInfo{
			Index:  i.Index,
			Name:   i.Name,
			Level:  i.Level,
			Class:  com.MirClass(i.Class),
			Gender: com.MirGender(i.Gender),
			//LastAccess:i,
		})
	}
	return result
}

func Login(c *com.Client, pkg *p.Packet) {
	if c.Status != LOGIN {
		return
	}
	// check accountId and password
	accountId := pkg.Data.(*cp.Login).AccountID
	password := pkg.Data.(*cp.Login).Password
	var account com.AccountInfo
	c.Env.DB.Where(&com.AccountInfo{AccountID: accountId, Password: password}).First(&account)
	if account.AccountID == "" {
		// login failed
		SendTo(c.Conn, &sp.Login{Result: byte(4)})
		return
	}
	c.Status = SELECT
	c.Info["AccountID"] = accountId
	c.Info["AccountInfoID"] = account.Index
	// query characters
	var characters []com.CharacterInfo
	c.Env.DB.Model(&account).Related(&characters)
	selectInfos := toSelectInfos(characters)
	SendTo(c.Conn, &sp.LoginSuccess{Characters: selectInfos})
}

func NewCharacter(c *com.Client, pkg *p.Packet) {
	if c.Status != SELECT {
		return
	}
	name := pkg.Data.(*cp.NewCharacter).Name
	gender := pkg.Data.(*cp.NewCharacter).Gender
	class := pkg.Data.(*cp.NewCharacter).Class
	var character com.CharacterInfo
	c.Env.DB.First(&character, "name = ?", name)
	if character.Name != "" {
		// 已经存在角色名name
		SendTo(c.Conn, &sp.NewCharacter{Result: 5})
		return
	}
	// TODO check gender class max...
	startPoint := com.GetStartPoint()
	characterInfo := &com.CharacterInfo{
		AccountInfoID: c.Info["AccountInfoID"].(uint32),
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
		Direction:        com.DOWN,
		HP:               1,
		MP:               1,
		Experience:       0,
	}
	c.Env.DB.Create(characterInfo)
	SendTo(c.Conn, &sp.NewCharacterSuccess{CharInfo: sp.SelectInfo{
		Name:   name,
		Level:  1,
		Class:  class,
		Gender: gender,
		//AccountID: c.info["accountId"].(uint),
	}})
}

func DeleteCharacter(c *com.Client, pkg *p.Packet) {

}

func StartGame(c *com.Client, pkg *p.Packet) {
	if c.Status != SELECT {
		return
	}

	index := pkg.Data.(*cp.StartGame).CharacterIndex
	accountInfoId := c.Info["AccountInfoID"].(uint32)
	var character com.CharacterInfo
	c.Env.DB.Where(&com.CharacterInfo{Index: uint32(index), AccountInfoID: accountInfoId}).First(&character)
	if character.AccountInfoID == 0 || character.Index == 0 {
		return
	}
	aoiEntity := com.GetAOIEntity((*c.Env.AOI)[character.CurrentMapIndex], com.Point{X: character.CurrentLocationX, Y: character.CurrentLocationY})
	if aoiEntity == nil {
		return
	}
	(*aoiEntity.Connections)[c.Id] = c.Conn
	c.AOIEntity = aoiEntity
	c.Player = &com.PlayerObject{
		MapObject: com.MapObject{
			ObjectID:        character.Index, // TODO 应该取map object id，随地图object 数量递增
			Name:            character.Name,
			CurrentMapIndex: character.CurrentMapIndex,
			CurrentLocation: com.Point{X: character.CurrentLocationX, Y: character.CurrentLocationY},
			Level:           character.Level,
			Direction:       character.Direction,
			// TODO
			// ...
		},
		HP: character.HP,
		MP: character.MP,
	}

	SendTo(c.Conn, &sp.StartGame{})
	SendTo(c.Conn, &sp.MapInformation{
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
	itemInfo := com.ItemInfo{
		Index:            uint32(248),                              // uint32
		Name:             "DragonSlayer",                           // string
		Type:             com.ItemType(1),                          // ItemType
		Grade:            com.ItemGrade(2),                         // ItemGrade
		RequiredType:     com.RequiredType(0),                      // RequiredType   // default Level
		RequiredClass:    com.RequiredClass(7),                     // RequiredClass  // default None
		RequiredGender:   com.RequiredGender(3),                    // RequiredGender // default None
		Set:              com.ItemSet(0),                           // ItemSet
		Shape:            com.BytesToUint16([]byte{29, 0}),         // uint16
		Weight:           92,                                       // byte
		Light:            0,                                        // byte
		RequiredAmount:   40,                                       // byte
		Image:            com.BytesToUint16([]byte{57, 0}),         // uint16
		Durability:       com.BytesToUint16([]byte{232, 128}),      // uint16
		Price:            com.BytesToUint32([]byte{1, 0, 0, 0}),    // uint32
		StackSize:        com.BytesToUint32([]byte{248, 36, 1, 0}), // uint32 //default 1;
		MinAC:            0,                                        // byte
		MaxAC:            0,                                        // byte
		MinMAC:           0,                                        // byte
		MaxMAC:           0,                                        // byte
		MinDC:            5,                                        // byte
		MaxDC:            40,                                       // byte
		MinMC:            0,                                        // byte
		MaxMC:            0,                                        // byte
		MinSC:            0,                                        // byte
		MaxSC:            0,                                        // byte
		Accuracy:         0,                                        // byte
		Agility:          0,                                        // byte
		HP:               uint16(0),                                // uint16
		MP:               uint16(0),                                // uint16
		AttackSpeed:      0,                                        // int8 // 需要是负数
		Luck:             0,                                        // int8
		BagWeight:        0,                                        // byte
		HandWeight:       0,                                        // byte
		WearWeight:       0,                                        // byte
		StartItem:        false,                                    // bool
		Effect:           0,                                        // byte
		Strong:           0,                                        // byte
		MagicResist:      0,                                        // byte
		PoisonResist:     0,                                        // byte
		HealthRecovery:   0,                                        // byte
		SpellRecovery:    0,                                        // byte
		PoisonRecovery:   0,                                        // byte
		HPrate:           0,                                        // byte
		MPrate:           0,                                        // byte
		CriticalRate:     0,                                        // byte
		CriticalDamage:   0,                                        // byte
		NeedIdentify:     false,                                    // bool
		ShowGroupPickup:  false,                                    // bool
		GlobalDropNotify: false,                                    // bool
		ClassBased:       false,                                    // bool
		LevelBased:       false,                                    // bool
		CanMine:          false,                                    // bool
		CanFastRun:       false,                                    // bool
		CanAwakening:     false,                                    // bool
		MaxAcRate:        0,                                        // byte
		MaxMacRate:       0,                                        // byte
		Holy:             0,                                        // byte
		Freezing:         0,                                        // byte
		PoisonAttack:     0,                                        // byte
		HpDrainRate:      0,                                        // byte
		Bind:             uint16(0),                                // uint16 // BindMode 这个枚举太大了，直接用uint16 // default none
		Reflect:          0,                                        // byte
		Unique:           uint16(0),                                // uint16 // SpecialItemMode ?? // default None;
		RandomStatsId:    0,                                        // byte
		//RandomStats      : , // RandomItemStat
		//ToolTip          : , // string //default ""
	}
	SendTo(c.Conn, &sp.NewItemInfo{Info: itemInfo})

	// TODO
	userInformation := &sp.UserInformation{
		ObjectID:                  1,                                                                       //uint32
		RealId:                    character.Index,                                                         //uint32
		Name:                      character.Name,                                                          //string
		GuildName:                 "测试工会名字",                                                                //string
		GuildRank:                 "测试工会Rank",                                                              //string
		NameColour:                1,                                                                       //uint32
		Class:                     com.MirClass(character.Class),                                           //com.MirClass
		Gender:                    com.MirGender(character.Gender),                                         //com.MirGender
		Level:                     character.Level,                                                         //uint16
		Location:                  com.Point{X: character.CurrentLocationX, Y: character.CurrentLocationY}, //Point
		Direction:                 1,                                                                       //com.MirDirection
		Hair:                      1,                                                                       //byte
		HP:                        character.HP,                                                            //uint16
		MP:                        character.MP,                                                            //uint16
		Experience:                1,                                                                       //uint64
		MaxExperience:             1,                                                                       //uint64
		LevelEffect:               1,                                                                       //LevelEffects

		Gold:                      666,                                                                     //uint32
		Credit:                    555,                                                                     //uint32
		HasExpandedStorage:        false,                                                                   //bool
		ExpandedStorageExpiryTime: 1,                                                                       //uint64      // DateTime
		Magics:                    1,                                                                       //interface{} // []ClientMagic
		IntelligentCreatures:      1,                                                                       //interface{} // []ClientIntelligentCreature
		IntelligentCreatureType:   1,                                                                       //com.IntelligentCreatureType
		CreatureSummoned:          false,                                                                   //bool
	}
	inventory := make([]com.UserItem, 0)
	inventory = append(inventory, com.UserItem{
		UniqueID:       132,                                            // uint64
		ItemIndex:      248,                                           // uint32
		CurrentDura:    com.BytesToUint16([]byte{152, 95}),            // uint16
		MaxDura:        com.BytesToUint16([]byte{232, 128}),           // uint16
		Count:          uint32(1),                                     // uint32
		AC:             0,                                             // byte
		MAC:            0,                                             // byte
		DC:             0,                                             // byte
		MC:             0,                                             // byte
		SC:             0,                                             // byte
		Accuracy:       0,                                             // byte
		Agility:        0,                                             // byte
		HP:             0,                                             // byte
		MP:             0,                                             // byte
		AttackSpeed:    0,                                             // byte
		Luck:           0,                                             // byte
		SoulBoundId:    com.BytesToUint32([]byte{255, 255, 255, 255}), // uint32
		Bools:          1,                                             // byte
		Identified:     false,                                         // bool
		Cursed:         false,                                         // bool
		Strong:         0,                                             // byte
		MagicResist:    0,                                             // byte
		PoisonResist:   0,                                             // byte
		HealthRecovery: 0,                                             // byte
		ManaRecovery:   0,                                             // byte
		PoisonRecovery: 0,                                             // byte
		CriticalRate:   0,                                             // byte
		CriticalDamage: 0,                                             // byte
		Freezing:       0,                                             // byte
		PoisonAttack:   0,                                             // byte
	})
	userInformation.Inventory = inventory
	userInformation.Equipment = make([]com.UserItem, 0)
	userInformation.QuestInventory = make([]com.UserItem, 0)
	SendTo(c.Conn, userInformation)

	aois := make([]com.AOIEntity, 0)
	aois = append(aois, *c.AOIEntity)
	aois = append(aois, c.AOIEntity.GetNearlyEightAOIs()...)
	allMonsters := make([]interface{}, 0)
	for _, aoi := range aois {
		monsterObjects := aoi.GetMonsterObjects()
		for _, m := range monsterObjects {
			objectMonster := sp.MonsterObjectToObjectMonster(m)
			allMonsters = append(allMonsters, objectMonster)
		}
	}
	SendTo(c.Conn, allMonsters...)

	c.Status = GAME
}

func Logout(c *com.Client, pkg *p.Packet) {

}

func Turn(c *com.Client, pkg *p.Packet) {
	if c.Status != GAME {
		return
	}
	if byte(pkg.Data.(*cp.Turn).Direction) == 100 {
		return
	}
	Broadcast(c, &sp.ObjectTurn{ObjectID: c.Player.ObjectID, Direction: pkg.Data.(*cp.Turn).Direction, Location: c.Player.CurrentLocation})
	SendTo(c.Conn, &sp.UserLocation{Direction: pkg.Data.(*cp.Turn).Direction, Location: c.Player.CurrentLocation})
}

func Walk(c *com.Client, pkg *p.Packet) {
	if c.Status != GAME {
		return
	}
	if !c.Player.CanWalk() || !c.Player.CanMove() {
		SendTo(c.Conn, &sp.UserLocation{c.Player.CurrentLocation, c.Player.Direction})
		return
	}
	playerMap := (*c.Env.Maps)[c.Player.CurrentMapIndex]
	targetDirection := pkg.Data.(*cp.Walk).Direction
	targetPoint := c.Player.CurrentLocation.Move(targetDirection, 1)
	if !c.AOIEntity.ValidPoint(playerMap, targetPoint) {
		SendTo(c.Conn, &sp.UserLocation{c.Player.CurrentLocation, targetDirection})
		Broadcast(c, &sp.ObjectTurn{c.Player.ObjectID, c.Player.CurrentLocation, targetDirection})
		return
	}
	// TODO ...剩下的各种判断
	c.Player.CurrentLocation = targetPoint
	c.Player.Direction = targetDirection
	// 广播给附近玩家，在其他client player视角里，本client player 就是object player
	SendTo(c.Conn, &sp.UserLocation{targetPoint, targetDirection})
	Broadcast(c, &sp.ObjectWalk{ObjectID: c.Player.ObjectID, Direction: targetDirection, Location: targetPoint})
}

func Run(c *com.Client, pkg *p.Packet) {
	if c.Status != GAME {
		return
	}
	if !c.Player.CanMove() || !c.Player.CanMove() || !c.Player.CanRun() {
		SendTo(c.Conn, &sp.UserLocation{c.Player.CurrentLocation, c.Player.Direction})
		return
	}
	playerMap := (*c.Env.Maps)[c.Player.CurrentMapIndex]
	playerLocation := c.Player.CurrentLocation
	targetDirection := pkg.Data.(*cp.Run).Direction
	targetPoint := c.Player.CurrentLocation.Move(targetDirection, 1)
	steps := 2
	for i := 1; i <= steps; i++ {
		// TODO check point
		targetPoint = c.Player.CurrentLocation.Move(targetDirection, 1)
		if !c.AOIEntity.ValidPoint(playerMap, targetPoint) {
			targetPoint = c.Player.CurrentLocation
			break
		}
		c.Player.CurrentLocation = targetPoint
		c.Player.Direction = targetDirection
	}
	if playerLocation != targetPoint {
		SendTo(c.Conn, &sp.UserLocation{targetPoint, targetDirection})
		Broadcast(c, &sp.ObjectRun{c.Player.ObjectID, targetPoint, targetDirection})
	} else {
		SendTo(c.Conn, &sp.UserLocation{playerLocation, targetDirection})
		Broadcast(c, &sp.UserLocation{c.Player.CurrentLocation, targetDirection})
	}
}

func Chat(c *com.Client, pkg *p.Packet) {
	if c.Status != GAME {
		return
	}
	msg := pkg.Data.(*cp.Chat).Message
	// TODO switch case msg type
	Broadcast(c, &sp.Chat{Message: msg, Type: com.CT_NORMAL})
}

func MoveItem(c *com.Client, pkg *p.Packet)           {}
func StoreItem(c *com.Client, pkg *p.Packet)          {}
func TakeBackItem(c *com.Client, pkg *p.Packet)       {}
func MergeItem(c *com.Client, pkg *p.Packet)          {}
func EquipItem(c *com.Client, pkg *p.Packet)          {}
func RemoveItem(c *com.Client, pkg *p.Packet)         {}
func RemoveSlotItem(c *com.Client, pkg *p.Packet)     {}
func SplitItem(c *com.Client, pkg *p.Packet)          {}
func UseItem(c *com.Client, pkg *p.Packet)            {}
func DropItem(c *com.Client, pkg *p.Packet)           {}
func DepositRefineItem(c *com.Client, pkg *p.Packet)  {}
func RetrieveRefineItem(c *com.Client, pkg *p.Packet) {}
func RefineCancel(c *com.Client, pkg *p.Packet)       {}
func RefineItem(c *com.Client, pkg *p.Packet)         {}
func CheckRefine(c *com.Client, pkg *p.Packet)        {}
func ReplaceWedRing(c *com.Client, pkg *p.Packet)     {}
func DepositTradeItem(c *com.Client, pkg *p.Packet)   {}
func RetrieveTradeItem(c *com.Client, pkg *p.Packet)  {}
func DropGold(c *com.Client, pkg *p.Packet)           {}
func PickUp(c *com.Client, pkg *p.Packet)             {}
func Inspect(c *com.Client, pkg *p.Packet)            {}
func ChangeAMode(c *com.Client, pkg *p.Packet)        {}
func ChangePMode(c *com.Client, pkg *p.Packet)        {}
func ChangeTrade(c *com.Client, pkg *p.Packet)        {}
func Attack(c *com.Client, pkg *p.Packet)             {}
func RangeAttack(c *com.Client, pkg *p.Packet)        {}
func Harvest(c *com.Client, pkg *p.Packet)            {}
func CallNPC(c *com.Client, pkg *p.Packet)            {}
func TalkMonsterNPC(c *com.Client, pkg *p.Packet)     {}
func BuyItem(c *com.Client, pkg *p.Packet)            {}
func SellItem(c *com.Client, pkg *p.Packet)           {}
func CraftItem(c *com.Client, pkg *p.Packet)          {}
func RepairItem(c *com.Client, pkg *p.Packet)         {}
func BuyItemBack(c *com.Client, pkg *p.Packet)        {}
func SRepairItem(c *com.Client, pkg *p.Packet)        {}
func MagicKey(c *com.Client, pkg *p.Packet)           {}
func Magic(c *com.Client, pkg *p.Packet)              {}
