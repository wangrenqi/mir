package server

import (
	"mir/com"
	"log"
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
	bytes := com.Uint16ToBytes(CONNECTED)
	return bytes
}

type ClientVersion struct {
	// 0 wrong version
	// 1 correct version
	Result byte
}

func (self *ClientVersion) ToBytes() []byte {
	bytes := com.Uint16ToBytes(CLIENT_VERSION)
	bytes = append(bytes, self.Result)
	return bytes
}

type Disconnect struct{}

type NewAccount struct {
	/** 0: Disabled 1: Bad AccountID 2: Bad Password 3: Bad Email 4: Bad Name 5: Bad Question 6: Bad Answer 7: Account Exists. 8: Success*/
	Result byte
}

func (self *NewAccount) ToBytes() []byte {
	bytes := com.Uint16ToBytes(NEW_ACCOUNT)
	return append(bytes, self.Result)
}

type ChangePassword struct {
	Result byte
}

func (self *ChangePassword) ToBytes() []byte {
	bytes := com.Uint16ToBytes(NEW_ACCOUNT)
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
	bytes := com.Uint16ToBytes(LOGIN)
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
	com.AccountInfo
	Characters []SelectInfo
}

type SelectInfo struct {
	Index      uint32
	Name       string
	Level      uint16
	Class      com.MirClass
	Gender     com.MirGender
	LastAccess uint64
}

func (self *SelectInfo) ToBytes() []byte {
	indexBytes := com.Uint32ToBytes(self.Index)
	nameBytes := com.StringToBytes(self.Name)
	levelBytes := com.Uint16ToBytes(self.Level)
	class := self.Class
	classBytes := []byte{byte(class)}
	gender := self.Gender
	genderBytes := []byte{byte(gender)}
	lastAccessBytes := com.Uint64ToBytes(uint64(0))
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
	bytes := com.Uint16ToBytes(LOGIN_SUCCESS)
	characters := self.Characters
	count := len(characters)
	if count == 0 {
		bytes = append(bytes, []byte{0, 0, 0, 0}...)
	} else {
		countBytes := com.Uint32ToBytes(uint32(count))
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
	bytes := com.Uint16ToBytes(NEW_CHARACTER)
	return append(bytes, self.Result)
}

type NewCharacterSuccess struct {
	CharInfo SelectInfo
}

func (self *NewCharacterSuccess) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(NEW_CHARACTER_SUCCESS)
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
	pkgBytes := com.Uint16ToBytes(START_GAME)
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
	pkgBytes := com.Uint16ToBytes(MAP_INFORMATION)
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
	indexBytes := com.Uint16ToBytes(self.Index)
	titleBytes := com.StringToBytes(self.Title)
	miniMapBytes := com.Uint16ToBytes(self.MiniMap)
	bigMapBytes := com.Uint16ToBytes(self.BigMap)
	musicBytes := com.Uint16ToBytes(self.Music)
	lightningBytes := com.BoolToBytes(self.Lightning)
	fireBytes := com.BoolToBytes(self.Fire)
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
		self.MAX_DURA_CHANCE, self.MAX_DURA_STAT_CHANCE, self.MAX_DURA_MAX_STAT, self.MAX_AC_CHANCE, self.MAX_AC_STAT_CHANCE,
		self.MAX_AC_MAX_STAT, self.MAX_MAC_CHANCE, self.MAX_MAC_STAT_CHANCE, self.MAX_MAC_MAX_STAT, self.MAX_DC_CHANCE,
		self.MAX_DC_STAT_CHANCE, self.MAX_DC_MAX_STAT, self.MAX_MC_CHANCE, self.MAX_MC_STAT_CHANCE, self.MAX_MC_MAX_STAT,
		self.MAX_SC_CHANCE, self.MAX_SC_STAT_CHANCE, self.MAX_SC_MAX_STAT, self.ACCURACY_CHANCE, self.ACCURACY_STAT_CHANCE,
		self.ACCURACY_MAX_STAT, self.AGILITY_CHANCE, self.AGILITY_STAT_CHANCE, self.AGILITY_MAX_STAT, self.HP_CHANCE,
		self.HP_STAT_CHANCE, self.HP_MAX_STAT, self.MP_CHANCE, self.MP_STAT_CHANCE, self.MP_MAX_STAT, self.STRONG_CHANCE,
		self.STRONG_STAT_CHANCE, self.STRONG_MAX_STAT, self.MAGIC_RESIST_CHANCE, self.MAGIC_RESIST_STAT_CHANCE, self.MAGIC_RESIST_MAX_STAT,
		self.POISON_RESIST_CHANCE, self.POISON_RESIST_STAT_CHANCE, self.POISON_RESIST_MAX_STAT, self.HP_RECOV_CHANCE,
		self.HP_RECOV_STAT_CHANCE, self.HP_RECOV_MAX_STAT, self.MP_RECOV_CHANCE, self.MP_RECOV_STAT_CHANCE, self.MP_RECOV_MAX_STAT,
		self.POISON_RECOV_CHANCE, self.POISON_RECOV_STAT_CHANCE, self.POISON_RECOV_MAX_STAT, self.CRITICAL_RATE_CHANCE,
		self.CRITICAL_RATE_STAT_CHANCE, self.CRITICAL_RATE_MAX_STAT, self.CRITICAL_DAMAGE_CHANCE, self.CRITICAL_DAMAGE_STAT_CHANCE,
		self.CRITICAL_DAMAGE_MAX_STAT, self.FREEZE_CHANCE, self.FREEZE_STAT_CHANCE, self.FREEZE_MAX_STAT, self.POISON_ATTACK_CHANCE,
		self.POISON_ATTACK_STAT_CHANCE, self.POISON_ATTACK_MAX_STAT, self.ATTACK_SPEED_CHANCE, self.ATTACK_SPEED_STAT_CHANCE,
		self.ATTACK_SPEED_MAX_STAT, self.LUCK_CHANCE, self.LUCK_STAT_CHANCE, self.LUCK_MAX_STAT, self.CURSE_CHANCE,
	}
}

