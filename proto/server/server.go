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
	Price            uint32
	StackSize        uint32 //default 1;
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
	indexBytes := cm.Uint32ToBytes(self.Index)
	nameBytes := cm.StringToBytes(self.Name)
	typeBytes := []byte{byte(self.Type)}
	gradeBytes := []byte{byte(self.Grade)}
	requiredTypeBytes := []byte{byte(self.RequiredType)}
	requiredClassBytes := []byte{byte(self.RequiredClass)}
	requiredGenderBytes := []byte{byte(self.RequiredGender)}
	setBytes := []byte{byte(self.Set)}
	shapeBytes := cm.Uint16ToBytes(self.Shape)
	weightBytes := []byte{self.Weight}
	lightBytes := []byte{self.Light}
	requiredAmountBytes := []byte{self.RequiredAmount}
	imageBytes := cm.Uint16ToBytes(self.Image)
	durabilityBytes := cm.Uint16ToBytes(self.Durability)
	priceBytes := cm.Uint32ToBytes(self.Price)
	stackSizeBytes := cm.Uint32ToBytes(self.StackSize)
	minACBytes := []byte{self.MinAC}
	maxACBytes := []byte{self.MaxAC}
	minMACBytes := []byte{self.MinMAC}
	maxMACBytes := []byte{self.MaxMAC}
	minDCBytes := []byte{self.MinDC}
	maxDCBytes := []byte{self.MaxDC}
	minMCBytes := []byte{self.MinMC}
	maxMCBytes := []byte{self.MaxMC}
	minSCBytes := []byte{self.MinSC}
	maxSCBytes := []byte{self.MaxSC}
	accuracyBytes := []byte{self.Accuracy}
	agilityBytes := []byte{self.Agility}
	hPBytes := cm.Uint16ToBytes(self.HP)
	mPBytes := cm.Uint16ToBytes(self.MP)
	attackSpeedBytes := []byte{byte(self.AttackSpeed)} // TODO int8 可能为负数，先当正数处理
	luckBytes := []byte{byte(self.Luck)}               // 同上
	bagWeightBytes := []byte{self.BagWeight}
	handWeightBytes := []byte{self.HandWeight}
	wearWeightBytes := []byte{self.WearWeight}
	startItemBytes := cm.BoolToBytes(self.StartItem)
	effectBytes := []byte{self.Effect}
	strongBytes := []byte{self.Strong}
	magicResistBytes := []byte{self.MagicResist}
	poisonResistBytes := []byte{self.PoisonResist}
	healthRecoveryBytes := []byte{self.HealthRecovery}
	spellRecoveryBytes := []byte{self.SpellRecovery}
	poisonRecoveryBytes := []byte{self.PoisonRecovery}
	hPrateBytes := []byte{self.HPrate}
	mPrateBytes := []byte{self.MPrate}
	criticalRateBytes := []byte{self.CriticalRate}
	criticalDamageBytes := []byte{self.CriticalDamage}
	needIdentifyBytes := cm.BoolToBytes(self.NeedIdentify)
	showGroupPickupBytes := cm.BoolToBytes(self.ShowGroupPickup)
	globalDropNotifyBytes := cm.BoolToBytes(self.GlobalDropNotify)
	classBasedBytes := cm.BoolToBytes(self.ClassBased)
	levelBasedBytes := cm.BoolToBytes(self.LevelBased)
	canMineBytes := cm.BoolToBytes(self.CanMine)
	canFastRunBytes := cm.BoolToBytes(self.CanFastRun)
	canAwakeningBytes := cm.BoolToBytes(self.CanAwakening)
	maxAcRateBytes := []byte{self.MaxAcRate}
	maxMacRateBytes := []byte{self.MaxMacRate}
	holyBytes := []byte{self.Holy}
	freezingBytes := []byte{self.Freezing}
	poisonAttackBytes := []byte{self.PoisonAttack}
	hpDrainRateBytes := []byte{self.HpDrainRate}
	bindBytes := cm.Uint16ToBytes(self.Bind)
	reflectBytes := []byte{self.Reflect}
	uniqueBytes := cm.Uint16ToBytes(self.Unique)
	randomStatsIdBytes := []byte{self.RandomStatsId}
	randomStatsBytes := self.RandomStats.ToBytes()
	toolTipBytes := cm.StringToBytes(self.ToolTip)
	result := make([]byte, 0)
	for _, r := range [][]byte{indexBytes, nameBytes, typeBytes, gradeBytes, requiredTypeBytes, requiredClassBytes,
		requiredGenderBytes, setBytes, shapeBytes, weightBytes, lightBytes, requiredAmountBytes, imageBytes,
		durabilityBytes, priceBytes, stackSizeBytes, minACBytes, maxACBytes, minMACBytes, maxMACBytes, minDCBytes,
		maxDCBytes, minMCBytes, maxMCBytes, minSCBytes, maxSCBytes, accuracyBytes, agilityBytes, hPBytes, mPBytes,
		attackSpeedBytes, luckBytes, bagWeightBytes, handWeightBytes, wearWeightBytes, startItemBytes, effectBytes,
		strongBytes, magicResistBytes, poisonResistBytes, healthRecoveryBytes, spellRecoveryBytes, poisonRecoveryBytes,
		hPrateBytes, mPrateBytes, criticalRateBytes, criticalDamageBytes, needIdentifyBytes, showGroupPickupBytes,
		globalDropNotifyBytes, classBasedBytes, levelBasedBytes, canMineBytes, canFastRunBytes, canAwakeningBytes,
		maxAcRateBytes, maxMacRateBytes, holyBytes, freezingBytes, poisonAttackBytes, hpDrainRateBytes, bindBytes,
		reflectBytes, uniqueBytes, randomStatsIdBytes, randomStatsBytes, toolTipBytes,
	} {
		result = append(result, r...)
	}
	return result
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

