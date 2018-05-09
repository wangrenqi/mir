package srv

import (
	"net"
	"sync/atomic"
	"log"
	"mir/env"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	"fmt"
	"mir/object"
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

			err := c.process(&p.Packet{index, structData})
			if err != nil {
				log.Printf("client process packet return err: %v\n", err)
			}
		}
	}
}

func (c *client) process(pkg *p.Packet) (err error) {
	if pkg == nil || pkg.Index == -1 {
		return nil
	}
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
	case cp.MOVE_ITEM:
		return c.MoveItem(pkg)
	case cp.STORE_ITEM:
		return c.StoreItem(pkg)
	case cp.TAKE_BACK_ITEM:
		return c.TakeBackItem(pkg)
	case cp.MERGE_ITEM:
		return c.MergeItem(pkg)
	case cp.EQUIP_ITEM:
		return c.EquipItem(pkg)
	case cp.REMOVE_ITEM:
		return c.RemoveItem(pkg)
	case cp.REMOVE_SLOT_ITEM:
		return c.RemoveSlotItem(pkg)
	case cp.SPLIT_ITEM:
		return c.SplitItem(pkg)
	case cp.USE_ITEM:
		return c.UseItem(pkg)
	case cp.DROP_ITEM:
		return c.DropItem(pkg)
	case cp.DEPOSIT_REFINE_ITEM:
		return c.DepositRefineItem(pkg)
	case cp.RETRIEVE_REFINE_ITEM:
		return c.RetrieveRefineItem(pkg)
	case cp.REFINE_CANCEL:
		return c.RefineCancel(pkg)
	case cp.REFINE_ITEM:
		return c.RefineItem(pkg)
	case cp.CHECK_REFINE:
		return c.CheckRefine(pkg)
	case cp.REPLACE_WED_RING:
		return c.ReplaceWedRing(pkg)
	case cp.DEPOSIT_TRADE_ITEM:
		return c.DepositTradeItem(pkg)
	case cp.RETRIEVE_TRADE_ITEM:
		return c.RetrieveTradeItem(pkg)
	case cp.DROP_GOLD:
		return c.DropGold(pkg)
	case cp.PICK_UP:
		return c.PickUp(pkg)
	case cp.INSPECT:
		return c.Inspect(pkg)
	case cp.CHANGE_A_MODE:
		return c.ChangeAMode(pkg)
	case cp.CHANGE_P_MODE:
		return c.ChangePMode(pkg)
	case cp.CHANGE_TRADE:
		return c.ChangeTrade(pkg)
	case cp.ATTACK:
		return c.Attack(pkg)
	case cp.RANGE_ATTACK:
		return c.RangeAttack(pkg)
	case cp.HARVEST:
		return c.Harvest(pkg)
	case cp.CALL_NPC:
		return c.CallNPC(pkg)
	case cp.TALK_MONSTER_NPC:
		return c.TalkMonsterNPC(pkg)
	case cp.BUY_ITEM:
		return c.BuyItem(pkg)
	case cp.SELL_ITEM:
		return c.SellItem(pkg)
	case cp.CRAFT_ITEM:
		return c.CraftItem(pkg)
	case cp.REPAIR_ITEM:
		return c.RepairItem(pkg)
	case cp.BUY_ITEM_BACK:
		return c.BuyItemBack(pkg)
	case cp.S_REPAIR_ITEM:
		return c.SRepairItem(pkg)
	case cp.MAGIC_KEY:
		return c.MagicKey(pkg)
	case cp.MAGIC:
		return c.Magic(pkg)
	}
	return fmt.Errorf("unknow pkg index: %d", pkg.Index)
}
