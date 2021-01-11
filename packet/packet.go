package packet

type Payload interface {
	ToByte() []byte
}

type Packet struct {
}