func (self *ClientMagic) ToBytes() []byte {
	spellBytes := []byte{byte(self.Spell)}
	baseCostBytes := []byte{byte(self.BaseCost)}
	levelCostBytes := []byte{byte(self.LevelCost)}
	iconBytes := []byte{byte(self.Icon)}
	level1Bytes := []byte{byte(self.Level1)}
	level2Bytes := []byte{byte(self.Level2)}
	level3Bytes := []byte{byte(self.Level3)}
	need1Bytes := cm.Uint16ToBytes(self.Need1)
	need2Bytes := cm.Uint16ToBytes(self.Need2)
	need3Bytes := cm.Uint16ToBytes(self.Need3)
	levelBytes := []byte{byte(self.Level)}
	keyBytes := []byte{byte(self.Key)}
	rangeBytes := []byte{byte(self.Range)}
	experienceBytes := cm.Uint16ToBytes(self.Experience)
	isTempSpellBytes := cm.BoolToBytes(self.IsTempSpell)
	castTimeBytes := cm.Uint64ToBytes(self.CastTime)
	delayBytes := cm.Uint64ToBytes(self.Delay)
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
	minimalFullnessBytes := cm.Uint32ToBytes(self.MinimalFullness)
	mousePickupEnabledBytes := cm.BoolToBytes(self.MousePickupEnabled)
	mousePickupRangeBytes := cm.Uint32ToBytes(self.MousePickupRange)
	autoPickupEnabledBytes := cm.BoolToBytes(self.AutoPickupEnabled)
	autoPickupRangeBytes := cm.Uint32ToBytes(self.AutoPickupRange)
	semiAutoPickupEnabledBytes := cm.BoolToBytes(self.SemiAutoPickupEnabled)
	semiAutoPickupRangeBytes := cm.Uint32ToBytes(self.SemiAutoPickupRange)
	canProduceBlackStoneBytes := cm.BoolToBytes(self.CanProduceBlackStone)
	infoBytes := cm.StringToBytes(self.Info)
	info1Bytes := cm.StringToBytes(self.Info1)
	info2Bytes := cm.StringToBytes(self.Info2)
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

func (self *IntelligentCreatureItemFilter) ToBytes() []byte {
	petPickupAllBytes := cm.BoolToBytes(self.PetPickupAll)
	petPickupGoldBytes := cm.BoolToBytes(self.PetPickupGold)
	petPickupWeaponsBytes := cm.BoolToBytes(self.PetPickupWeapons)
	petPickupArmoursBytes := cm.BoolToBytes(self.PetPickupArmours)
	petPickupHelmetsBytes := cm.BoolToBytes(self.PetPickupHelmets)
	petPickupBootsBytes := cm.BoolToBytes(self.PetPickupBoots)
	petPickupBeltsBytes := cm.BoolToBytes(self.PetPickupBelts)
	petPickupAccessoriesBytes := cm.BoolToBytes(self.PetPickupAccessories)
	petPickupOthersBytes := cm.BoolToBytes(self.PetPickupOthers)
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
	PetType          cm.IntelligentCreatureType
	Icon             uint32
	CustomName       string
	Fullness         uint32
	SlotIndex        uint32
	ExpireTime       uint64                           // long
	BlackstoneTime   uint64                           // long
	MaintainFoodTime uint64                           // long
	PetMode          cm.IntelligentCreaturePickupMode // default SemiAutomatic
	CreatureRules    IntelligentCreatureRules
	Filter           IntelligentCreatureItemFilter
}

func (self *ClientIntelligentCreature) ToBytes() []byte {
	petTypeBytes := []byte{byte(self.PetType)}
	iconBytes := cm.Uint32ToBytes(self.Icon)
	customNameBytes := cm.StringToBytes(self.CustomName)
	fullnessBytes := cm.Uint32ToBytes(self.Fullness)
	slotIndexBytes := cm.Uint32ToBytes(self.SlotIndex)
	expireTimeBytes := cm.Uint64ToBytes(self.ExpireTime)
	blackstoneTimeBytes := cm.Uint64ToBytes(self.BlackstoneTime)
	maintainFoodTimeBytes := cm.Uint64ToBytes(self.MaintainFoodTime)
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
	Info ItemInfo
}

// TODO
func (self *NewItemInfo) ToBytes() []byte {
	return nil
}

type MoveItem struct {
	Grid    cm.MirGridType
	From    uint32
	To      uint32
	Success bool
}

func (self *MoveItem) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(MOVE_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	fromBytes := cm.Uint32ToBytes(self.From)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type EquipItem struct {
	Grid     cm.MirGridType
	UniqueID uint64
	To       uint32
	Success  bool
}

func (self *EquipItem) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(EQUIP_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	uniqueIdBytes := cm.Uint64ToBytes(self.UniqueID)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, uniqueIdBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type MergeItem struct {
	GridFrom cm.MirGridType
	GridTo   cm.MirGridType
	IDFrom   uint64
	IDTo     uint64
	Success  bool
}

func (self *MergeItem) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(MERGE_ITEM)
	gridFromBytes := []byte{byte(self.GridFrom)}
	gridToBytes := []byte{byte(self.GridTo)}
	idFromBytes := cm.Uint64ToBytes(self.IDFrom)
	idToBytes := cm.Uint64ToBytes(self.IDTo)
	successBytes := cm.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridFromBytes, gridToBytes, idFromBytes, idToBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type RemoveItem struct {
	Grid     cm.MirGridType
	UniqueID uint64
	To       uint32
	Success  bool
}

func (self *RemoveItem) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(REMOVE_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	uniqueIdBytes := cm.Uint64ToBytes(self.UniqueID)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, gridBytes, uniqueIdBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type RemoveSlotItem struct {
	Grid     cm.MirGridType
	GridTo   cm.MirGridType
	UniqueID uint64
	To       uint32
	Success  bool
}

func (self *RemoveSlotItem) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(REMOVE_SLOT_ITEM)
	gridBytes := []byte{byte(self.Grid)}
	gridToBytes := []byte{byte(self.GridTo)}
	uniqueIdBytes := cm.Uint64ToBytes(self.UniqueID)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(TAKE_BACK_ITEM)
	fromBytes := cm.Uint32ToBytes(self.From)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(STORE_ITEM)
	fromBytes := cm.Uint32ToBytes(self.From)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, fromBytes, toBytes, successBytes} {
		result = append(result, r...)
	}
	return result
}

type SplitItem struct {
	Item UserItem
	Grid cm.MirGridType
}

// TODO
func (self SplitItem) ToBytes() []byte {
	//pkgBytes := cm.Uint16ToBytes(SPLIT_ITEM)
	return nil
}

type SplitItem1 struct {
	Grid     cm.MirGridType
	UniqueID uint64
	Count    uint32
	Success  bool
}

func (self *SplitItem1) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(SPLIT_ITEM1)
	gridBytes := []byte{byte(self.Grid)}
	uniqueIdBytes := cm.Uint64ToBytes(self.UniqueID)
	countBytes := cm.Uint32ToBytes(self.Count)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(DEPOSIT_REFINE_ITEM)
	fromBytes := cm.Uint32ToBytes(self.From)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(RETRIEVE_REFINE_ITEM)
	fromBytes := cm.Uint32ToBytes(self.From)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(REFINE_CANCEL)
	unlockBytes := cm.BoolToBytes(self.Unlock)
	return append(pkgBytes, unlockBytes...)
}

type RefineItem struct {
	UniqueID uint64
}

func (self *RefineItem) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(REFINE_ITEM)
	uniqueIdBytes := cm.Uint64ToBytes(self.UniqueID)
	return append(pkgBytes, uniqueIdBytes...)
}

