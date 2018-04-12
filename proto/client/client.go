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
