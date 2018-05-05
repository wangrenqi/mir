package proto

import "encoding/binary"

import (
	"log"
	cm "mir/common"
	cp "mir/proto/client"
)

type Packet struct {
	Index int
	Data  interface{}
}

type Null struct{}

func (self *Null) ToBytes() []byte {
	return make([]byte, 0)
}

func BytesToStruct(bytes []byte) (int, interface{}) {
	var data interface{}
	index := int(binary.LittleEndian.Uint16(bytes[0:2]))
	log.Println("BytesToStruct bytes ->", bytes, "index ->", index)
	bytes = bytes[2:]
	switch index {
	case cp.CLIENT_VERSION:
		data = cp.GetClientVersion(bytes)
	case cp.DISCONNECT:
		data = cp.GetDisconnect(bytes)
	case cp.KEEPALIVE:
		data = cp.GetKeepAlive(bytes)
	case cp.NEW_ACCOUNT:
		data = cp.GetNewAccount(bytes)
	case cp.CHANGE_PASSWORD:
		data = cp.GetChangePassword(bytes)
	case cp.LOGIN:
		data = cp.GetLogin(bytes)
	case cp.NEW_CHARACTER:
		data = cp.GetNewCharacter(bytes)
	case cp.DELETE_CHARACTER:
		data = cp.GetDeleteCharacter(bytes)
	case cp.START_GAME:
		data = cp.GetStartGame(bytes)
	case cp.LOGOUT:
		data = cp.GetLogout(bytes)
	case cp.TURN:
		data = cp.GetTurn(bytes)
	case cp.WALK:
		data = cp.GetWalk(bytes)
	case cp.RUN:
		data = cp.GetRun(bytes)
	case cp.CHAT:
		data = cp.GetChat(bytes)
	case cp.MOVE_ITEM:
		data = cp.GetMoveItem(bytes)
	case cp.STORE_ITEM:
		data = cp.GetStoreItem(bytes)
	case cp.TAKE_BACK_ITEM:
		data = cp.GetTakeBackItem(bytes)
	case cp.MERGE_ITEM:
		data = cp.GetMergeItem(bytes)
	case cp.EQUIP_ITEM:
		data = cp.GetEquipItem(bytes)
	case cp.REMOVE_ITEM:
		data = cp.GetRemoveItem(bytes)
	case cp.REMOVE_SLOT_ITEM:
		data = cp.GetRemoveSlotItem(bytes)
	case cp.SPLIT_ITEM:
		data = cp.GetSplitItem(bytes)
	case cp.USE_ITEM:
		data = cp.GetUseItem(bytes)
	case cp.DROP_ITEM:
		data = cp.GetDropItem(bytes)
	case cp.DEPOSIT_REFINE_ITEM:
		data = cp.GetDepositRefineItem(bytes)
	case cp.RETRIEVE_REFINE_ITEM:
		data = cp.GetRetrieveRefineItem(bytes)
	case cp.REFINE_CANCEL:
		data = cp.GetRefineCancel(bytes)
	case cp.REFINE_ITEM:
		data = cp.GetRefineItem(bytes)
	case cp.CHECK_REFINE:
		data = cp.GetCheckRefine(bytes)
	case cp.REPLACE_WED_RING:
		data = cp.GetReplaceWedRing(bytes)
	case cp.DEPOSIT_TRADE_ITEM:
		data = cp.GetDepositTradeItem(bytes)
	case cp.RETRIEVE_TRADE_ITEM:
		data = cp.GetRetrieveTradeItem(bytes)
	case cp.DROP_GOLD:
		data = cp.GetDropGold(bytes)
	case cp.PICK_UP:
		data = cp.GetPickUp(bytes)
	case cp.INSPECT:
		data = cp.GetInspect(bytes)
	case cp.CHANGE_A_MODE:
		data = cp.GetChangeAMode(bytes)
	case cp.CHANGE_P_MODE:
		data = cp.GetChangePMode(bytes)
	default:
		data = &Null{}
		log.Println("未知的client packet index:", index)
	}
	return index, data
}

type Parser interface {
	ToBytes() []byte
}