type ClientMagic struct {
	Spell       com.Spell
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

func (self *ClientMagic) ToBytes() []byte {
	spellBytes := []byte{byte(self.Spell)}
	baseCostBytes := []byte{byte(self.BaseCost)}
	levelCostBytes := []byte{byte(self.LevelCost)}
	iconBytes := []byte{byte(self.Icon)}
	level1Bytes := []byte{byte(self.Level1)}
	level2Bytes := []byte{byte(self.Level2)}
	level3Bytes := []byte{byte(self.Level3)}
	need1Bytes := com.Uint16ToBytes(self.Need1)
	need2Bytes := com.Uint16ToBytes(self.Need2)
	need3Bytes := com.Uint16ToBytes(self.Need3)
	levelBytes := []byte{byte(self.Level)}
	keyBytes := []byte{byte(self.Key)}
	rangeBytes := []byte{byte(self.Range)}
	experienceBytes := com.Uint16ToBytes(self.Experience)
	isTempSpellBytes := com.BoolToBytes(self.IsTempSpell)
	castTimeBytes := com.Uint64ToBytes(self.CastTime)
	delayBytes := com.Uint64ToBytes(self.Delay)
	result := make([]byte, 0)
	for _, r := range [][]byte{spellBytes, baseCostBytes, levelCostBytes, iconBytes, level1Bytes, level2Bytes, level3Bytes,
		need1Bytes, need2Bytes, need3Bytes, levelBytes, keyBytes, rangeBytes, experienceBytes, isTempSpellBytes,
		castTimeBytes, delayBytes,
	} {
		result = append(result, r...)
	}
	return result
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

func (self *IntelligentCreatureRules) ToBytes() []byte {
	minimalFullnessBytes := com.Uint32ToBytes(self.MinimalFullness)
	mousePickupEnabledBytes := com.BoolToBytes(self.MousePickupEnabled)
	mousePickupRangeBytes := com.Uint32ToBytes(self.MousePickupRange)
	autoPickupEnabledBytes := com.BoolToBytes(self.AutoPickupEnabled)
	autoPickupRangeBytes := com.Uint32ToBytes(self.AutoPickupRange)
	semiAutoPickupEnabledBytes := com.BoolToBytes(self.SemiAutoPickupEnabled)
	semiAutoPickupRangeBytes := com.Uint32ToBytes(self.SemiAutoPickupRange)
	canProduceBlackStoneBytes := com.BoolToBytes(self.CanProduceBlackStone)
	infoBytes := com.StringToBytes(self.Info)
	info1Bytes := com.StringToBytes(self.Info1)
	info2Bytes := com.StringToBytes(self.Info2)
	result := make([]byte, 0)
	for _, r := range [][]byte{minimalFullnessBytes, mousePickupEnabledBytes, mousePickupRangeBytes, autoPickupEnabledBytes,
		autoPickupRangeBytes, semiAutoPickupEnabledBytes, semiAutoPickupRangeBytes, canProduceBlackStoneBytes, infoBytes,
		info1Bytes, info2Bytes,
	} {
		result = append(result, r...)
	}
	return result
}

type IntelligentCreatureItemFilter struct {
	PetPickupAll         bool          // default true
	PetPickupGold        bool          // default false
	PetPickupWeapons     bool          // default false
	PetPickupArmours     bool          // default false
	PetPickupHelmets     bool          // default false
	PetPickupBoots       bool          // default false
	PetPickupBelts       bool          // default false
	PetPickupAccessories bool          // default false
	PetPickupOthers      bool          // default false
	PickupGrade          com.ItemGrade // default ItemGrade.None;
}

func (self *IntelligentCreatureItemFilter) ToBytes() []byte {
	petPickupAllBytes := com.BoolToBytes(self.PetPickupAll)
	petPickupGoldBytes := com.BoolToBytes(self.PetPickupGold)
	petPickupWeaponsBytes := com.BoolToBytes(self.PetPickupWeapons)
	petPickupArmoursBytes := com.BoolToBytes(self.PetPickupArmours)
	petPickupHelmetsBytes := com.BoolToBytes(self.PetPickupHelmets)
	petPickupBootsBytes := com.BoolToBytes(self.PetPickupBoots)
	petPickupBeltsBytes := com.BoolToBytes(self.PetPickupBelts)
	petPickupAccessoriesBytes := com.BoolToBytes(self.PetPickupAccessories)
	petPickupOthersBytes := com.BoolToBytes(self.PetPickupOthers)
	pickupGradeBytes := []byte{byte(self.PickupGrade)}
	result := make([]byte, 0)
	for _, r := range [][]byte{petPickupAllBytes, petPickupGoldBytes, petPickupWeaponsBytes, petPickupArmoursBytes,
		petPickupHelmetsBytes, petPickupBootsBytes, petPickupBeltsBytes, petPickupAccessoriesBytes, petPickupOthersBytes, pickupGradeBytes,
	} {
		result = append(result, r...)
	}
	return result
}

type ClientIntelligentCreature struct {
	PetType          com.IntelligentCreatureType
	Icon             uint32
	CustomName       string
	Fullness         uint32
	SlotIndex        uint32
	ExpireTime       uint64                            // long
	BlackstoneTime   uint64                            // long
	MaintainFoodTime uint64                            // long
	PetMode          com.IntelligentCreaturePickupMode // default SemiAutomatic
	CreatureRules    IntelligentCreatureRules
	Filter           IntelligentCreatureItemFilter
}

func (self *ClientIntelligentCreature) ToBytes() []byte {
	petTypeBytes := []byte{byte(self.PetType)}
	iconBytes := com.Uint32ToBytes(self.Icon)
	customNameBytes := com.StringToBytes(self.CustomName)
	fullnessBytes := com.Uint32ToBytes(self.Fullness)
	slotIndexBytes := com.Uint32ToBytes(self.SlotIndex)
	expireTimeBytes := com.Uint64ToBytes(self.ExpireTime)
	blackstoneTimeBytes := com.Uint64ToBytes(self.BlackstoneTime)
	maintainFoodTimeBytes := com.Uint64ToBytes(self.MaintainFoodTime)
	petModeBytes := []byte{byte(self.PetMode)}
	creatureRulesBytes := self.CreatureRules.ToBytes()
	filterBytes := self.Filter.ToBytes()
	result := make([]byte, 0)
	for _, r := range [][]byte{petTypeBytes, iconBytes, customNameBytes, fullnessBytes, slotIndexBytes,
		expireTimeBytes, blackstoneTimeBytes, maintainFoodTimeBytes, petModeBytes, creatureRulesBytes, filterBytes,
	} {
		result = append(result, r...)
	}
	return result
}

type UserInformation struct {
	ObjectID                  uint32
	RealId                    uint32
	Name                      string
	GuildName                 string
	GuildRank                 string
	NameColour                uint32
	Class                     com.MirClass
	Gender                    com.MirGender
	Level                     uint16
	Location                  com.Point
	Direction                 com.MirDirection
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
	IntelligentCreatureType   com.IntelligentCreatureType
	CreatureSummoned          bool
}

func (self *UserInformation) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(USER_INFORMATION)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
	realIdBytes := com.Uint32ToBytes(self.RealId)
	nameBytes := com.StringToBytes(self.Name)
	guildNameBytes := com.StringToBytes(self.GuildName)
	guildRankBytes := com.StringToBytes(self.GuildRank)
	nameColorBytes := []byte{255, 255, 255, 255}
	classBytes := []byte{byte(self.Class)}
	genderBytes := []byte{byte(self.Gender)}
	levelBytes := com.Uint16ToBytes(self.Level)
	locationBytes := self.Location.ToBytes()
	directionBytes := []byte{byte(self.Direction)}
	hairBytes := []byte{byte(self.Hair)}
	hpBytes := com.Uint16ToBytes(self.HP)
	mpBytes := com.Uint16ToBytes(self.MP)
	experienceBytes := com.Uint64ToBytes(self.Experience)
	maxExperienceBytes := com.Uint64ToBytes(self.MaxExperience)
	levelEffectBytes := []byte{byte(self.LevelEffect)}
	//unknowBytes := []byte{1, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	unknowBytes1 := []byte{1, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14, 0, 0, 0, 1, 32, 0, 0, 0, 0, 0, 0, 0, 248, 0, 0, 0, 158, 95, 232, 128, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	goldBytes := com.Uint32ToBytes(self.Gold)
	creditBytes := com.Uint32ToBytes(self.Credit)
	unknowBytes2 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 99, 0}
	intelligentCreatureTypeBytes := []byte{byte(self.IntelligentCreatureType)}
	creatureSummonedBytes := com.BoolToBytes(self.CreatureSummoned)

	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, realIdBytes, nameBytes, guildNameBytes, guildRankBytes, nameColorBytes, classBytes, genderBytes, levelBytes, locationBytes, directionBytes, hairBytes, hpBytes, mpBytes, experienceBytes, maxExperienceBytes, levelEffectBytes, unknowBytes1, goldBytes, creditBytes, unknowBytes2, intelligentCreatureTypeBytes, creatureSummonedBytes} {
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
	//	1, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	//       0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	99, // intelligent creature type
	//	0,  // creature summoned
	//}...)
	//return result

	// {{ 10, 1,		length
	// 18, 0, 			pkgIndex
	// 165, 7, 0, 0,	object id
	// 7, 0, 0, 0,		real id
	// 3, 49, 49, 49, 	name
	// 0,
	// 0,
	// 255, 255, 255, 255,
	// 0,
	// 0,
	// 255, 255,
	// 29, 1, 0, 0, 97, 2, 0, 0,
	// 2,
	// 1,
	// 255, 255,
	// 255, 255,
	// 3, 0, 0, 0, 0, 0, 0, 0,
	// 0, 0, 0, 0, 0, 0, 0, 0,
	// 0,
	// 1, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14, 0, 0, 0, 1, 32, 0, 0, 0, 0, 0, 0, 0, 248, 0, 0, 0, 158, 95, 232, 128, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 123, 214, 16, 0, 		gold
	// 215, 136, 1, 0, 			credit
	// 0, 						HasExpandedStorage
	// 0, 0, 0, 0, 0, 0, 0, 0, 	ExpandedStorageExpiryTime
	// 0, 0, 0, 0, 0, 0, 0, 0,  ???  Magics([]ClientMagic)  IntelligentCreatures([]ClientIntelligentCreature )
	// 99,						IntelligentCreatureType
	// 0, 						CreatureSummoned
	// }}
}

