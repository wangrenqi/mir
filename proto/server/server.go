package server

import (
	"mir/orm"
	cm "mir/common"
)

const (
	CONNECTED                 = iota
	CLIENT_VERSION
	DISCONNECT
	KEEPALIVE
	NEW_ACCOUNT
	CHANGE_PASSWORD
	CHANGE_PASSWORD_BANNED
	LOGIN
	LOGIN_BANNED
	LOGIN_SUCCESS
	NEW_CHARACTER
	NEW_CHARACTER_SUCCESS
	DELETE_CHARACTER
	DELETE_CHARACTER_SUCCESS
	START_GAME
	START_GAME_BANNED
	START_GAME_DELAY
	MAP_INFORMATION
	USER_INFORMATION
	USER_LOCATION
	OBJECT_PLAYER
	OBJECT_REMOVE
	OBJECT_TURN
	OBJECT_WALK
	OBJECT_RUN
	CHAT
	OBJECT_CHAT
	NEW_ITEM_INFO
	MOVE_ITEM
	EQUIP_ITEM
	MERGE_ITEM
	REMOVE_ITEM
	REMOVE_SLOT_ITEM
	TAKE_BACK_ITEM
	STORE_ITEM
	SPLIT_ITEM
	SPLIT_ITEM1
	DEPOSIT_REFINE_ITEM
	RETRIEVE_REFINE_ITEM
	REFINE_CANCEL
	REFINE_ITEM
	DEPOSIT_TRADE_ITEM
	RETRIEVE_TRADE_ITEM
	USE_ITEM
	DROP_ITEM
	PLAYER_UPDATE
	PLAYER_INSPECT
	LOGOUT_SUCCESS
	LOGOUT_FAILED
	TIME_OF_DAY
	CHANGE_A_MODE
	CHANGE_P_MODE
	OBJECT_ITEM
	OBJECT_GOLD
	GAINED_ITEM
	GAINED_GOLD
	LOSE_GOLD
	GAINED_CREDIT
	LOSE_CREDIT
	OBJECT_MONSTER
	OBJECT_ATTACK
	STRUCK
	OBJECT_STRUCK
	DAMAGE_INDICATOR
	DURA_CHANGED
	HEALTH_CHANGED
	DELETE_ITEM
	DEATH
	OBJECT_DIED
	COLOUR_CHANGED
	OBJECT_COLOUR_CHANGED
	OBJECT_GUILD_NAME_CHANGED
	GAIN_EXPERIENCE
	LEVEL_CHANGED
	OBJECT_LEVELED
	OBJECT_HARVEST
	OBJECT_HARVESTED
	OBJECT_NPC
	NPC_RESPONSE
	OBJECT_HIDE
	OBJECT_SHOW
	POISONED
	OBJECT_POISONED
	MAP_CHANGED
	OBJECT_TELEPORT_OUT
	OBJECT_TELEPORT_IN
	TELEPORT_IN
	NPC_GOODS
	NPC_SELL
	NPC_REPAIR
	NPC_SREPAIR
	NPC_REFINE
	NPC_CHECK_REFINE
	NPC_COLLECT_REFINE
	NPC_REPLACE_WED_RING
	NPC_STORAGE
	SELL_ITEM
	CRAFT_ITEM
	REPAIR_ITEM
	ITEM_REPAIRED
	NEW_MAGIC
	REMOVE_MAGIC
	MAGIC_LEVELED
	MAGIC
	MAGIC_DELAY
	MAGIC_CAST
	// ObjectMagic
	ObjectEffect
	RangeAttack
	Pushed
	ObjectPushed
	ObjectName
	UserStorage
	SwitchGroup
	DeleteGroup
	DeleteMember
	GroupInvite
	AddMember
	Revived
	ObjectRevived
	SpellToggle
	ObjectHealth
	MapEffect
	ObjectRangeAttack
	AddBuff
	RemoveBuff
	ObjectHidden
	RefreshItem
	ObjectSpell
	UserDash
	ObjectDash
	UserDashFail
	ObjectDashFail
	NPCConsign
	NPCMarket
	NPCMarketPage
	ConsignItem
	MarketFail
	MarketSuccess
	ObjectSitDown
	InTrapRock
	BaseStatsInfo
	UserName
	ChatItemStats
	GuildNoticeChange
	GuildMemberChange
	GuildStatus
	GuildInvite
	GuildExpGain
	GuildNameRequest
	GuildStorageGoldChange
	GuildStorageItemChange
	GuildStorageList
	GuildRequestWar
	DefaultNPC
	NPCUpdate
	NPCImageUpdate
	MarriageRequest
	DivorceRequest
	MentorRequest
	TradeRequest
	TradeAccept
	TradeGold
	TradeItem
	TradeConfirm
	TradeCancel
	MountUpdate
	EquipSlotItem
	FishingUpdate
	ChangeQuest
	CompleteQuest
	ShareQuest
	NewQuestInfo
	GainedQuestItem
	DeleteQuestItem
	CancelReincarnation
	RequestReincarnation
	UserBackStep
	ObjectBackStep
	UserDashAttack
	ObjectDashAttack
	UserAttackMove
	CombineItem
	ItemUpgraded
	SetConcentration
	SetObjectConcentration
	SetElemental
	SetObjectElemental
	RemoveDelayedExplosion
	ObjectDeco
	ObjectSneaking
	ObjectLevelEffects
	SetBindingShot
	SendOutputMessage

	NPCAwakening
	NPCDisassemble
	NPCDowngrade
	NPCReset
	AwakeningNeedMaterials
	AwakeningLockedItem
	Awakening

	ReceiveMail
	MailLockedItem
	MailSendRequest
	MailSent
	ParcelCollected
	MailCost
	ResizeInventory
	ResizeStorage
	NewIntelligentCreature
	UpdateIntelligentCreatureList
	IntelligentCreatureEnableRename
	IntelligentCreaturePickup
	NPCPearlGoods

	TransformUpdate
	FriendUpdate
	LoverUpdate
	MentorUpdate
	GuildBuffList
	NPCRequestInput
	GameShopInfo
	GameShopStock
	Rankings
	Opendoor

	GetRentedItems
	ItemRentalRequest
	ItemRentalFee
	ItemRentalPeriod
	DepositRentalItem
	RetrieveRentalItem
	UpdateRentalItem
	CancelItemRental
	ItemRentalLock
	ItemRentalPartnerLock
	CanConfirmItemRental
	ConfirmItemRental
	NewRecipeInfo
)

