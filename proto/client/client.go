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

	return nil
}

func (self *StartGame) ToBytes() []byte {
	// 8, 0 (6 + 2)
	return []byte{8, 0, 2, 0, 0, 0}
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

type StoreItem struct {
}

type TakeBackItem struct {
}

type MergeItem struct {
}

type EquipItem struct {
}

type RemoveItem struct {
}

type RemoveSlotItem struct {
}

type SplitItem struct {
}

type UseItem struct {
}

type DropItem struct {
}

type DepositRefineItem struct {
}

type RetrieveRefineItem struct {
}

type RefineCancel struct {
}

type RefineItem struct {
}

type CheckRefine struct {
}

type ReplaceWedRing struct {
}

type DepositTradeItem struct {
}

type RetrieveTradeItem struct {
}

type DropGold struct {
}

type PickUp struct {
}

type Inspect struct {
}

type ChangeAMode struct {
}

type ChangePMode struct {
}
