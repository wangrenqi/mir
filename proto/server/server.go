package server

import (
	cm "mir/common"
	"mir/orm"
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
	NPC_S_REPAIR
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
	OBJECT_MAGIC
	OBJECT_EFFECT
	RANGE_ATTACK
	PUSHED
	OBJECT_PUSHED
	OBJECT_NAME
	USER_STORAGE
	SWITCH_GROUP
	DELETE_GROUP
	DELETE_MEMBER
	GROUP_INVITE
	ADD_MEMBER
	REVIVED
	OBJECT_REVIVED
	SPELL_TOGGLE
	OBJECT_HEALTH
	MAP_EFFECT
	OBJECT_RANGE_ATTACK
	ADD_BUFF
	REMOVE_BUFF
	OBJECT_HIDDEN
	REFRESH_ITEM
	OBJECT_SPELL
	USER_DASH
	OBJECT_DASH
	USER_DASH_FAIL
	OBJECT_DASH_FAIL
	NPC_CONSIGN
	NPC_MARKET
	NPC_MARKET_PAGE
	CONSIGN_ITEM
	MARKET_FAIL
	MARKET_SUCCESS
	OBJECT_SIT_DOWN
	IN_TRAP_ROCK
	BASE_STATS_INFO
	USER_NAME
	CHAT_ITEM_STATS
	GUILD_NOTICE_CHANGE
	GUILD_MEMBER_CHANGE
	GUILD_STATUS
	GUILD_INVITE
	GUILD_EXP_GAIN
	GUILD_NAME_REQUEST
	GUILD_STORAGE_GOLD_CHANGE
	GUILD_STORAGE_ITEM_CHANGE
	GUILD_STORAGE_LIST
	GUILD_REQUEST_WAR
	DEFAULT_NPC
	NPC_UPDATE
	NPC_IMAGE_UPDATE
	MARRIAGE_REQUEST
	DIVORCE_REQUEST
	MENTOR_REQUEST
	TRADE_REQUEST
	TRADE_ACCEPT
	TRADE_GOLD
	TRADE_ITEM
	TRADE_CONFIRM
	TRADE_CANCEL
	MOUNT_UPDATE
	EQUIP_SLOT_ITEM
	FISHING_UPDATE
	CHANGE_QUEST
	COMPLETE_QUEST
	SHARE_QUEST
	NEW_QUEST_INFO
	GAINED_QUEST_ITEM
	DELETE_QUEST_ITEM
	CANCEL_REINCARNATION
	REQUEST_REINCARNATION
	USER_BACK_STEP
	OBJECT_BACK_STEP
	USER_DASH_ATTACK
	OBJECT_DASH_ATTACK
	USER_ATTACK_MOVE
	COMBINE_ITEM
	ITEM_UPGRADED
	SET_CONCENTRATION
	SET_OBJECT_CONCENTRATION
	SET_ELEMENTAL
	SET_OBJECT_ELEMENTAL
	REMOVE_DELAYED_EXPLOSION
	OBJECT_DECO
	OBJECT_SNEAKING
	OBJECT_LEVEL_EFFECTS
	SET_BINDING_SHOT
	SEND_OUTPUT_MESSAGE

	NPC_AWAKENING
	NPC_DISASSEMBLE
	NPC_DOWNGRADE
	NPC_RESET
	AWAKENING_NEED_MATERIALS
	AWAKENING_LOCKEDiTEM
	AWAKENING

	RECEIVE_MAIL
	MAIL_LOCKED_ITEM
	MAIL_SEND_REQUEST
	MAIL_SENT
	PARCEL_COLLECTED
	MAIL_COST
	RESIZE_INVENTORY
	RESIZE_STORAGE
	NEW_INTELLIGENT_CREATURE
	UPDATE_INTELLIGENT_CREATURE_LIST
	INTELLIGENT_CREATURE_ENABLE_RENAME
	INTELLIGENT_CREATURE_PICKUP
	NPC_PEARL_GOODS

	TRANSFORM_UPDATE
	FRIEND_UPDATE
	LOVER_UPDATE
	MENTOR_UPDATE
	GUILD_BUFF_LIST
	NPC_REQUEST_INPUT
	GAME_SHOP_INFO
	GAME_SHOP_STOCK
	RANKINGS
	OPENDOOR

	GET_RENTED_ITEMS
	ITEM_RENTAL_REQUEST
	ITEM_RENTAL_FEE
	ITEM_RENTAL_PERIOD
	DEPOSIT_RENTAL_ITEM
	RETRIEVE_RENTAL_ITEM
	UPDATE_RENTAL_ITEM
	CANCEL_ITEM_RENTAL
	ITEM_RENTAL_LOCK
	ITEM_RENTAL_PARTNER_LOCK
	CAN_CONFIRM_ITEM_RENTAL
	CONFIRM_ITEM_RENTAL
	NEW_RECIPE_INFO
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
	/** 0: Disabled 1: Bad AccountID 2: Bad Password 3: Bad Email 4: Bad Name 5: Bad Question 6: Bad Answer 7: Account Exists. 8: Success*/
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
	/** 0: Disabled 1: Bad AccountID 2: Bad Password 3: Account Not Exist 4: Wrong Password*/
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

type Account struct {
	orm.AccountInfo
	Characters []SelectInfo
}

type SelectInfo struct {
	Index      uint32
	Name       string
	Level      uint16
	Class      cm.MirClass
	Gender     cm.MirGender
	LastAccess uint64
}

func (self *SelectInfo) ToBytes() []byte {
	indexBytes := cm.Uint32ToBytes(self.Index)
	nameBytes := cm.StringToBytes(self.Name)
	levelBytes := cm.Uint16ToBytes(self.Level)
	class := self.Class
	classBytes := []byte{byte(class)}
	gender := self.Gender
	genderBytes := []byte{byte(gender)}
	lastAccessBytes := cm.Uint64ToBytes(uint64(0))
	result := make([]byte, 0)
	for _, r := range [][]byte{indexBytes, nameBytes, levelBytes, classBytes, genderBytes, lastAccessBytes} {
		result = append(result, r...)
	}
	return result
}

type LoginSuccess struct {
	// c#
	// count(int32 4byte) [ index(int32 4byte) name(string) level(int16 2byte) class(1byte) gender(1byte) lastAccess(int64 8byte) ]
	Characters []SelectInfo
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
	/** 0: Disabled. 1: Bad Character Name 2: Bad Gender 3: Bad Class 4: Max Characters 5: Character Exists.**/
	Result byte
}

func (self *NewCharacter) ToBytes() []byte {
	bytes := cm.Uint16ToBytes(NEW_CHARACTER)
	return append(bytes, self.Result)
}

type NewCharacterSuccess struct {
	CharInfo SelectInfo
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
	Index        uint16
	Filename     string
	Title        string
	MiniMap      uint16
	BigMap       uint16
	Music        uint16
	Lightning    bool
	Fire         bool
	MapDarkLight byte
}

func (self *MapInformation) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(MAP_INFORMATION)
	//tmp := []byte{
	//	1, 48,                                                                // index
	//	14, 66, 105, 99, 104, 111, 110, 80, 114, 111, 118, 105, 110, 99, 101, // filename and title
	//	101, 0,                                                               // mini map
	//	135, 0,                                                               // big map
	//	0, 0,                                                                 // music
	//	0,                                                                    // lighting
	//	0,                                                                    // fire
	//	0,                                                                    // map dark light
	//}
	//result := append(pkgBytes, tmp...)
	indexBytes := cm.Uint16ToBytes(self.Index)
	titleBytes := cm.StringToBytes(self.Title)
	miniMapBytes := cm.Uint16ToBytes(self.MiniMap)
	bigMapBytes := cm.Uint16ToBytes(self.BigMap)
	musicBytes := cm.Uint16ToBytes(self.Music)
	lightningBytes := cm.BoolToBytes(self.Lightning)
	fireBytes := cm.BoolToBytes(self.Fire)
	mapDarkLightBytes := []byte{self.MapDarkLight}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, indexBytes, titleBytes, miniMapBytes, bigMapBytes, musicBytes, lightningBytes, fireBytes, mapDarkLightBytes} {
		result = append(result, r...)
	}
	return result
}

