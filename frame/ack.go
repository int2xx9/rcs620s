package frame

import "github.com/int2xx9/rcs620s/packet"

type Ack struct {
	payload packet.Payload
}

type AckPayload struct{}

func (*AckPayload) ToByte() []byte {
	return []byte{}
}

func NewAckFrame() (*Ack, error) {
	return &Ack{
		payload: &AckPayload{},
	}, nil
}

func (a *Ack) Payload() packet.Payload {
	return a.payload
}

func (*Ack) ToByte() []byte {
	return []byte{0x00, 0x00, 0xff, 0x00, 0xff, 0x00}
}

func IsAck(f Packet) bool {
	_, ok := f.(*Ack)
	return ok
}
