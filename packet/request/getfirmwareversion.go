package request

import "github.com/int2xx9/rcs620s/packet"

type GetFirmwareVersion struct {
	RequestPacket
}

func (g *GetFirmwareVersion) ToByte() []byte {
	return []byte{g.CommandCode, g.SubCommandCode}
}

func NewGetFirmwareVersion() (packet.Payload, error) {
	return &GetFirmwareVersion{
		RequestPacket: RequestPacket{
			CommandCode:    0xd4,
			SubCommandCode: 0x02,
		},
	}, nil
}
