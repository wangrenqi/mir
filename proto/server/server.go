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

func GetConnected(bytes []byte) *Connected {
	return &Connected{}
}

func (self *Connected) ToBytes() []byte {
	bytes := util.IndexToBytes(CONNECTED)
	return bytes
}

type ClientVersion struct{}

func GetClientVersion(bytes []byte) *ClientVersion {
	return &ClientVersion{}
}

func (self *ClientVersion) ToBytes() []byte {
	bytes := util.IndexToBytes(CLIENT_VERSION)
	bytes = append(bytes, byte(1))
	return bytes
}
