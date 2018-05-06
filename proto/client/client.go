package client

import (
	cm "mir/common"
)

const (
	CLIENT_VERSION   = iota
	DISCONNECT
	KEEPALIVE
	NEW_ACCOUNT
	CHANGE_PASSWORD
	LOGIN
	NEW_CHARACTER
	DELETE_CHARACTER
	START_GAME
	LOGOUT
	TURN
	WALK
	RUN
	CHAT

	MOVE_ITEM
	STORE_ITEM
	TAKE_BACK_ITEM
	MERGE_ITEM
	EQUIP_ITEM
	REMOVE_ITEM
	REMOVE_SLOT_ITEM
	SPLIT_ITEM
	USE_ITEM
	DROP_ITEM
	DEPOSIT_REFINE_ITEM
	RETRIEVE_REFINE_ITEM
	REFINE_CANCEL
	REFINE_ITEM
	CHECK_REFINE
	REPLACE_WED_RING
	DEPOSIT_TRADE_ITEM
	RETRIEVE_TRADE_ITEM
	DROP_GOLD
	PICK_UP
	INSPECT
	CHANGE_A_MODE
	CHANGE_P_MODE
	CHANGE_TRADE
	ATTACK
	RANGE_ATTACK
	HARVEST
	CALL_NPC
	TALK_MONSTER_NPC
	BUY_ITEM
	SELL_ITEM
	CRAFT_ITEM
	REPAIR_ITEM
	BUY_ITEM_BACK
	S_REPAIR_ITEM
	MAGIC_KEY
	MAGIC
	SWITCH_GROUP
	ADD_MEMBER
	DELL_MEMBER
	GROUP_INVITE
	TOWN_REVIVE
	SPELL_TOGGLE
	CONSIGN_ITEM
	MARKET_SEARCH
	MARKET_REFRESH
	MARKET_PAGE
	MARKET_BUY
	MARKET_GET_BACK
	REQUEST_USER_NAME
	REQUEST_CHAT_ITEM
	EDIT_GUILD_MEMBER
	EDIT_GUILD_NOTICE
	GUILD_INVITE
	GUILD_NAME_RETURN
	REQUEST_GUILD_INFO
	GUILD_STORAGE_GOLD_CHANGE
	GUILD_STORAGE_ITEM_CHANGE
	GUILD_WAR_RETURN
	MARRIAGE_REQUEST
	MARRIAGE_REPLY
	CHANGE_MARRIAGE
	DIVORCE_REQUEST
	DIVORCE_REPLY
	ADD_MENTOR
	MENTOR_REPLY
	ALLOW_MENTOR
	CANCEL_MENTOR
	TRADE_REQUEST
	TRADE_REPLY
	TRADE_GOLD
	TRADE_CONFIRM
	TRADE_CANCEL
	EQUIP_SLOT_ITEM
	FISHING_CAST
	FISHING_CHANGE_AUTOCAST
	ACCEPT_QUEST
	FINISH_QUEST
	ABANDON_QUEST
	SHARE_QUEST

	ACCEPT_REINCARNATION
	CANCEL_REINCARNATION
	COMBINE_ITEM

	SET_CONCENTRATION
	AWAKENING_NEED_MATERIALS
	AWAKENING_LOCKED_ITEM
	AWAKENING
	DISASSEMBLE_ITEM
	DOWNGRADE_AWAKENING
	RESET_ADDED_ITEM

	SEND_MAIL
	READ_MAIL
	COLLECT_PARCEL
	DELETE_MAIL
	LOCK_MAIL
	MAIL_LOCKED_ITEM
	MAIL_COST

	UPDATE_INTELLIGENT_CREATURE
	INTELLIGENT_CREATURE_PICKUP

	ADD_FRIEND
	REMOVE_FRIEND
	REFRESH_FRIENDS
	ADD_MEMO
	GUILD_BUFF_UPDATE
	NPC_CONFIRM_INPUT
	GAMESHOP_BUY

	REPORT_ISSUE
	GET_RANKING
	OPENDOOR

	GET_RENTED_ITEMS
	ITEM_RENTAL_REQUEST
	ITEM_RENTAL_FEE
	ITEM_RENTAL_PERIOD
	DEPOSIT_RENTAL_ITEM
	RETRIEVE_RENTAL_ITEM
	CANCEL_ITEM_RENTAL
	ITEM_RENTAL_LOCK_FEE
	ITEM_RENTAL_LOCK_ITEM
	CONFIRM_ITEM_RENTAL
)

