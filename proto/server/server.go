package server

import (
	"mir/util"
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
	bytes = append(bytes, self.Result)
	return bytes
}