type DepositTradeItem struct {
	From    uint32
	To      uint32
	Success bool
}

func (self *DepositTradeItem) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(DEPOSIT_TRADE_ITEM)
	fromBytes := cm.Uint32ToBytes(self.From)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(RETRIEVE_TRADE_ITEM)
	fromBytes := cm.Uint32ToBytes(self.From)
	toBytes := cm.Uint32ToBytes(self.To)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(USE_ITEM)
	uniqueIdBytes := cm.Uint64ToBytes(self.UniqueID)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(DROP_ITEM)
	uniqueIdBytes := cm.Uint64ToBytes(self.UniqueID)
	countBytes := cm.Uint32ToBytes(self.Count)
	successBytes := cm.BoolToBytes(self.Success)
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
	pkgBytes := cm.Uint16ToBytes(PLAYER_UPDATE)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
	lightBytes := []byte{self.Light}
	weaponBytes := cm.Uint16ToBytes(self.Weapon)
	weaponEffectBytes := cm.Uint16ToBytes(self.WeaponEffect)
	armourBytes := cm.Uint16ToBytes(self.Armour)
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
	Location          cm.Point
	Image             cm.Monster
	Direction         cm.MirDirection
	Effect            byte
	AI                byte
	Light             byte
	Dead              bool
	Skeleton          bool
	Poison            cm.PoisonType
	Hidden            bool
	Extra             bool
	ExtraByte         byte
	ShockTime         uint64 // long
	BindingShotCenter bool
}