type Connected struct{}

func (self *Connected) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(CONNECTED)
	return bytes
}

type ClientVersion struct {
	// 0 wrong version
	// 1 correct version
	Result byte
}

func (self *ClientVersion) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(CLIENT_VERSION)
	bytes = append(bytes, self.Result)
	return bytes
}

type Disconnect struct{}

type NewAccount struct {
	/*
 	* 0: Disabled
 	* 1: Bad AccountID
 	* 2: Bad Password
 	* 3: Bad Email
 	* 4: Bad Name
 	* 5: Bad Question
 	* 6: Bad Answer
 	* 7: Account Exists.
 	* 8: Success
 	*/
	Result byte
}

func (self *NewAccount) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(NEW_ACCOUNT)
	return append(bytes, self.Result)
}

type ChangePassword struct {
	Result byte
}

func (self *ChangePassword) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(NEW_ACCOUNT)
	return append(bytes, self.Result)
}

type ChangePasswordBanned struct {
	//reason string
	//expiryDate time
}

type Login struct {
	/*
	  * 0: Disabled
	  * 1: Bad AccountID
	  * 2: Bad Password
	  * 3: Account Not Exist
	  * 4: Wrong Password
	  */
	Result byte
}

func (self *Login) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(LOGIN)
	return append(bytes, self.Result)
}

type LoginBanned struct {
	//Reason string
	//ExpiryDate time
}

func (self *LoginBanned) ToBytes() []byte {
	return nil
}

