package response

import (
	"fmt"

	"github.com/int2xx9/rcs620s/packet"
)

type InListPassiveTarget struct {
	ResponsePacket
	NbTg       byte
	TargetData []byte
}

func (i *InListPassiveTarget) ToByte() []byte {
	d := []byte{i.ResponseCode, i.SubResponseCode, i.NbTg}
	return append(d, i.TargetData...)
}

func (i *InListPassiveTarget) String() string {
	return fmt.Sprintf("response.InListPassiveTarget: NbTg:%d", i.NbTg)
}

func ParseInListPassiveTarget(data []byte) (packet.Payload, error) {
	return &InListPassiveTarget{
		ResponsePacket: ResponsePacket{
			ResponseCode:    data[0],
			SubResponseCode: data[1],
		},
		NbTg:       data[2],
		TargetData: data[3:],
	}, nil
}