type ClientVersion struct {
	VersionHash string
}

func GetClientVersion(bytes []byte) *ClientVersion {

	return nil
}

func (self *ClientVersion) ToBytes() []byte {
	//24, 0 (22 + 2)
	return []byte{0, 0, 16, 0, 0, 0, 196, 46, 198, 6, 217, 38, 102, 128, 242, 128, 185, 164, 66, 146, 36, 34}
}

type Disconnect struct{}

func GetDisconnect(bytes []byte) *Disconnect {

	return nil
}

func (self *Disconnect) ToBytes() []byte {
	return nil
}

type KeepAlive struct {
	//time time.Time
}

func GetKeepAlive(bytes []byte) *KeepAlive {

	return nil
}

func (self *KeepAlive) ToBytes() []byte {
	return nil
}

type NewAccount struct {
	AccountID string
	Password  string
	//BirthDate      datetime
	//UserName
	SecretQuestion string
	SecretAnswer   string
	EMailAddress   string
}

func GetNewAccount(bytes []byte) *NewAccount {
	index, accountId := cm.ReadString(bytes, 0)
	index, password := cm.ReadString(bytes, index)
	// TODO birthday datetime from binary int64
	return &NewAccount{AccountID: accountId, Password: password}
}

func (self *NewAccount) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(NEW_ACCOUNT)
	accountIdBytes := cm.StringToBytes(self.AccountID)
	passwordBytes := cm.StringToBytes(self.Password)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, accountIdBytes, passwordBytes} {
		result = append(result, r...)
	}
	return result
}

type ChangePassword struct {
	AccountID       string
	CurrentPassword string
	NewPassword     string
}

func GetChangePassword(bytes []byte) *ChangePassword {

	return nil
}

func (self *ChangePassword) ToBytes() []byte {
	return nil
}

type Login struct {
	AccountID string
	Password  string
}

func GetLogin(bytes []byte) *Login {
	index, accountId := cm.ReadString(bytes, 0)
	index, password := cm.ReadString(bytes, index)
	return &Login{AccountID: accountId, Password: password}
}

func (self *Login) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(LOGIN)
	accountIdBytes := cm.StringToBytes(self.AccountID)
	passwordBytes := cm.StringToBytes(self.Password)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, accountIdBytes, passwordBytes} {
		result = append(result, r...)
	}
	return result
}

type NewCharacter struct {
	Name   string
	Gender cm.MirGender
	Class  cm.MirClass
}

func GetNewCharacter(bytes []byte) *NewCharacter {
	index, name := cm.ReadString(bytes, 0)
	gender := bytes[index]
	class := bytes[index+1]
	return &NewCharacter{Name: name, Gender: cm.MirGender(gender), Class: cm.MirClass(class)}
}

func (self *NewCharacter) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(NEW_CHARACTER)
	nameBytes := cm.StringToBytes(self.Name)
	genderBytes := []byte{byte(self.Gender)}
	classBytes := []byte{byte(self.Class)}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, nameBytes, genderBytes, classBytes} {
		result = append(result, r...)
	}
	return result
}

type DeleteCharacter struct {
	CharacterIndex int
}

func GetDeleteCharacter(bytes []byte) *DeleteCharacter {

	return nil
}

func (self *DeleteCharacter) ToBytes() []byte {
	return nil
}

type StartGame struct {
	CharacterIndex int
}

