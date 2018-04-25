package server

import (
	"mir/orm"
	cm "mir/common"
)

const (
	CONNECTED                = iota
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
	Resolution int
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