type RandomItemStat struct {
	MAX_DURA_CHANCE             byte
	MAX_DURA_STAT_CHANCE        byte
	MAX_DURA_MAX_STAT           byte
	MAX_AC_CHANCE               byte
	MAX_AC_STAT_CHANCE          byte
	MAX_AC_MAX_STAT             byte
	MAX_MAC_CHANCE              byte
	MAX_MAC_STAT_CHANCE         byte
	MAX_MAC_MAX_STAT            byte
	MAX_DC_CHANCE               byte
	MAX_DC_STAT_CHANCE          byte
	MAX_DC_MAX_STAT             byte
	MAX_MC_CHANCE               byte
	MAX_MC_STAT_CHANCE          byte
	MAX_MC_MAX_STAT             byte
	MAX_SC_CHANCE               byte
	MAX_SC_STAT_CHANCE          byte
	MAX_SC_MAX_STAT             byte
	ACCURACY_CHANCE             byte
	ACCURACY_STAT_CHANCE        byte
	ACCURACY_MAX_STAT           byte
	AGILITY_CHANCE              byte
	AGILITY_STAT_CHANCE         byte
	AGILITY_MAX_STAT            byte
	HP_CHANCE                   byte
	HP_STAT_CHANCE              byte
	HP_MAX_STAT                 byte
	MP_CHANCE                   byte
	MP_STAT_CHANCE              byte
	MP_MAX_STAT                 byte
	STRONG_CHANCE               byte
	STRONG_STAT_CHANCE          byte
	STRONG_MAX_STAT             byte
	MAGIC_RESIST_CHANCE         byte
	MAGIC_RESIST_STAT_CHANCE    byte
	MAGIC_RESIST_MAX_STAT       byte
	POISON_RESIST_CHANCE        byte
	POISON_RESIST_STAT_CHANCE   byte
	POISON_RESIST_MAX_STAT      byte
	HP_RECOV_CHANCE             byte
	HP_RECOV_STAT_CHANCE        byte
	HP_RECOV_MAX_STAT           byte
	MP_RECOV_CHANCE             byte
	MP_RECOV_STAT_CHANCE        byte
	MP_RECOV_MAX_STAT           byte
	POISON_RECOV_CHANCE         byte
	POISON_RECOV_STAT_CHANCE    byte
	POISON_RECOV_MAX_STAT       byte
	CRITICAL_RATE_CHANCE        byte
	CRITICAL_RATE_STAT_CHANCE   byte
	CRITICAL_RATE_MAX_STAT      byte
	CRITICAL_DAMAGE_CHANCE      byte
	CRITICAL_DAMAGE_STAT_CHANCE byte
	CRITICAL_DAMAGE_MAX_STAT    byte
	FREEZE_CHANCE               byte
	FREEZE_STAT_CHANCE          byte
	FREEZE_MAX_STAT             byte
	POISON_ATTACK_CHANCE        byte
	POISON_ATTACK_STAT_CHANCE   byte
	POISON_ATTACK_MAX_STAT      byte
	ATTACK_SPEED_CHANCE         byte
	ATTACK_SPEED_STAT_CHANCE    byte
	ATTACK_SPEED_MAX_STAT       byte
	LUCK_CHANCE                 byte
	LUCK_STAT_CHANCE            byte
	LUCK_MAX_STAT               byte
	CURSE_CHANCE                byte
}