type LoginSuccess struct {
	// c#
	// count(int32 4byte) [ index(int32 4byte) name(string) level(int16 2byte) class(1byte) gender(1byte) lastAccess(int64 8byte) ]
	Characters []orm.SelectInfo
}

func (self *LoginSuccess) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(LOGIN_SUCCESS)
	characters := self.Characters
	count := len(characters)
	if count == 0 {
		bytes = append(bytes, []byte{0, 0, 0, 0}...)
	} else {
		countBytes := cm.Uint32ToBytes(uint32(count))
		for _, character := range characters {
			countBytes = append(countBytes, character.ToBytes()...)
		}
		bytes = append(bytes, countBytes...)
	}
	return bytes
}

type NewCharacter struct {
	/*
	  * 0: Disabled.
	  * 1: Bad Character Name
	  * 2: Bad Gender
	  * 3: Bad Class
	  * 4: Max Characters
	  * 5: Character Exists.
	  * */
	Result byte
}

func (self *NewCharacter) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(NEW_CHARACTER)
	return append(bytes, self.Result)
}

type NewCharacterSuccess struct {
	CharInfo orm.SelectInfo
}

func (self *NewCharacterSuccess) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(NEW_CHARACTER_SUCCESS)
	characterInfoBytes := self.CharInfo.ToBytes()
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, characterInfoBytes} {
		result = append(result, r...)
	}
	return result
}

type DeleteCharacter struct {
	Result byte
}

type DeleteCharacterSuccess struct {
	CharacterIndex int
}

type StartGame struct {
	Result     byte
	Resolution int // 分辨率的意思??
}

func (self *StartGame) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(START_GAME)
	result := append(pkgBytes, []byte{4, 0, 4, 0, 0}...)
	return result
}

type StartGameBanned struct {
	Reason string
	//ExpiryDate time date
}

type StartGameDelay struct {
	//Milliseconds c# long
}

type LightSetting byte

const (
	NORMAL  LightSetting = iota
	DAWN
	DAY
	EVENING
	NIGHT
)

type MapInformation struct {
	Index        int
	Filename     string
	Title        string
	MiniMap      int
	BigMap       int
	Music        int
	Lightning    bool
	Fire         bool
	MapDarkLight byte
}

func (self *MapInformation) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(MAP_INFORMATION)
	tmp := []byte{1, 48, 14, 66, 105, 99, 104, 111, 110, 80, 114, 111, 118, 105, 110, 99, 101, 101, 0, 135, 0, 0, 0, 0, 0, 0}
	result := append(pkgBytes, tmp...)
	return result
}

type UserInformation struct {
	ObjectID  int
	RealId    int
	Name      string
	GuildName string
	GuildRank string
	//public Color NameColour; int ?
	Class  cm.MirClass
	Gender cm.MirGender
	Level  int
	//public Point Location; [][]int ?
	Direction cm.MirDirection
	Hair      byte
	HP        int
	MP        int
	// long Experience, MaxExperience;
	// LevelEffects
	// public UserItem[] Inventory, Equipment, QuestInventory;
	// public uint Gold, Credit;

	// public bool HasExpandedStorage;
	// public DateTime ExpandedStorageExpiryTime;

	// public List<ClientMagic> Magics = new List<ClientMagic>();

	// public List<ClientIntelligentCreature> IntelligentCreatures = new List<ClientIntelligentCreature>();//IntelligentCreature
	// public IntelligentCreatureType SummonedCreatureType = IntelligentCreatureType.None;//IntelligentCreature
	// public bool CreatureSummoned;//IntelligentCreature
}

func (self *UserInformation) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(USER_INFORMATION)
	result := append(pkgBytes, []byte{165, 7, 0, 0, 1, 0, 0, 0, 3, 54, 54, 51, 0, 0, 255, 255, 255, 255, 0, 0, 1, 0, 31, 1, 0, 0, 100, 2, 0, 0, 4, 4, 18, 0, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 1, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 99, 0}...)
	return result
}

// ......

//type NewQuestInfo struct {
//
//}
//
//func (self *NewQuestInfo) ToBytes() []byte {
//	return
//}
