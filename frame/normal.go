package frame

import (
	"errors"
	"fmt"

	"github.com/int2xx9/rcs620s/packet"
)

var (
	ErrWrongDataLength = errors.New("wrong data length (must be <= 255 bytes)")
)

type NormalPayload []byte

func (n NormalPayload) ToByte() []byte {
	return n
}

type Normal struct {
	payload packet.Payload
}

func NewNormalFrame(data packet.Payload) (*Normal, error) {
	if len(data.ToByte()) > 255 {
		return nil, ErrWrongDataLength
	}
	return &Normal{
		payload: data,
	}, nil
}

func (n *Normal) Payload() packet.Payload {
	return n.payload
}

func (n *Normal) ToByte() []byte {
	data := n.payload.ToByte()
	datalen := byte(len(data))

	r := []byte{0x00, 0x00, 0xff, datalen, CalcLCS(datalen)}
	r = append(r, data...)
	r = append(r, CalcDCS(data), 0x00)
	return append(r)
}

func (n *Normal) String() string {
	return fmt.Sprintf("NormalFrame: %s", n.Payload())
}

func IsNormal(f Packet) bool {
	_, ok := f.(*Normal)
	return ok
}
