package proto

type Packet struct {
	IsServer bool
	Index    int
	Data     interface{}
}

type NewAccount struct {
	UserName string
	Password string
}

const (
	CLIENT_VERSION = iota
	DISCONNECT
	KEEPALIVE
	NEW_ACCOUNT
)

func ToPacket([]byte) *Packet {

}

func (pkg *Packet) ToBytes() []byte {

}
