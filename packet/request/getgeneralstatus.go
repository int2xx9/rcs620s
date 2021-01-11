package request

import "github.com/int2xx9/rcs620s/packet"

type GetGeneralStatus struct {
	RequestPacket
}

func (g *GetGeneralStatus) ToByte() []byte {
	return []byte{g.CommandCode, g.SubCommandCode}
}

func NewGetGeneralStatus() (packet.Payload, error) {
	return &GetGeneralStatus{
		RequestPacket: RequestPacket{
			CommandCode:    0xd4,
			SubCommandCode: 0x04,
		},
	}, nil
}