func GetStartGame(bytes []byte) *StartGame {
	characterIndex := int(cm.BytesToUint32(bytes))
	return &StartGame{CharacterIndex: characterIndex}
}

func (self *StartGame) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(START_GAME)
	characterIndexBytes := cm.Uint32ToBytes(uint32(self.CharacterIndex))
	return append(pkgBytes, characterIndexBytes...)
}

type Logout struct{}

func GetLogout(bytes []byte) *Logout {

	return nil
}

func (self *Logout) ToBytes() []byte {

	return nil
}

type Turn struct {
	Direction cm.MirDirection
}

func GetTurn(bytes []byte) *Turn {

	return nil
}

func (self *Turn) ToBytes() []byte {
	return nil
}

type Walk struct {
	Direction cm.MirDirection
}

func GetWalk(bytes []byte) *Walk {
	if len(bytes) != 1 {
		return nil
	}
	in := false
	for _, dir := range []cm.MirDirection{cm.UP, cm.UP_RIGHT, cm.RIGHT, cm.DOWN_RIGHT, cm.DOWN, cm.DOWN_LEFT, cm.LEFT, cm.UP_LEFT} {
		if cm.MirDirection(bytes[0]) == dir {
			in = true
		}
	}
	if in {
		return &Walk{cm.MirDirection(bytes[0])}
	}
	return nil
}

func (self *Walk) ToBytes() []byte {
	// up upright right downright down downleft left upleft
	// 5, 0 (3 + 2)
	return []byte{11, 0, byte(self.Direction)}
}

type Run struct {
	Direction cm.MirDirection
}

func GetRun(bytes []byte) *Run {

	return nil
}

func (self *Run) ToBytes() []byte {
	return nil
}

type Chat struct {
	Message string
}

func GetChat(bytes []byte) *Chat {
	_, msg := cm.ReadString(bytes, 0)
	return &Chat{Message: msg}
}

func (self *Chat) ToBytes() []byte {
	msgBytes := []byte(self.Message)
	index := cm.Uint16ToBytes(CHAT)
	index = append(index, byte(len(msgBytes)))
	bytes := append(index, msgBytes...)
	return bytes
}

type MoveItem struct {
}

func GetMoveItem(bytes []byte) *MoveItem {
	return nil
}

func (self *MoveItem) ToBytes() []byte {
	return nil
}

type StoreItem struct {
}

func GetStoreItem(bytes []byte) *StoreItem {
	return nil
}

func (self *StoreItem) ToBytes() []byte {
	return nil
}

type TakeBackItem struct {
}

func GetTakeBackItem(bytes []byte) *TakeBackItem {
	return nil
}

func (self *TakeBackItem) ToBytes() []byte {
	return nil
}

type MergeItem struct {
}

func GetMergeItem(bytes []byte) *MergeItem {
	return nil
}

func (self *MergeItem) ToBytes() []byte {
	return nil
}

type EquipItem struct {
}

func GetEquipItem(bytes []byte) *EquipItem {
	return nil
}

func (self *EquipItem) ToBytes() []byte {
	return nil
}

type RemoveItem struct {
}

func GetRemoveItem(bytes []byte) *RemoveItem {
	return nil
}

func (self *RemoveItem) ToBytes() []byte {
	return nil
}

type RemoveSlotItem struct {
}

func GetRemoveSlotItem(bytes []byte) *RemoveSlotItem {
	return nil
}

func (self *RemoveSlotItem) ToBytes() []byte {
	return nil
}

type SplitItem struct {
}

func GetSplitItem(bytes []byte) *SplitItem {
	return nil
}

func (self *SplitItem) ToBytes() []byte {
	return nil
}

type UseItem struct {
}

func GetUseItem(bytes []byte) *UseItem {
	return nil
}

func (self *UseItem) ToBytes() []byte {
	return nil
}

type DropItem struct {
}

func GetDropItem(bytes []byte) *DropItem {
	return nil
}

