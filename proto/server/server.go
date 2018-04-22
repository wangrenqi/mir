package server

import (
	"mir/util"
	"mir/orm"
	"encoding/binary"
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
	bytes := util.IndexToBytes(CONNECTED)
	return bytes
}

type ClientVersion struct {
	// 0 wrong version
	// 1 correct version
	Result byte
}

func (self *ClientVersion) ToBytes() []byte {
	bytes := util.IndexToBytes(CLIENT_VERSION)
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
	bytes := util.IndexToBytes(NEW_ACCOUNT)
	return append(bytes, self.Result)
}

type ChangePassword struct {
	Result byte
}

func (self *ChangePassword) ToBytes() []byte {
	bytes := util.IndexToBytes(NEW_ACCOUNT)
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
	bytes := util.IndexToBytes(LOGIN)
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
	bytes := util.IndexToBytes(LOGIN_SUCCESS)
	characters := self.Characters
	if len(characters) == 0 {
		bytes = append(bytes, []byte{0, 0, 0, 0}...)
	} else {
		// TODO 有角色时LoginSuccess处理
		for _, character := range characters {
			bytes = append(bytes, character.ToBytes()...)
		}
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
	bytes := util.IndexToBytes(NEW_CHARACTER)
	return append(bytes, self.Result)
}

type NewCharacterSuccess struct {
	CharInfo orm.SelectInfo
}

func (self *NewCharacterSuccess) ToBytes() []byte {
	// index(int32 4byte)
	pkgBytes := util.IndexToBytes(NEW_CHARACTER_SUCCESS)
	index := uint32(self.CharInfo.ID)
	indexBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBytes, index)
	// name (string)
	name := self.CharInfo.Name
	nameBytes := []byte(name)
	nameBytesLenBytes := []byte{byte(len(nameBytes))}
	nameBytes = append(nameBytesLenBytes, nameBytes...)
	// level (int16 2byte)
	level := uint16(self.CharInfo.Level)
	levelBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(levelBytes, level)
	// class (byte)
	class := self.CharInfo.Class
	classBytes := []byte{class}
	// gender (byte)
	gender := self.CharInfo.Gender
	genderBytes := []byte{gender}
	// lastAccess (int64 8byte)
	lastAccess := uint64(0)
	lastAccessBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lastAccessBytes, lastAccess)

	result := make([]byte, 0)
	for _, r := range [][]byte{pkgBytes, indexBytes, nameBytes, levelBytes, classBytes, genderBytes, lastAccessBytes} {
		result = append(result, r...)
	}
	return result
}
