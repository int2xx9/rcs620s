package response

import (
	"fmt"

	"github.com/int2xx9/rcs620s/packet"
)

type GetGeneralStatus struct {
	ResponsePacket
	Err   byte
	NA1   byte
	NbTg  byte
	Tg1   byte
	BrRx1 byte
	BrTx1 byte
	Type1 byte
	NA2   byte
	Mode  byte
}

func (g *GetGeneralStatus) ToByte() []byte {
	return []byte{g.ResponseCode, g.SubResponseCode, g.Err, g.NA1, g.NbTg, g.Tg1, g.BrRx1, g.BrTx1, g.Type1, g.NA2, g.Mode}
}

func (g *GetGeneralStatus) String() string {
	if g.NbTg == 0x00 {
		return fmt.Sprintf("response.GetGeneralStatus: Err:%02x, NbTg:%d, Mode:%02x", g.Err, g.NbTg, g.Mode)
	} else {
		return fmt.Sprintf("response.GetGeneralStatus: Err:%02x, NbTg:%d, Type1:%02x, Mode:%02x", g.Err, g.NbTg, g.Type1, g.Mode)
	}
}

func ParseGetGeneralStatus(data []byte) (packet.Payload, error) {
	var status *GetGeneralStatus
	if data[4] != 0x00 {
		status = &GetGeneralStatus{
			ResponsePacket: ResponsePacket{
				ResponseCode:    data[0],
				SubResponseCode: data[1],
			},
			Err:   data[2],
			NA1:   data[3],
			NbTg:  data[4],
			Tg1:   data[5],
			BrRx1: data[6],
			BrTx1: data[7],
			Type1: data[8],
			NA2:   data[9],
			Mode:  data[10],
		}
	} else {
		status = &GetGeneralStatus{
			ResponsePacket: ResponsePacket{
				ResponseCode:    data[0],
				SubResponseCode: data[1],
			},
			Err:  data[2],
			NA1:  data[3],
			NbTg: data[4],
			NA2:  data[5],
			Mode: data[6],
		}
	}

	return status, nil
}