func (self *DropItem) ToBytes() []byte {
	return nil
}

type DepositRefineItem struct {
}

func GetDepositRefineItem(bytes []byte) *DepositRefineItem {
	return nil
}

func (self *DepositRefineItem) ToBytes() []byte {
	return nil
}

type RetrieveRefineItem struct {
}

func GetRetrieveRefineItem(bytes []byte) *RetrieveRefineItem {
	return nil
}

func (self *RetrieveRefineItem) ToBytes() []byte {
	return nil
}

type RefineCancel struct {
}

func GetRefineCancel(bytes []byte) *RefineCancel {
	return nil
}

func (self *RefineCancel) ToBytes() []byte {
	return nil
}

type RefineItem struct {
}

func GetRefineItem(bytes []byte) *RefineItem {
	return nil
}

func (self *RefineItem) ToBytes() []byte {
	return nil
}

type CheckRefine struct {
}

func GetCheckRefine(bytes []byte) *CheckRefine {
	return nil
}

func (self *CheckRefine) ToBytes() []byte {
	return nil
}

type ReplaceWedRing struct {
}

func GetReplaceWedRing(bytes []byte) *ReplaceWedRing {
	return nil
}

func (self *ReplaceWedRing) ToBytes() []byte {
	return nil
}

type DepositTradeItem struct {
}

func GetDepositTradeItem(bytes []byte) *DepositTradeItem {
	return nil
}

func (self *DepositTradeItem) ToBytes() []byte {
	return nil
}

type RetrieveTradeItem struct {
}

func GetRetrieveTradeItem(bytes []byte) *RetrieveTradeItem {
	return nil
}

func (self *RetrieveTradeItem) ToBytes() []byte {
	return nil
}

type DropGold struct {
}

func GetDropGold(bytes []byte) *DropGold {
	return nil
}

func (self *DropGold) ToBytes() []byte {
	return nil
}

type PickUp struct {
}

func GetPickUp(bytes []byte) *PickUp {
	return nil
}

func (self *PickUp) ToBytes() []byte {
	return nil
}

type Inspect struct {
}

func GetInspect(bytes []byte) *Inspect {
	return nil
}

func (self *Inspect) ToBytes() []byte {
	return nil
}

type ChangeAMode struct {
}

func GetChangeAMode(bytes []byte) *ChangeAMode {
	return nil
}

func (self *ChangeAMode) ToBytes() []byte {
	return nil
}

type ChangePMode struct {
}

func GetChangePMode(bytes []byte) *ChangePMode {
	return nil
}

func (self *ChangePMode) ToBytes() []byte {
	return nil
}

type ChangeTrade struct {
}

func (self *ChangeTrade) ToBytes() []byte { return nil }

type Attack struct {
}

func (self *Attack) ToBytes() []byte { return nil }

type RangeAttack struct {
}

func (self *RangeAttack) ToBytes() []byte { return nil }

type Harvest struct {
}

func (self *Harvest) ToBytes() []byte { return nil }

type CallNPC struct {
}

func (self *CallNPC) ToBytes() []byte { return nil }

type TalkMonsterNPC struct {
}

func (self *TalkMonsterNPC) ToBytes() []byte { return nil }

type BuyItem struct {
}

func (self *BuyItem) ToBytes() []byte { return nil }

type SellItem struct {
}

func (self *SellItem) ToBytes() []byte { return nil }

type CraftItem struct {
}

func (self *CraftItem) ToBytes() []byte { return nil }

type RepairItem struct {
}

func (self *RepairItem) ToBytes() []byte { return nil }

type BuyItemBack struct {
}

func (self *BuyItemBack) ToBytes() []byte { return nil }

type SRepairItem struct {
}

func (self *SRepairItem) ToBytes() []byte { return nil }

type MagicKey struct {
}

func (self *MagicKey) ToBytes() []byte { return nil }

type Magic struct {
}

func (self *Magic) ToBytes() []byte { return nil }