func (self *ObjectMonster) ToBytes() []byte {
	// TODO test
	pkgBytes := cm.Uint16ToBytes(OBJECT_MONSTER)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
	nameBytes := cm.StringToBytes(self.Name)
	nameColorBytes := cm.Uint32ToBytes(self.NameColour)
	locationBytes := self.Location.ToBytes()
	imageBytes := []byte{byte(self.Image)}
	directionBytes := []byte{byte(self.Direction)}
	effectBytes := []byte{self.Effect}
	aiBytes := []byte{self.AI}
	lightBytes := []byte{self.Light}
	deadBytes := cm.BoolToBytes(self.Dead)
	skeletonBytes := cm.BoolToBytes(self.Skeleton)
	poisonBytes := []byte{byte(self.Poison)}
	hiddenBytes := cm.BoolToBytes(self.Hidden)
	extraBytes := cm.BoolToBytes(self.Extra)
	extraByteBytes := []byte{self.ExtraByte}
	shockTimeBytes := cm.Uint64ToBytes(self.ShockTime)
	bindingShotCenterBytes := cm.BoolToBytes(self.BindingShotCenter)
	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, objectIdBytes, nameBytes, nameColorBytes, locationBytes,
		imageBytes, directionBytes, effectBytes, aiBytes, lightBytes, deadBytes, skeletonBytes, poisonBytes,
		hiddenBytes, extraBytes, extraByteBytes, shockTimeBytes, bindingShotCenterBytes,
	} {
		result = append(result, r...)
	}
	return result
}

type ObjectAttack struct {
	ObjectID  uint32
	Location  cm.Point
	Direction cm.MirDirection
	Spell     cm.Spell
	Level     byte
	Type      byte
}

func (self *ObjectAttack) ToBytes() []byte {
	pkgBytes := cm.Uint16ToBytes(OBJECT_ATTACK)
	objectIdBytes := cm.Uint32ToBytes(self.ObjectID)
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
	pkgBytes := cm.Uint16ToBytes(HEALTH_CHANGED)
	hpBytes := cm.Uint16ToBytes(self.HP)
	mpBytes := cm.Uint16ToBytes(self.MP)
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