func (self *RandomItemStat) ToBytes() []byte {
	return []byte{
		self.MAX_DURA_CHANCE,
		self.MAX_DURA_STAT_CHANCE,
		self.MAX_DURA_MAX_STAT,
		self.MAX_AC_CHANCE,
		self.MAX_AC_STAT_CHANCE,
		self.MAX_AC_MAX_STAT,
		self.MAX_MAC_CHANCE,
		self.MAX_MAC_STAT_CHANCE,
		self.MAX_MAC_MAX_STAT,
		self.MAX_DC_CHANCE,
		self.MAX_DC_STAT_CHANCE,
		self.MAX_DC_MAX_STAT,
		self.MAX_MC_CHANCE,
		self.MAX_MC_STAT_CHANCE,
		self.MAX_MC_MAX_STAT,
		self.MAX_SC_CHANCE,
		self.MAX_SC_STAT_CHANCE,
		self.MAX_SC_MAX_STAT,
		self.ACCURACY_CHANCE,
		self.ACCURACY_STAT_CHANCE,
		self.ACCURACY_MAX_STAT,
		self.AGILITY_CHANCE,
		self.AGILITY_STAT_CHANCE,
		self.AGILITY_MAX_STAT,
		self.HP_CHANCE,
		self.HP_STAT_CHANCE,
		self.HP_MAX_STAT,
		self.MP_CHANCE,
		self.MP_STAT_CHANCE,
		self.MP_MAX_STAT,
		self.STRONG_CHANCE,
		self.STRONG_STAT_CHANCE,
		self.STRONG_MAX_STAT,
		self.MAGIC_RESIST_CHANCE,
		self.MAGIC_RESIST_STAT_CHANCE,
		self.MAGIC_RESIST_MAX_STAT,
		self.POISON_RESIST_CHANCE,
		self.POISON_RESIST_STAT_CHANCE,
		self.POISON_RESIST_MAX_STAT,
		self.HP_RECOV_CHANCE,
		self.HP_RECOV_STAT_CHANCE,
		self.HP_RECOV_MAX_STAT,
		self.MP_RECOV_CHANCE,
		self.MP_RECOV_STAT_CHANCE,
		self.MP_RECOV_MAX_STAT,
		self.POISON_RECOV_CHANCE,
		self.POISON_RECOV_STAT_CHANCE,
		self.POISON_RECOV_MAX_STAT,
		self.CRITICAL_RATE_CHANCE,
		self.CRITICAL_RATE_STAT_CHANCE,
		self.CRITICAL_RATE_MAX_STAT,
		self.CRITICAL_DAMAGE_CHANCE,
		self.CRITICAL_DAMAGE_STAT_CHANCE,
		self.CRITICAL_DAMAGE_MAX_STAT,
		self.FREEZE_CHANCE,
		self.FREEZE_STAT_CHANCE,
		self.FREEZE_MAX_STAT,
		self.POISON_ATTACK_CHANCE,
		self.POISON_ATTACK_STAT_CHANCE,
		self.POISON_ATTACK_MAX_STAT,
		self.ATTACK_SPEED_CHANCE,
		self.ATTACK_SPEED_STAT_CHANCE,
		self.ATTACK_SPEED_MAX_STAT,
		self.LUCK_CHANCE,
		self.LUCK_STAT_CHANCE,
		self.LUCK_MAX_STAT,
		self.CURSE_CHANCE,
	}
}

