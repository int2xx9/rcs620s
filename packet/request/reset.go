package request

import "github.com/int2xx9/rcs620s/packet"

type Reset struct {
	RequestPacket
	Function byte
}

func (r *Reset) ToByte() []byte {
	return []byte{r.CommandCode, r.SubCommandCode, r.Function}
}

func NewReset() (packet.Payload, error) {
	return &Reset{
		RequestPacket: RequestPacket{
			CommandCode:    0xd4,
			SubCommandCode: 0x18,
		},
		Function: 0x01,
	}, nil
}