type UserLocation struct {
	Location  com.Point
	Direction com.MirDirection
}

func (self *UserLocation) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(USER_LOCATION)
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
	Location  com.Point
	Direction com.MirDirection
}

func (self *ObjectTurn) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(OBJECT_TURN)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
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
	Location  com.Point
	Direction com.MirDirection
}

func (self *ObjectWalk) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(OBJECT_WALK)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
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
	Location  com.Point
	Direction com.MirDirection
}

func (self *ObjectRun) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(OBJECT_RUN)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
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
	Type    com.ChatType
}

func (self *Chat) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(CHAT)
	messageBytes := com.StringToBytes(self.Message)
	typeBytes := []byte{byte(self.Type)}
	tmp := append(pkgBytes, messageBytes...)
	result := append(tmp, typeBytes...)
	return result
}

type ObjectChat struct {
	ObjectID uint32
	Text     string
	Type     com.ChatType
}

func (self *ObjectChat) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(OBJECT_CHAT)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
	messageBytes := com.StringToBytes(self.Text)
	typeBytes := []byte{byte(self.Type)}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, messageBytes, typeBytes} {
		result = append(result, r...)
	}
	return result
}

type NewItemInfo struct {
	Info com.ItemInfo
}

// TODO
func (self *NewItemInfo) ToBytes() []byte {
	// {{ 94, 0,
	// 27, 0,			pkg index
	// 248, 0, 0, 0, 	index
	// 12, 68, 114, 97, 103, 111, 110, 83, 108, 97, 121, 101, 114,		name
	// 1,				item type
	// 2, 				item grade
	// 0,				required type
	// 7,				required class
	// 3,				required gender
	// 0,				item set
	// 29, 0, 			shape
	// 92, 				weight
	// 0,				light
	// 40,				required amount
	// 57, 0, 			image
	// 232, 128,		durability
	// 1, 0, 0, 0,		price
	// 248, 36, 1, 		stack size
	// 0, 0, 0, 0, 0, 5, 40, 0, 0, 0, 0, 0, 	MinAC ~ Agility
	// 0, 0, 			hp
	// 0, 0, 			mp
	// 0, 				attack speed
	// 0,				luck
	// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 	bag weight ~ MPrate
	// 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 	CriticalRate ~ HpDrainRate
	// 0, 1, 1, 0, }}   ???
	return []byte{27, 0, 248, 0, 0, 0, 12, 68, 114, 97, 103, 111, 110, 83, 108, 97, 121, 101, 114, 1, 2, 0, 7, 3, 0, 29, 0, 92, 0, 40, 57, 0, 232, 128, 1, 0, 0, 0, 248, 36, 1, 0, 0, 0, 0, 0, 5, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0}
}

