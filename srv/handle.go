package srv

import (
	"net"
	"sync/atomic"
	"mir/env"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	"mir/object"
	"log"
)

type client struct {
	id        int32
	conn      net.Conn
	reqChan   <-chan []byte
	env       *env.Environ
	status    int
	info      map[string]interface{}
	player    *object.PlayerObject
	aoiEntity *env.AOIEntity
}

var id int32 = 0
var clients = make(map[int32]*client)
// TODO map 换成atomic map

func HandleClient(conn net.Conn, env *env.Environ) {
	reqChan := make(chan []byte, 1024)
	client := &client{
		id:        id,
		conn:      conn,
		reqChan:   reqChan,
		env:       env,
		status:    0,
		info:      make(map[string]interface{}),
		player:    nil,
		aoiEntity: nil,
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

			c.process(&p.Packet{index, structData})
		}
	}
}

func (c *client) process(pkg *p.Packet) {
	if pkg == nil || pkg.Index == -1 {
		return
	}
	switch pkg.Index {
	case cp.CLIENT_VERSION:
		c.ClientVersion(pkg)
	case cp.DISCONNECT:
		c.Disconnect(pkg)
	case cp.KEEPALIVE:
		c.Keepalive(pkg)
	case cp.NEW_ACCOUNT:
		c.NewAccount(pkg)
	case cp.CHANGE_PASSWORD:
		c.ChangePassword(pkg)
	case cp.LOGIN:
		c.Login(pkg)
	case cp.NEW_CHARACTER:
		c.NewCharacter(pkg)
	case cp.DELETE_CHARACTER:
		c.DeleteCharacter(pkg)
	case cp.START_GAME:
		c.StartGame(pkg)
	case cp.LOGOUT:
		c.Logout(pkg)
	case cp.TURN:
		c.Turn(pkg)
	case cp.WALK:
		c.Walk(pkg)
	case cp.RUN:
		c.Run(pkg)
	case cp.CHAT:
		c.Chat(pkg)
	case cp.MOVE_ITEM:
		c.MoveItem(pkg)
	case cp.STORE_ITEM:
		c.StoreItem(pkg)
	case cp.TAKE_BACK_ITEM:
		c.TakeBackItem(pkg)
	case cp.MERGE_ITEM:
		c.MergeItem(pkg)
	case cp.EQUIP_ITEM:
		c.EquipItem(pkg)
	case cp.REMOVE_ITEM:
		c.RemoveItem(pkg)
	case cp.REMOVE_SLOT_ITEM:
		c.RemoveSlotItem(pkg)
	case cp.SPLIT_ITEM:
		c.SplitItem(pkg)
	case cp.USE_ITEM:
		c.UseItem(pkg)
	case cp.DROP_ITEM:
		c.DropItem(pkg)
	case cp.DEPOSIT_REFINE_ITEM:
		c.DepositRefineItem(pkg)
	case cp.RETRIEVE_REFINE_ITEM:
		c.RetrieveRefineItem(pkg)
	case cp.REFINE_CANCEL:
		c.RefineCancel(pkg)
	case cp.REFINE_ITEM:
		c.RefineItem(pkg)
	case cp.CHECK_REFINE:
		c.CheckRefine(pkg)
	case cp.REPLACE_WED_RING:
		c.ReplaceWedRing(pkg)
	case cp.DEPOSIT_TRADE_ITEM:
		c.DepositTradeItem(pkg)
	case cp.RETRIEVE_TRADE_ITEM:
		c.RetrieveTradeItem(pkg)
	case cp.DROP_GOLD:
		c.DropGold(pkg)
	case cp.PICK_UP:
		c.PickUp(pkg)
	case cp.INSPECT:
		c.Inspect(pkg)
	case cp.CHANGE_A_MODE:
		c.ChangeAMode(pkg)
	case cp.CHANGE_P_MODE:
		c.ChangePMode(pkg)
	case cp.CHANGE_TRADE:
		c.ChangeTrade(pkg)
	case cp.ATTACK:
		c.Attack(pkg)
	case cp.RANGE_ATTACK:
		c.RangeAttack(pkg)
	case cp.HARVEST:
		c.Harvest(pkg)
	case cp.CALL_NPC:
		c.CallNPC(pkg)
	case cp.TALK_MONSTER_NPC:
		c.TalkMonsterNPC(pkg)
	case cp.BUY_ITEM:
		c.BuyItem(pkg)
	case cp.SELL_ITEM:
		c.SellItem(pkg)
	case cp.CRAFT_ITEM:
		c.CraftItem(pkg)
	case cp.REPAIR_ITEM:
		c.RepairItem(pkg)
	case cp.BUY_ITEM_BACK:
		c.BuyItemBack(pkg)
	case cp.S_REPAIR_ITEM:
		c.SRepairItem(pkg)
	case cp.MAGIC_KEY:
		c.MagicKey(pkg)
	case cp.MAGIC:
		c.Magic(pkg)
	default:
		log.Println("unkonw pkg index:", pkg.Index)
	}
}
