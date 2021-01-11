package response

import (
	"fmt"

	"github.com/int2xx9/rcs620s/packet"
)

type Reset struct {
	ResponsePacket
}

func (r *Reset) ToByte() []byte {
	return []byte{r.ResponseCode, r.SubResponseCode}
}

func (r *Reset) String() string {
	return fmt.Sprintf("response.Reset")
}

func ParseReset(data []byte) (packet.Payload, error) {
	return &Reset{
		ResponsePacket: ResponsePacket{
			ResponseCode:    data[0],
			SubResponseCode: data[1],
		},
	}, nil
}
