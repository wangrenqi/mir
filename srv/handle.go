package srv

import (
	"net"
	"sync/atomic"
	p "mir/proto"
	cp "mir/proto/client"
	sp "mir/proto/server"
	"log"
	"mir/com"
)

var id int32 = 0
var clients = make(map[int32]*com.Client)
// TODO map 换成atomic map

func GetClients() map[int32]*com.Client {
	return clients
}

func GetClientsAOIs() []*com.AOIEntity {
	visited := make(map[*com.AOIEntity]bool)
	aois := make([]*com.AOIEntity, 0)
	for _, client := range GetClients() {
		aoi := client.AOIEntity
		if visited[aoi] {
			continue
		}
		aois = append(aois, aoi)
	}
	return aois
}

func HandleClient(conn net.Conn, env *com.Environ) {
	reqChan := make(chan []byte, 1024)
	client := &com.Client{
		Id:        id,
		Conn:      conn,
		ReqChan:   reqChan,
		Env:       env,
		Status:    0,
		Info:      make(map[string]interface{}),
		Player:    nil,
		AOIEntity: nil,
	}
	atomic.AddInt32(&id, 1)
	clients[id] = client
	go func() {
		for {
			select {
			case bytes := <-client.ReqChan:
				index, structData := p.BytesToStruct(bytes)

				process(client, &p.Packet{index, structData})
			}
		}
	}()

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

func process(c *com.Client, pkg *p.Packet) {
	if pkg == nil || pkg.Index == -1 {
		return
	}
	switch pkg.Index {
	case cp.CLIENT_VERSION:
		ClientVersion(c, pkg)
	case cp.DISCONNECT:
		Disconnect(c, pkg)
	case cp.KEEPALIVE:
		Keepalive(c, pkg)
	case cp.NEW_ACCOUNT:
		NewAccount(c, pkg)
	case cp.CHANGE_PASSWORD:
		ChangePassword(c, pkg)
	case cp.LOGIN:
		Login(c, pkg)
	case cp.NEW_CHARACTER:
		NewCharacter(c, pkg)
	case cp.DELETE_CHARACTER:
		DeleteCharacter(c, pkg)
	case cp.START_GAME:
		StartGame(c, pkg)
	case cp.LOGOUT:
		Logout(c, pkg)
	case cp.TURN:
		Turn(c, pkg)
	case cp.WALK:
		Walk(c, pkg)
	case cp.RUN:
		Run(c, pkg)
	case cp.CHAT:
		Chat(c, pkg)
	case cp.MOVE_ITEM:
		MoveItem(c, pkg)
	case cp.STORE_ITEM:
		StoreItem(c, pkg)
	case cp.TAKE_BACK_ITEM:
		TakeBackItem(c, pkg)
	case cp.MERGE_ITEM:
		MergeItem(c, pkg)
	case cp.EQUIP_ITEM:
		EquipItem(c, pkg)
	case cp.REMOVE_ITEM:
		RemoveItem(c, pkg)
	case cp.REMOVE_SLOT_ITEM:
		RemoveSlotItem(c, pkg)
	case cp.SPLIT_ITEM:
		SplitItem(c, pkg)
	case cp.USE_ITEM:
		UseItem(c, pkg)
	case cp.DROP_ITEM:
		DropItem(c, pkg)
	case cp.DEPOSIT_REFINE_ITEM:
		DepositRefineItem(c, pkg)
	case cp.RETRIEVE_REFINE_ITEM:
		RetrieveRefineItem(c, pkg)
	case cp.REFINE_CANCEL:
		RefineCancel(c, pkg)
	case cp.REFINE_ITEM:
		RefineItem(c, pkg)
	case cp.CHECK_REFINE:
		CheckRefine(c, pkg)
	case cp.REPLACE_WED_RING:
		ReplaceWedRing(c, pkg)
	case cp.DEPOSIT_TRADE_ITEM:
		DepositTradeItem(c, pkg)
	case cp.RETRIEVE_TRADE_ITEM:
		RetrieveTradeItem(c, pkg)
	case cp.DROP_GOLD:
		DropGold(c, pkg)
	case cp.PICK_UP:
		PickUp(c, pkg)
	case cp.INSPECT:
		Inspect(c, pkg)
	case cp.CHANGE_A_MODE:
		ChangeAMode(c, pkg)
	case cp.CHANGE_P_MODE:
		ChangePMode(c, pkg)
	case cp.CHANGE_TRADE:
		ChangeTrade(c, pkg)
	case cp.ATTACK:
		Attack(c, pkg)
	case cp.RANGE_ATTACK:
		RangeAttack(c, pkg)
	case cp.HARVEST:
		Harvest(c, pkg)
	case cp.CALL_NPC:
		CallNPC(c, pkg)
	case cp.TALK_MONSTER_NPC:
		TalkMonsterNPC(c, pkg)
	case cp.BUY_ITEM:
		BuyItem(c, pkg)
	case cp.SELL_ITEM:
		SellItem(c, pkg)
	case cp.CRAFT_ITEM:
		CraftItem(c, pkg)
	case cp.REPAIR_ITEM:
		RepairItem(c, pkg)
	case cp.BUY_ITEM_BACK:
		BuyItemBack(c, pkg)
	case cp.S_REPAIR_ITEM:
		SRepairItem(c, pkg)
	case cp.MAGIC_KEY:
		MagicKey(c, pkg)
	case cp.MAGIC:
		Magic(c, pkg)
	default:
		log.Println("unkonw pkg index:", pkg.Index)
	}
}