type ItemInfo struct {
	Index            uint32
	Name             string
	Type             cm.ItemType
	Grade            cm.ItemGrade
	RequiredType     cm.RequiredType   // default Level
	RequiredClass    cm.RequiredClass  // default None
	RequiredGender   cm.RequiredGender // default None
	Set              cm.ItemSet
	Shape            uint16
	Weight           byte
	Light            byte
	RequiredAmount   byte
	Image            uint16
	Durability       uint16
	Price            uint16
	StackSize        uint16 //default 1;
	MinAC            byte
	MaxAC            byte
	MinMAC           byte
	MaxMAC           byte
	MinDC            byte
	MaxDC            byte
	MinMC            byte
	MaxMC            byte
	MinSC            byte
	MaxSC            byte
	Accuracy         byte
	Agility          byte
	HP               uint16
	MP               uint16
	AttackSpeed      int8 // 需要是负数
	Luck             int8
	BagWeight        byte
	HandWeight       byte
	WearWeight       byte
	StartItem        bool
	Effect           byte
	Strong           byte
	MagicResist      byte
	PoisonResist     byte
	HealthRecovery   byte
	SpellRecovery    byte
	PoisonRecovery   byte
	HPrate           byte
	MPrate           byte
	CriticalRate     byte
	CriticalDamage   byte
	NeedIdentify     bool
	ShowGroupPickup  bool
	GlobalDropNotify bool
	ClassBased       bool
	LevelBased       bool
	CanMine          bool
	CanFastRun       bool
	CanAwakening     bool
	MaxAcRate        byte
	MaxMacRate       byte
	Holy             byte
	Freezing         byte
	PoisonAttack     byte
	HpDrainRate      byte
	Bind             uint16 // BindMode 这个枚举太大了，直接用uint16 // default none
	Reflect          byte
	Unique           uint16 // SpecialItemMode ?? // default None;
	RandomStatsId    byte
	RandomStats      RandomItemStat
	ToolTip          string //default ""
}

