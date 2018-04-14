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

type Disconnect struct{}

//type KeepAlive struct {
//	time time.Time
//}

type NewAccount struct {
	UserName string
	Password string
}

type ChangePassword struct {
	AccountId       string
	CurrentPassword string
	NewPassword     string
}

type Login struct {
	AccountId string
	Password  string
}

type StartGame struct {
	CharacterIndex int
}

type Logout struct{}

type Turn struct {
	Dir Direction
}

type Walk struct {
	Dir Direction
}
