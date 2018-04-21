package client

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

type Direction byte

const (
	Up        Direction = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

type ClientVersion struct {
	VersionHash string
}

func GetClientVersion(bytes []byte) *ClientVersion {

}

type Disconnect struct{}

func GetDisconnect(bytes []byte) *Disconnect {

}

type KeepAlive struct {
	//time time.Time
}

func GetKeepAlive(bytes []byte) *KeepAlive {

}

type NewAccount struct {
	UserName string
	Password string
}

func GetNewAccount(bytes []byte) *NewAccount {

}

type ChangePassword struct {
	AccountId       string
	CurrentPassword string
	NewPassword     string
}

func GetChangePassword(bytes []byte) *ChangePassword {

}

type Login struct {
	AccountId string
	Password  string
}

func GetLogin(bytes []byte) *Login {

}

type NewCharacter struct {
	//Gender MirGender
	//Class  MirClass
}

func GetNewCharacter(bytes []byte) *NewCharacter {

}

type DeleteCharacter struct {
	CharacterIndex int
}

func GetDeleteCharacter(bytes []byte) *DeleteCharacter {

}

type StartGame struct {
	CharacterIndex int
}

func GetStartGame(bytes []byte) *StartGame {

}

type Logout struct{}

func GetLogout(bytes []byte) *Logout {

}

type Turn struct {
	Dir Direction
}

func GetTurn(bytes []byte) *Turn {

}

type Walk struct {
	Dir Direction
}

func GetWalk(bytes []byte) *Walk {

}

type Run struct {
	Dir Direction
}

func GetRun(bytes []byte) *Run {

}

type Chat struct {
	Message string
}

func GetChat(bytes []byte) *Chat {

}
