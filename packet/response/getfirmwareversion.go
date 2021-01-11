package response

import (
	"encoding/binary"
	"fmt"

	"github.com/int2xx9/rcs620s/packet"
)

type GetFirmwareVersion struct {
	ResponsePacket
	ICType byte
	Ver    uint16
	NA     byte
}

func (g *GetFirmwareVersion) ToByte() []byte {
	return []byte{g.ResponseCode, g.SubResponseCode, g.ICType, byte(g.Ver >> 8), byte(g.Ver), g.NA}
}

func (g *GetFirmwareVersion) String() string {
	return fmt.Sprintf("response.GetFirmwareVersion: ICType:0x%02x, Ver:0x%04x, N/A:0x%02x", g.ICType, g.Ver, g.NA)
}

func ParseGetFirmwareVersion(data []byte) (packet.Payload, error) {
	return &GetFirmwareVersion{
		ResponsePacket: ResponsePacket{
			ResponseCode:    data[0],
			SubResponseCode: data[1],
		},
		ICType: data[2],
		Ver:    binary.BigEndian.Uint16([]byte{data[3], data[4]}),
		NA:     data[5],
	}, nil
}