func (pkg *Packet) ToBytes() []byte {
	var parser Parser
	switch pkg.Index {
	case cp.CLIENT_VERSION:
		parser = pkg.Data.(*cp.ClientVersion)
	case cp.DISCONNECT:
		parser = pkg.Data.(*cp.Disconnect)
	case cp.KEEPALIVE:
		parser = pkg.Data.(*cp.KeepAlive)
	case cp.NEW_ACCOUNT:
		parser = pkg.Data.(*cp.NewAccount)
	case cp.CHANGE_PASSWORD:
		parser = pkg.Data.(*cp.ChangePassword)
	case cp.LOGIN:
		parser = pkg.Data.(*cp.Login)
	case cp.NEW_CHARACTER:
		parser = pkg.Data.(*cp.NewCharacter)
	case cp.DELETE_CHARACTER:
		parser = pkg.Data.(*cp.DeleteCharacter)
	case cp.START_GAME:
		parser = pkg.Data.(*cp.StartGame)
	case cp.LOGOUT:
		parser = pkg.Data.(*cp.Logout)
	case cp.TURN:
		parser = pkg.Data.(*cp.Turn)
	case cp.WALK:
		parser = pkg.Data.(*cp.Walk)
	case cp.RUN:
		parser = pkg.Data.(*cp.Run)
	case cp.CHAT:
		parser = pkg.Data.(*cp.Chat)
	case cp.MOVE_ITEM:
		parser = pkg.Data.(*cp.MoveItem)
	case cp.STORE_ITEM:
		parser = pkg.Data.(*cp.StoreItem)
	case cp.TAKE_BACK_ITEM:
		parser = pkg.Data.(*cp.TakeBackItem)
	case cp.MERGE_ITEM:
		parser = pkg.Data.(*cp.MergeItem)
	case cp.EQUIP_ITEM:
		parser = pkg.Data.(*cp.EquipItem)
	case cp.REMOVE_ITEM:
		parser = pkg.Data.(*cp.RemoveItem)
	case cp.REMOVE_SLOT_ITEM:
		parser = pkg.Data.(*cp.RemoveSlotItem)
	case cp.SPLIT_ITEM:
		parser = pkg.Data.(*cp.SplitItem)
	case cp.USE_ITEM:
		parser = pkg.Data.(*cp.UseItem)
	case cp.DROP_ITEM:
		parser = pkg.Data.(*cp.DropItem)
	case cp.DEPOSIT_REFINE_ITEM:
		parser = pkg.Data.(*cp.DepositRefineItem)
	case cp.RETRIEVE_REFINE_ITEM:
		parser = pkg.Data.(*cp.RetrieveRefineItem)
	case cp.REFINE_CANCEL:
		parser = pkg.Data.(*cp.RefineCancel)
	case cp.REFINE_ITEM:
		parser = pkg.Data.(*cp.RefineItem)
	case cp.CHECK_REFINE:
		parser = pkg.Data.(*cp.CheckRefine)
	case cp.REPLACE_WED_RING:
		parser = pkg.Data.(*cp.ReplaceWedRing)
	case cp.DEPOSIT_TRADE_ITEM:
		parser = pkg.Data.(*cp.DepositTradeItem)
	case cp.RETRIEVE_TRADE_ITEM:
		parser = pkg.Data.(*cp.RetrieveTradeItem)
	case cp.DROP_GOLD:
		parser = pkg.Data.(*cp.DropGold)
	case cp.PICK_UP:
		parser = pkg.Data.(*cp.PickUp)
	case cp.INSPECT:
		parser = pkg.Data.(*cp.Inspect)
	case cp.CHANGE_A_MODE:
		parser = pkg.Data.(*cp.ChangeAMode)
	case cp.CHANGE_P_MODE:
		parser = pkg.Data.(*cp.ChangePMode)
	default:
		parser = &Null{}
	}
	return parser.ToBytes()
}

// 封包
func Pack(data []byte) []byte {
	length := len(data) + 2
	lenBytes := cm.Uint16ToBytes(uint16(length))
	return append(lenBytes, data...)
}

// 解包
func UnPack(buffer []byte, readerChan chan []byte) []byte {
	bufLen := len(buffer)

	var i int
	for i = 0; i < bufLen; i = i + 1 {
		if bufLen < 4 {
			break
		}
		dataLen := int(buffer[i+1]<<8 + buffer[i])
		readerChan <- buffer[2+i : dataLen+i]
		i = i + dataLen - 1
	}
	if i >= bufLen {
		return make([]byte, 0)
	}
	return buffer[i:]
}