// TODO
func (self *ItemInfo) ToBytes() []byte {
	return nil
}

type UserItem struct {
	UniqueID       uint64
	ItemIndex      uint32
	Info           ItemInfo
	CurrentDura    uint16
	MaxDura        uint16
	Count          uint32
	GemCount       uint32
	AC             byte
	MAC            byte
	DC             byte
	MC             byte
	SC             byte
	Accuracy       byte
	Agility        byte
	HP             byte
	MP             byte
	Strong         byte
	MagicResist    byte
	PoisonResist   byte
	HealthRecovery byte
	ManaRecovery   byte
	PoisonRecovery byte
	CriticalRate   byte
	CriticalDamage byte
	Freezing       byte
	PoisonAttack   byte
	AttackSpeed    byte
	Luck           byte
	RefinedValue   cm.RefinedValue
	RefineAdded    byte
	DuraChanged    bool
	SoulBoundId    uint32
	Identified     bool
	Cursed         bool
	WeddingRing    uint32
	//public UserItem[] Slots = new UserItem[5];
	//public DateTime BuybackExpiryDate;
	//public ExpireInfo ExpireInfo;
	//public RentalInformation RentalInformation;
	//public Awake Awake = new Awake();
}

// TODO
func (self *UserItem) ToBytes() []byte {
	return nil
}

type ClientMagic struct {
	Spell       cm.Spell
	BaseCost    byte
	LevelCost   byte
	Icon        byte
	Level1      byte
	Level2      byte
	Level3      byte
	Need1       uint16
	Need2       uint16
	Need3       uint16
	Level       byte
	Key         byte
	Range       byte
	Experience  uint16
	IsTempSpell bool
	CastTime    uint64
	Delay       uint64
}

// TODO
func (self *ClientMagic) ToBytes() []byte {
	return nil
}

type IntelligentCreatureRules struct {
	MinimalFullness       uint32 // default 1
	MousePickupEnabled    bool   // default false
	MousePickupRange      uint32 // default 0
	AutoPickupEnabled     bool   // default false
	AutoPickupRange       uint32 // default 0
	SemiAutoPickupEnabled bool   // default false
	SemiAutoPickupRange   uint32 // default 0
	CanProduceBlackStone  bool   // default false
	Info                  string // default ""
	Info1                 string // default ""
	Info2                 string // default ""
}

// TODO
func (self *IntelligentCreatureRules) ToBytes() []byte {
	return nil
}

type IntelligentCreatureItemFilter struct {
	PetPickupAll         bool         // default true
	PetPickupGold        bool         // default false
	PetPickupWeapons     bool         // default false
	PetPickupArmours     bool         // default false
	PetPickupHelmets     bool         // default false
	PetPickupBoots       bool         // default false
	PetPickupBelts       bool         // default false
	PetPickupAccessories bool         // default false
	PetPickupOthers      bool         // default false
	PickupGrade          cm.ItemGrade // default ItemGrade.None;
}

