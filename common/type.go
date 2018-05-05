package common

type MirDirection byte

const (
	UP         MirDirection = iota
	UP_RIGHT
	RIGHT
	DOWN_RIGHT
	DOWN
	DOWN_LEFT
	LEFT
	UP_LEFT
)

type MirGender byte

const (
	MALE   MirGender = iota
	FEMALE
)

type MirClass byte

const (
	WARRIOR  MirClass = iota
	WIZARD
	TAOIST
	ASSASSIN
	ARCHER
)

// public enum LevelEffects : byte
// {
// 		None = 0,
// 		ist = 0x0001,
// 		RedDragon = 0x0002,
// 		BlueDragon = 0x0004
// }

type ItemType byte

const (
	IT_NOTHING           ItemType = iota
	IT_WEAPON
	IT_ARMOUR
	IT_HELMET
	IT_NECKLACE
	IT_BRACELET
	IT_RING
	IT_AMULET
	IT_BELT
	IT_BOOTS
	IT_STONE
	IT_TORCH
	IT_POTION
	IT_ORE
	IT_MEAT
	IT_CRAFTING_MATERIAL
	IT_SCROLL
	IT_GEM
	IT_MOUNT
	IT_BOOK
	IT_SCRIPT
	IT_REINS
	IT_BELLS
	IT_SADDLE
	IT_RIBBON
	IT_MASK
	IT_FOOD
	IT_HOOK
	IT_FLOAT
	IT_BAIT
	IT_FINDER
	IT_REEL
	IT_FISH
	IT_QUEST
	IT_AWAKENING
	IT_PETS
	IT_TRANSFORM
)

type RequiredType byte

const (
	RT_LEVEL     RequiredType = iota
	RT_MAX_AC
	RT_MAX_MAC
	RT_MAX_DC
	RT_MAX_MC
	RT_MAX_SC
	RT_MAX_LEVEL
	RT_MIN_AC
	RT_MIN_MAC
	RT_MIN_DC
	RT_MIN_MC
	RT_MIN_SC
)

type RequiredClass byte

const (
	_          RequiredClass = iota
	RC_WARRIOR  // 1  // 因为与上面MirClass 重名 加上前缀RC
	RC_WIZARD   // 2
	_           // 3
	RC_TAOIST   // 4
	//Assassin  // 8
	//Archer  	// 16
)

type RequiredGender byte

const (
	_         RequiredGender = iota
	RG_MALE    // 1  // 因为与上面MirGender 重名 加上前缀RG
	RG_FEMALE  // 2
)

type ItemSet byte

const (
	IS_NONE         ItemSet = iota
	IS_SPIRIT
	IS_RECALL
	IS_RED_ORCHID
	IS_RED_FLOWER
	IS_SMASH
	IS_HWAN_DEVIL
	IS_PURITY
	IS_FIVE_STRING
	IS_MUNDANE
	IS_NOK_CHI
	IS_TAO_PROTECT
	IS_MIR
	IS_BONE
	IS_BUG
	IS_WHITE_GOLD
	IS_WHITE_GOLD_H
	IS_RED_JADE
	IS_RED_JADE_H
	IS_NEPHRITE
	IS_NEPHRITE_H
	IS_WHISKER_1
	IS_WHISKER_2
	IS_WHISKER_3
	IS_WHISKER_4
	IS_WHISKER_5
	IS_HYEOLRYONG
	IS_MONITOR
	IS_OPPRESSIVE
	IS_PAEOK
	IS_SULGWAN
)

// type BindMode uint16

const (
	BM_NONE                  = 0
	BM_DONT_DEATHDROP        = 1
	BM_DONT_DROP             = 2
	BM_DONT_SELL             = 4
	BM_DONT_STORE            = 8
	BM_DONT_TRADE            = 16
	BM_DONT_REPAIR           = 32
	BM_DONT_UPGRADE          = 64
	BM_DESTROY_ON_DROP       = 128
	BM_BREAK_ON_DEATH        = 256
	BM_BIND_ON_EQUIP         = 512
	BM_NO_S_REPAIR           = 1024
	BM_NO_WEDDING_RING       = 2048
	BM_UNABLE_TO_RENT        = 4096
	BM_UNABLE_TO_DISASSEMBLE = 8192
)

// type SpecialItemMode uint16
// TODO 转成十进制

const (
	SIM_NONE         = 0
	SIM_PARALIZE     = 0x0001
	SIM_TELEPORT     = 0x0002
	SIM_CLEARRING    = 0x0004
	SIM_PROTECTION   = 0x0008
	SIM_REVIVAL      = 0x0010
	SIM_MUSCLE       = 0x0020
	SIM_FLAME        = 0x0040
	SIM_HEALING      = 0x0080
	SIM_PROBE        = 0x0100
	SIM_SKILL        = 0x0200
	SIM_NO_DURA_LOSS = 0x0400
	SIM_BLINK        = 0x800
)

//type LevelEffects byte

const (
	LE_NONE        = 0
	LE_MIST        = 0x0001
	LE_RED_DRAGON  = 0x0002
	LE_BLUE_DRAGON = 0x0004
)

type IntelligentCreatureType byte

const
(
	ICT_BABY_PIG      IntelligentCreatureType = iota
	ICT_CHICK
	ICT_KITTEN
	ICT_BABY_SKELETON
	ICT_BAEKDON
	ICT_WIMAEN
	ICT_BLACK_KITTEN
	ICT_BABY_DRAGON
	ICT_OLYMPIC_FLAME
	ICT_BABY_SNOW_MAN
	ICT_FROG
	ICT_BABY_MONKEY
	ICT_ANGRY_BIRD
	ICT_FOXEY
	ICT_NONE          = 99
)

type IntelligentCreaturePickupMode byte

const (
	ICP_Automatic     IntelligentCreaturePickupMode = iota
	ICP_SemiAutomatic
)

type ChatType byte

const (
	CT_NORMAL       ChatType = iota
	CT_SHOUT
	CT_SYSTEM
	CT_HINT
	CT_ANNOUNCEMENT
	CT_GROUP
	CT_WHISPERiN
	CT_WHISPERoUT
	CT_GUILD
	CT_TRAINER
	CT_LEVELuP
	CT_SYSTEM2
	CT_RELATIONSHIP
	CT_MENTOR
	CT_SHOUT2
	CT_SHOUT3
)

type ItemGrade byte

const (
	IG_NONE      ItemGrade = iota
	IG_COMMON
	IG_RARE
	IG_LEGENDARY
	IG_MYTHICAL
)

type RefinedValue byte

const (
	RV_NONE RefinedValue = iota
	RV_DC
	RV_MC
	RV_SC
)

// TODO 这个太多
type Spell byte

type MirGridType byte

const (
	MGT_NONE            MirGridType = iota
	MGT_INVENTORY
	MGT_EQUIPMENT
	MGT_TRADE
	MGT_STORAGE
	MGT_BUY_BACK
	MGT_DROP_PANEL
	MGT_INSPECT
	MGT_TRUST_MERCHANT
	MGT_GUILD_STORAGE
	MGT_GUEST_TRADE
	MGT_MOUNT
	MGT_FISHING
	MGT_QUEST_INVENTORY
	MGT_AWAKEN_ITEM
	MGT_MAIL
	MGT_REFINE
	MGT_RENTING
	MGT_GUEST_RENTING
	MGT_CRAFT
)
