package request

import "github.com/int2xx9/rcs620s/packet"

type BRTYType byte

const (
	BRTY212 BRTYType = 0x01
	BRTY424 BRTYType = 0x02
)

type InListPassiveTarget struct {
	RequestPacket
	MaxTg         byte
	BRTY          BRTYType
	InitiatorData []byte
}

func (i *InListPassiveTarget) ToByte() []byte {
	d := []byte{i.CommandCode, i.SubCommandCode, i.MaxTg, byte(i.BRTY)}
	return append(d, i.InitiatorData...)
}

func NewInListPassiveTarget(brty BRTYType, initiatorData []byte) (packet.Payload, error) {
	return &InListPassiveTarget{
		RequestPacket: RequestPacket{
			CommandCode:    0xd4,
			SubCommandCode: 0x4a,
		},
		MaxTg:         0x01,
		BRTY:          brty,
		InitiatorData: initiatorData,
	}, nil
}
