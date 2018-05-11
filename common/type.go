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

type BindMode uint16

const (
	BM_NONE                  BindMode = iota
	BM_DONT_DEATHDROP                 = 1
	BM_DONT_DROP                      = 2
	BM_DONT_SELL                      = 4
	BM_DONT_STORE                     = 8
	BM_DONT_TRADE                     = 16
	BM_DONT_REPAIR                    = 32
	BM_DONT_UPGRADE                   = 64
	BM_DESTROY_ON_DROP                = 128
	BM_BREAK_ON_DEATH                 = 256
	BM_BIND_ON_EQUIP                  = 512
	BM_NO_S_REPAIR                    = 1024
	BM_NO_WEDDING_RING                = 2048
	BM_UNABLE_TO_RENT                 = 4096
	BM_UNABLE_TO_DISASSEMBLE          = 8192
)

type SpecialItemMode uint16

// TODO 转成十进制

const (
	SIM_NONE         SpecialItemMode = iota
	SIM_PARALIZE                     = 0x0001
	SIM_TELEPORT                     = 0x0002
	SIM_CLEARRING                    = 0x0004
	SIM_PROTECTION                   = 0x0008
	SIM_REVIVAL                      = 0x0010
	SIM_MUSCLE                       = 0x0020
	SIM_FLAME                        = 0x0040
	SIM_HEALING                      = 0x0080
	SIM_PROBE                        = 0x0100
	SIM_SKILL                        = 0x0200
	SIM_NO_DURA_LOSS                 = 0x0400
	SIM_BLINK                        = 0x800
)

type LevelEffects byte

const (
	LE_NONE        LevelEffects = iota
	LE_MIST                     = 0x0001
	LE_RED_DRAGON               = 0x0002
	LE_BLUE_DRAGON              = 0x0004
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

const (
	S_NONE Spell = iota

	//Warrior
	S_FENCING          = 1
	S_SLAYING          = 2
	S_THRUSTING        = 3
	S_HALF_MOON        = 4
	S_SHOULDER_DASH    = 5
	S_TWIN_DRAKE_BLADE = 6
	S_ENTRAPMENT       = 7
	S_FLAMING_SWORD    = 8
	S_LION_ROAR        = 9
	S_CROSS_HALF_MOON  = 10
	S_BLADE_AVALANCHE  = 11
	S_PROTECTION_FIELD = 12
	S_RAGE             = 13
	S_COUNTER_ATTACK   = 14
	S_SLASHING_BURST   = 15
	S_FURY             = 16
	S_IMMORTAL_SKIN    = 17

	//Wizard
	S_FIRE_BALL       = 31
	S_REPULSION       = 32
	S_ELECTRIC_SHOCK  = 33
	S_GREAT_FIRE_BALL = 34
	S_HELL_FIRE       = 35
	S_THUNDER_BOLT    = 36
	S_TELEPORT        = 37
	S_FIRE_BANG       = 38
	S_FIRE_WALL       = 39
	S_LIGHTNING       = 40
	S_FROST_CRUNCH    = 41
	S_THUNDER_STORM   = 42
	S_MAGIC_SHIELD    = 43
	S_TURN_UNDEAD     = 44
	S_VAMPIRISM       = 45
	S_ICE_STORM       = 46
	S_FLAME_DISRUPTOR = 47
	S_MIRRORING       = 48
	S_FLAME_FIELD     = 49
	S_BLIZZARD        = 50
	S_MAGIC_BOOSTER   = 51
	S_METEOR_STRIKE   = 52
	S_ICE_THRUST      = 53
	S_FAST_MOVE       = 54
	S_STORM_ESCAPE    = 55

	//Taoist
	S_HEALING           = 61
	S_SPIRIT_SWORD      = 62
	S_POISONING         = 63
	S_SOUL_FIRE_BALL    = 64
	S_SUMMON_SKELETON   = 65
	S_HIDING            = 67
	S_MASS_HIDING       = 68
	S_SOUL_SHIELD       = 69
	S_REVELATION        = 70
	S_BLESSED_ARMOUR    = 71
	S_ENERGY_REPULSOR   = 72
	S_TRAP_HEXAGON      = 73
	S_PURIFICATION      = 74
	S_MASS_HEALING      = 75
	S_HALLUCINATION     = 76
	S_ULTIMATE_ENHANCER = 77
	S_SUMMON_SHINSU     = 78
	S_REINCARNATION     = 79
	S_SUMMON_HOLY_DEVA  = 80
	S_CURSE             = 81
	S_PLAGUE            = 82
	S_POISON_CLOUD      = 83
	S_ENERGY_SHIELD     = 84
	S_PET_ENHANCER      = 85
	S_HEALING_CIRCLE    = 86

	//Assassin

	//Archer

	//Custom
	S_BLINK      = 151
	S_PORTAL     = 152
	S_BATTLE_Cry = 153

	//Map Events
	S_DIG_OUT_ZOMBIE = 200
	S_RUBBLE         = 201
	S_MAP_LIGHTNING  = 202
	S_MAP_LAVA       = 203
	S_MAP_QUAKE1     = 204
	S_MAP_QUAKE2     = 205
)

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

type Monster uint16

const (
	GUARD          Monster = iota
	TAOIST_GUARD
	GUARD2
	HEN
	DEER
	SCARECROW
	HOOKING_CAT
	RAKING_CAT
	YOB
	OMA
	CANNIBAL_PLANT
	// TODO ......
)

type PoisonType uint16

const (
	PT_None             PoisonType = iota
	PT_Green                       = 1
	PT_Red                         = 2
	PT_Slow                        = 4
	PT_Frozen                      = 8
	PT_Stun                        = 16
	PT_Paralysis                   = 32
	PT_DelayedExplosion            = 64
	PT_Bleeding                    = 128
	PT_LRParalysis                 = 256
)
