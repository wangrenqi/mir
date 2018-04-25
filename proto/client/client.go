package client

import (
	cm "mir/common"
)

const (
	CLIENT_VERSION   = iota
	DISCONNECT
	KEEPALIVE
	NEW_ACCOUNT
	CHANGE_PASSWORD
	LOGIN
	NEW_CHARACTER
	DELETE_CHARACTER
	START_GAME
	LOGOUT
	TURN
	WALK
	RUN
	CHAT
)

type ClientVersion struct {
	VersionHash string
}

func GetClientVersion(bytes []byte) *ClientVersion {

	return nil
}

func (self *ClientVersion) ToBytes() []byte {
	//24, 0 (22 + 2)
	return []byte{0, 0, 16, 0, 0, 0, 196, 46, 198, 6, 217, 38, 102, 128, 242, 128, 185, 164, 66, 146, 36, 34}
}

type Disconnect struct{}

func GetDisconnect(bytes []byte) *Disconnect {

	return nil
}

func (self *Disconnect) ToBytes() []byte {
	return nil
}

type KeepAlive struct {
	//time time.Time
}

func GetKeepAlive(bytes []byte) *KeepAlive {

	return nil
}

func (self *KeepAlive) ToBytes() []byte {
	return nil
}

type NewAccount struct {
	UserName string
	Password string
	//BirthDate      datetime
	SecretQuestion string
	SecretAnswer   string
	EMailAddress   string
}

func GetNewAccount(bytes []byte) *NewAccount {
	index, username := cm.ReadString(bytes, 0)
	index, password := cm.ReadString(bytes, index)
	// TODO birthday datetime from binary int64
	return &NewAccount{UserName: username, Password: password}
}

func (self *NewAccount) ToBytes() []byte {
	return nil
}

type ChangePassword struct {
	AccountId       string
	CurrentPassword string
	NewPassword     string
}

func GetChangePassword(bytes []byte) *ChangePassword {

	return nil
}

func (self *ChangePassword) ToBytes() []byte {
	return nil
}

type Login struct {
	AccountId string
	Password  string
}

func GetLogin(bytes []byte) *Login {
	index, username := cm.ReadString(bytes, 0)
	index, password := cm.ReadString(bytes, index)
	return &Login{AccountId: username, Password: password}
}

func (self *Login) ToBytes() []byte {
	//data := pkg.Data.(*cp.Login)
	// 15, 0 (13 + 2)
	return []byte{5, 0, 3, 50, 50, 50, 6, 50, 50, 50, 50, 50, 50}
}

type NewCharacter struct {
	Name   string
	Gender cm.MirGender
	Class  cm.MirClass
}

func GetNewCharacter(bytes []byte) *NewCharacter {
	index, name := cm.ReadString(bytes, 0)
	gender := bytes[index]
	class := bytes[index+1]
	return &NewCharacter{Name: name, Gender: cm.MirGender(gender), Class: cm.MirClass(class)}
}

func (self *NewCharacter) ToBytes() []byte {
	return nil
}

type DeleteCharacter struct {
	CharacterIndex int
}

func GetDeleteCharacter(bytes []byte) *DeleteCharacter {

	return nil
}

func (self *DeleteCharacter) ToBytes() []byte {
	return nil
}

type StartGame struct {
	CharacterIndex int
}

func GetStartGame(bytes []byte) *StartGame {

	return nil
}

func (self *StartGame) ToBytes() []byte {
	// 8, 0 (6 + 2)
	return []byte{8, 0, 2, 0, 0, 0}
}

type Logout struct{}

func GetLogout(bytes []byte) *Logout {

	return nil
}

func (self *Logout) ToBytes() []byte {

	return nil
}

type Turn struct {
	Dir cm.MirDirection
}

func GetTurn(bytes []byte) *Turn {

	return nil
}

func (self *Turn) ToBytes() []byte {
	return nil
}

type Walk struct {
	Dir cm.MirDirection
}

func GetWalk(bytes []byte) *Walk {

	return nil
}

func (self *Walk) ToBytes() []byte {
	// up upright right downright down downleft left upleft
	// 5, 0 (3 + 2)
	return []byte{11, 0, byte(self.Dir)}
}

type Run struct {
	Dir cm.MirDirection
}

func GetRun(bytes []byte) *Run {

	return nil
}

func (self *Run) ToBytes() []byte {
	return nil
}

type Chat struct {
	Message string
}

func GetChat(bytes []byte) *Chat {
	_, msg := cm.ReadString(bytes, 0)
	return &Chat{Message: msg}
}

func (self *Chat) ToBytes() []byte {
	msgBytes := []byte(self.Message)
	index := cm.IndexToBytes(CHAT)
	index = append(index, byte(len(msgBytes)))
	bytes := append(index, msgBytes...)
	return bytes
}