type MoveItem struct {
	Grid    com.MirGridType
	From    uint32
	To      uint32
	Success bool
}

func (self *MoveItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(MOVE_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	fromBytes := com.Uint32ToBytes(self.From)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type EquipItem struct {
	Grid     com.MirGridType
	UniqueID uint64
	To       uint32
	Success  bool
}

func (self *EquipItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(EQUIP_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	uniqueIdBytes := com.Uint64ToBytes(self.UniqueID)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, uniqueIdBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type MergeItem struct {
	GridFrom com.MirGridType
	GridTo   com.MirGridType
	IDFrom   uint64
	IDTo     uint64
	Success  bool
}

func (self *MergeItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(MERGE_ITEM)
	gridFromBytes := []byte{byte(self.GridFrom)}
	gridToBytes := []byte{byte(self.GridTo)}
	idFromBytes := com.Uint64ToBytes(self.IDFrom)
	idToBytes := com.Uint64ToBytes(self.IDTo)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridFromBytes, gridToBytes, idFromBytes, idToBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type RemoveItem struct {
	Grid     com.MirGridType
	UniqueID uint64
	To       uint32
	Success  bool
}

func (self *RemoveItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(REMOVE_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	uniqueIdBytes := com.Uint64ToBytes(self.UniqueID)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, uniqueIdBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type RemoveSlotItem struct {
	Grid     com.MirGridType
	GridTo   com.MirGridType
	UniqueID uint64
	To       uint32
	Success  bool
}

func (self *RemoveSlotItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(REMOVE_SLOT_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	gridToBytes := []byte{byte(self.GridTo)}
	uniqueIdBytes := com.Uint64ToBytes(self.UniqueID)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, gridToBytes, uniqueIdBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type TakeBackItem struct {
	From    uint32
	To      uint32
	Success bool
}

func (self *TakeBackItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(TAKE_BACK_ITEM)
	fromBytes := com.Uint32ToBytes(self.From)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type StoreItem struct {
	From    uint32
	To      uint32
	Success bool
}

func (self *StoreItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(STORE_ITEM)
	fromBytes := com.Uint32ToBytes(self.From)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type SplitItem struct {
	Item com.UserItem
	Grid com.MirGridType
}

// TODO
func (self SplitItem) ToBytes() []byte {
	//pkgBytes := com.Uint16ToBytes(SPLIT_ITEM)
	return nil
}

type SplitItem1 struct {
	Grid     com.MirGridType
	UniqueID uint64
	Count    uint32
	Success  bool
}

func (self *SplitItem1) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(SPLIT_ITEM1)
	gridBytes := []byte{byte(self.Grid)}
	uniqueIdBytes := com.Uint64ToBytes(self.UniqueID)
	countBytes := com.Uint32ToBytes(self.Count)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, uniqueIdBytes, countBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type DepositRefineItem struct {
	From    uint32
	To      uint32
	Success bool
}

func (self *DepositRefineItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(DEPOSIT_REFINE_ITEM)
	fromBytes := com.Uint32ToBytes(self.From)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type RetrieveRefineItem struct {
	From    uint32
	To      uint32
	Success bool
}

func (self *RetrieveRefineItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(RETRIEVE_REFINE_ITEM)
	fromBytes := com.Uint32ToBytes(self.From)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type RefineCancel struct {
	Unlock bool
}

func (self *RefineCancel) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(REFINE_CANCEL)
	unlockBytes := com.BoolToBytes(self.Unlock)
	return append(pkgBytes, unlockBytes...)
}

type RefineItem struct {
	UniqueID uint64
}

func (self *RefineItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(REFINE_ITEM)
	uniqueIdBytes := com.Uint64ToBytes(self.UniqueID)
	return append(pkgBytes, uniqueIdBytes...)
}

type DepositTradeItem struct {
	From    uint32
	To      uint32
	Success bool
}

func (self *DepositTradeItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(DEPOSIT_TRADE_ITEM)
	fromBytes := com.Uint32ToBytes(self.From)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type RetrieveTradeItem struct {
	From    uint32
	To      uint32
	Success bool
}

func (self *RetrieveTradeItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(RETRIEVE_TRADE_ITEM)
	fromBytes := com.Uint32ToBytes(self.From)
	toBytes := com.Uint32ToBytes(self.To)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type UseItem struct {
	UniqueID uint64
	Success  bool
}

func (self *UseItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(USE_ITEM)
	uniqueIdBytes := com.Uint64ToBytes(self.UniqueID)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, uniqueIdBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type DropItem struct {
	UniqueID uint64
	Count    uint32
	Success  bool
}

func (self *DropItem) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(DROP_ITEM)
	uniqueIdBytes := com.Uint64ToBytes(self.UniqueID)
	countBytes := com.Uint32ToBytes(self.Count)
	successBytes := com.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, uniqueIdBytes, countBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type PlayerUpdate struct {
	ObjectID     uint32
	Light        byte
	Weapon       uint16
	WeaponEffect uint16
	Armour       uint16
	WingEffect   byte
}

func (self *PlayerUpdate) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(PLAYER_UPDATE)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
	lightBytes := []byte{self.Light}
	weaponBytes := com.Uint16ToBytes(self.Weapon)
	weaponEffectBytes := com.Uint16ToBytes(self.WeaponEffect)
	armourBytes := com.Uint16ToBytes(self.Armour)
	wingEffectBytes := []byte{self.WingEffect}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, lightBytes, weaponBytes, weaponEffectBytes, armourBytes, wingEffectBytes} {
		result = append(result, r...)
	}
	return result
}

type PlayerInspect struct {
	// TODO
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

type ObjectItem struct{}

func (self *ObjectItem) ToBytes() []byte { return nil }

type ObjectGold struct{}

func (self *ObjectGold) ToBytes() []byte { return nil }

type GainedItem struct{}

func (self *GainedItem) ToBytes() []byte { return nil }

type GainedGold struct{}

func (self *GainedGold) ToBytes() []byte { return nil }

type LoseGold struct{}

func (self *LoseGold) ToBytes() []byte { return nil }

type GainedCredit struct{}

func (self *GainedCredit) ToBytes() []byte { return nil }

type LoseCredit struct{}

func (self *LoseCredit) ToBytes() []byte { return nil }

type ObjectMonster struct {
	ObjectID          uint32
	Name              string
	NameColour        uint32 // Color
	Location          com.Point
	Image             com.Monster
	Direction         com.MirDirection
	Effect            byte
	AI                byte
	Light             byte
	Dead              bool
	Skeleton          bool
	Poison            com.PoisonType
	Hidden            bool
	Extra             bool
	ExtraByte         byte
	ShockTime         uint64 // long
	BindingShotCenter bool
}

// 把 内存中 monster object 转成 protocol object monster
// 为了避免循环引用放在这里
func MonsterObjectToObjectMonster(obj com.MonsterObject) *ObjectMonster {
	return &ObjectMonster{
		ObjectID:          obj.ObjectID,           // uint32
		Name:              obj.Name,               // string
		NameColour:        4294967295,             // uint32 // Color
		Location:          obj.CurrentLocation,    // com.Point
		Image:             com.Monster(obj.Image), // com.Monster
		Direction:         obj.Direction,          // com.MirDirection
		Effect:            1,                      // TODO byte
		AI:                1,                      // TODO byte
		Light:             byte(0),                // byte
		Dead:              false,                  // bool
		Skeleton:          false,                  // bool
		Poison:            com.PoisonType(0),      // com.PoisonType
		Hidden:            false,                  // bool
		Extra:             false,                  // bool
		ExtraByte:         0,                      // byte
		ShockTime:         1,                      // uint64 // long
		BindingShotCenter: false,                  // bool
	}
}

func (self *ObjectMonster) ToBytes() []byte {
	// TODO test
	/*
	59, 0,
	173, 4, 0, 0,   object id
	4, 68, 101, 101, 114,    name
	255, 255, 255, 255,  color
	30, 1, 0, 0,    location
	82, 2, 0, 0,
	4, 0,   image
	4,   direction
	0,   effect
	2,   ai
	0,  light
	0,   dead
	0,  skeleton
	0, 0,  poison
	0,  hidden
	0,  extra
	0,  extra byte
	0, 0, 0, 0, 0, 0, 0, 0,   shock time
	0    binding shot
	*/
	pkgBytes := com.Uint16ToBytes(OBJECT_MONSTER)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
	nameBytes := com.StringToBytes(self.Name)
	nameColorBytes := com.Uint32ToBytes(self.NameColour)
	locationBytes := self.Location.ToBytes()
	imageBytes := com.Uint16ToBytes(uint16(self.Image))
	directionBytes := []byte{byte(self.Direction)}
	effectBytes := []byte{self.Effect}
	aiBytes := []byte{self.AI}
	lightBytes := []byte{self.Light}
	deadBytes := com.BoolToBytes(self.Dead)
	skeletonBytes := com.BoolToBytes(self.Skeleton)
	poisonBytes := com.Uint16ToBytes(uint16(self.Poison))
	hiddenBytes := com.BoolToBytes(self.Hidden)
	extraBytes := com.BoolToBytes(self.Extra)
	extraByteBytes := []byte{self.ExtraByte}
	shockTimeBytes := com.Uint64ToBytes(self.ShockTime)
	bindingShotCenterBytes := com.BoolToBytes(self.BindingShotCenter)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, nameBytes, nameColorBytes, locationBytes,
		imageBytes, directionBytes, effectBytes, aiBytes, lightBytes, deadBytes, skeletonBytes, poisonBytes,
		hiddenBytes, extraBytes, extraByteBytes, shockTimeBytes, bindingShotCenterBytes,
	} {
		log.Println(r)
		result = append(result, r...)
	}
	return result
}

type ObjectAttack struct {
	ObjectID  uint32
	Location  com.Point
	Direction com.MirDirection
	Spell     com.Spell
	Level     byte
	Type      byte
}

func (self *ObjectAttack) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(OBJECT_ATTACK)
	objectIdBytes := com.Uint32ToBytes(self.ObjectID)
	locationBytes := self.Location.ToBytes()
	directionBytes := []byte{byte(self.Direction)}
	spellBytes := []byte{byte(self.Spell)}
	levelBytes := []byte{byte(self.Level)}
	typeBytes := []byte{byte(self.Type)}
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, locationBytes, directionBytes, spellBytes, levelBytes, typeBytes} {
		result = append(result, r...)
	}
	return result
}

type Struck struct{}

func (self *Struck) ToBytes() []byte { return nil }

type ObjectStruck struct{}

func (self *ObjectStruck) ToBytes() []byte { return nil }

type DamageIndicator struct{}

func (self *DamageIndicator) ToBytes() []byte { return nil }

type DuraChanged struct{}

func (self *DuraChanged) ToBytes() []byte { return nil }

type HealthChanged struct {
	HP uint16
	MP uint16
}

func (self *HealthChanged) ToBytes() []byte {
	pkgBytes := com.Uint16ToBytes(HEALTH_CHANGED)
	hpBytes := com.Uint16ToBytes(self.HP)
	mpBytes := com.Uint16ToBytes(self.MP)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, hpBytes, mpBytes} {
		result = append(result, r...)
	}
	return result
}

type DeleteItem struct{}

func (self *DeleteItem) ToBytes() []byte { return nil }

type Death struct{}

func (self *Death) ToBytes() []byte { return nil }

type ObjectDied struct{}

func (self *ObjectDied) ToBytes() []byte { return nil }

type ColourChanged struct{}

func (self *ColourChanged) ToBytes() []byte { return nil }

type ObjectColourChanged struct{}

func (self *ObjectColourChanged) ToBytes() []byte { return nil }

type ObjectGuildNameChanged struct{}

func (self *ObjectGuildNameChanged) ToBytes() []byte { return nil }

type GainExperience struct{}

func (self *GainExperience) ToBytes() []byte { return nil }

type LevelChanged struct{}

func (self *LevelChanged) ToBytes() []byte { return nil }

type ObjectLeveled struct{}

func (self *ObjectLeveled) ToBytes() []byte { return nil }

type ObjectHarvest struct{}

func (self *ObjectHarvest) ToBytes() []byte { return nil }

type ObjectHarvested struct{}

func (self *ObjectHarvested) ToBytes() []byte { return nil }

type ObjectNpc struct{}

func (self *ObjectNpc) ToBytes() []byte { return nil }

type NPCResponse struct{}

func (self *NPCResponse) ToBytes() []byte { return nil }

type ObjectHide struct{}

func (self *ObjectHide) ToBytes() []byte { return nil }

type ObjectShow struct{}

func (self *ObjectShow) ToBytes() []byte { return nil }

type Poisoned struct{}

func (self *Poisoned) ToBytes() []byte { return nil }

type ObjectPoisoned struct{}

func (self *ObjectPoisoned) ToBytes() []byte { return nil }

type MapChanged struct{}

func (self *MapChanged) ToBytes() []byte { return nil }

type ObjectTeleportOut struct{}

func (self *ObjectTeleportOut) ToBytes() []byte { return nil }

type ObjectTeleportIn struct{}

func (self *ObjectTeleportIn) ToBytes() []byte { return nil }

type TeleportIn struct{}

func (self *TeleportIn) ToBytes() []byte { return nil }

type NPCGoods struct{}

func (self *NPCGoods) ToBytes() []byte { return nil }

type NPCSell struct{}

func (self *NPCSell) ToBytes() []byte { return nil }

type NPCRepair struct{}

func (self *NPCRepair) ToBytes() []byte { return nil }

type NPCSRepair struct{}

func (self *NPCSRepair) ToBytes() []byte { return nil }

type NPCRefine struct{}

func (self *NPCRefine) ToBytes() []byte { return nil }

type NPCCheckRefine struct{}

func (self *NPCCheckRefine) ToBytes() []byte { return nil }

type NPCCollectRefine struct{}

func (self *NPCCollectRefine) ToBytes() []byte { return nil }

type NPCReplaceWedRing struct{}

func (self *NPCReplaceWedRing) ToBytes() []byte { return nil }

type NPCStorage struct{}

func (self *NPCStorage) ToBytes() []byte { return nil }

type SellItem struct{}

func (self *SellItem) ToBytes() []byte { return nil }

type CraftItem struct{}

func (self *CraftItem) ToBytes() []byte { return nil }

type RepairItem struct{}

func (self *RepairItem) ToBytes() []byte { return nil }

type ItemRepaired struct{}

func (self *ItemRepaired) ToBytes() []byte { return nil }

type NewMagic struct{}

func (self *NewMagic) ToBytes() []byte { return nil }

type RemoveMagic struct{}

func (self *RemoveMagic) ToBytes() []byte { return nil }

type MagicLeveled struct{}

func (self *MagicLeveled) ToBytes() []byte { return nil }

type Magic struct{}

func (self *Magic) ToBytes() []byte { return nil }