// TODO
func (self *IntelligentCreatureItemFilter) ToBytes() []byte {
	return nil
}

type ClientIntelligentCreature struct {
	PetType          cm.IntelligentCreatureType
	Icon             uint32
	CustomName       string
	Fullness         uint32
	SlotIndex        uint32
	ExpireTime       uint64 // long
	BlackstoneTime   uint64
	MaintainFoodTime uint64
	petMode          cm.IntelligentCreaturePickupMode // default SemiAutomatic
	CreatureRules    IntelligentCreatureRules
	Filter           IntelligentCreatureItemFilter
}

// TODO
func (self *ClientIntelligentCreature) ToBytes() []byte {
	return nil
}

type UserInformation struct {
	ObjectID                  uint32
	RealId                    uint32
	Name                      string
	GuildName                 string
	GuildRank                 string
	NameColour                uint32
	Class                     cm.MirClass
	Gender                    cm.MirGender
	Level                     uint16
	Location                  cm.Point
	Direction                 cm.MirDirection
	Hair                      byte
	HP                        uint16
	MP                        uint16
	Experience                uint64
	MaxExperience             uint64
	LevelEffect               byte        // LevelEffects
	Inventory                 interface{} // []UserItem				// TODO UNKNOW
	Equipment                 interface{} // []UserItem
	QuestInventory            interface{} // []UserItem
	Gold                      uint32
	Credit                    uint32
	HasExpandedStorage        bool
	ExpandedStorageExpiryTime uint64      // DateTime
	Magics                    interface{} // []ClientMagic
	IntelligentCreatures      interface{} // []ClientIntelligentCreature // TODO
	IntelligentCreatureType   cm.IntelligentCreatureType
	CreatureSummoned          bool
}

func (self *UserInformation) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(USER_INFORMATION)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
	realIdBytes := cm.Uint32ToBytes(self.RealId)
	nameBytes := cm.StringToBytes(self.Name)
	guildNameBytes := cm.StringToBytes(self.GuildName)
	guildRankBytes := cm.StringToBytes(self.GuildRank)
	nameColorBytes := []byte{255, 255, 255, 255}
	classBytes := []byte{byte(self.Class)}
	genderBytes := []byte{byte(self.Gender)}
	levelBytes := cm.Uint16ToBytes(self.Level)
	locationBytes := self.Location.ToBytes()
	directionBytes := []byte{byte(self.Direction)}
	hairBytes := []byte{byte(self.Hair)}
	hpBytes := cm.Uint16ToBytes(self.HP)
	mpBytes := cm.Uint16ToBytes(self.MP)
	experienceBytes := cm.Uint64ToBytes(self.Experience)
	maxExperienceBytes := cm.Uint64ToBytes(self.MaxExperience)
	levelEffectBytes := []byte{byte(self.LevelEffect)}
	unknowBytes := []byte{1, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	intelligentCreatureTypeBytes := []byte{byte(self.IntelligentCreatureType)}
	creatureSummonedBytes := cm.BoolToBytes(self.CreatureSummoned)

	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, realIdBytes, nameBytes, guildNameBytes, guildRankBytes, nameColorBytes, classBytes, genderBytes, levelBytes, locationBytes, directionBytes, hairBytes, hpBytes, mpBytes, experienceBytes, maxExperienceBytes, levelEffectBytes, unknowBytes, intelligentCreatureTypeBytes, creatureSummonedBytes} {
		result = append(result, r...)
	}
	return result
	//result := append(pkgBytes, []byte{
	//	165, 7, 0, 0,              // object id
	//	1, 0, 0, 0,                // real id
	//	3, 54, 54, 51,             // name
	//	0,                         // guild name
	//	0,                         // guild rank
	//	255, 255, 255, 255,        // name color
	//	0,                         // class
	//	0,                         // gender
	//	1, 0,                      // level
	//	31, 1, 0, 0, 100, 2, 0, 0, // location ??
	//	4,                         // direction
	//	4,                         // hair
	//	18, 0,                     // hp
	//	14, 0,                     // mp
	//	0, 0, 0, 0, 0, 0, 0, 0,    // experience
	//	100, 0, 0, 0, 0, 0, 0, 0,  // max experience
	//	0,                         // level effect
	//	1, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	99, // intelligent creature type
	//	0,  // creature summoned
	//}...)
	//return result
}

type UserLocation struct {
	Location  cm.Point
	Direction cm.MirDirection
}

func (self *UserLocation) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(USER_LOCATION)
	locationBytes := self.Location.ToBytes()
	directionBytes := []byte{byte(self.Direction)}
	tmp := append(pkgBytes, locationBytes...)
	result := append(tmp, directionBytes...)
	return result
}

type ObjectPlayer struct {
}

func (self *ObjectPlayer) ToBytes() []byte {
	return nil
}

type ObjectRemove struct {
}

type ObjectTurn struct {
	ObjectID  uint32
	Location  cm.Point
	Direction cm.MirDirection
}

func (self *ObjectTurn) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(OBJECT_TURN)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
	locationBytes := self.Location.ToBytes()
	directionBytes := []byte{byte(self.Direction)}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, locationBytes, directionBytes} {
		result = append(result, r...)
	}
	return result
}

type ObjectWalk struct {
	ObjectID  uint32
	Location  cm.Point
	Direction cm.MirDirection
}

func (self *ObjectWalk) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(OBJECT_WALK)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
	locationBytes := self.Location.ToBytes()
	directionBytes := []byte{byte(self.Direction)}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, locationBytes, directionBytes} {
		result = append(result, r...)
	}
	return result
}

type ObjectRun struct {
	ObjectID  uint32
	Location  cm.Point
	Direction cm.MirDirection
}

func (self *ObjectRun) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(OBJECT_RUN)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
	locationBytes := self.Location.ToBytes()
	directionBytes := []byte{byte(self.Direction)}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, locationBytes, directionBytes} {
		result = append(result, r...)
	}
	return result
}

type Chat struct {
	Message string
	Type    cm.ChatType
}

func (self *Chat) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(CHAT)
	messageBytes := cm.StringToBytes(self.Message)
	typeBytes := []byte{byte(self.Type)}
	tmp := append(pkgBytes, messageBytes...)
	result := append(tmp, typeBytes...)
	return result
}

type ObjectChat struct {
	ObjectID uint32
	Text     string
	Type     cm.ChatType
}

func (self *ObjectChat) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(OBJECT_CHAT)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
	messageBytes := cm.StringToBytes(self.Text)
	typeBytes := []byte{byte(self.Type)}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, messageBytes, typeBytes} {
		result = append(result, r...)
	}
	return result
}

type NewItemInfo struct {
}

type MoveItem struct {
}

type EquipItem struct {
}

type MergeItem struct {
}

type RemoveItem struct {
}

type RemoveSlotItem struct {
}

type TakeBackItem struct {
}

type StoreItem struct {
}

type SplitItem struct {
}

type SplitItem1 struct {
}

type DepositRefineItem struct {
}

type RetrieveRefineItem struct {
}

type RefineCancel struct {
}

type RefineItem struct {
}

type DepositTradeItem struct {
}

type RetrieveTradeItem struct {
}

type UseItem struct {
}

type DropItem struct {
}

type PlayerUpdate struct {
}

type PlayerInspect struct {
}

type LogOutSuccess struct {
}

type LogOutFailed struct {
}

type TimeOfDay struct {
}

type ChangeAMode struct {
}

type ChangePMode struct {
}

//type NewQuestInfo struct {
//
//}
//
//func (self *NewQuestInfo) ToBytes() []byte {
//	return
//}
