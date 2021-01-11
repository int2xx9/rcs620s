package response

import (
	"fmt"

	"github.com/int2xx9/rcs620s/packet"
)

type RFConfiguration struct {
	ResponsePacket
}

func (r *RFConfiguration) ToByte() []byte {
	return []byte{r.ResponseCode, r.SubResponseCode}
}

func (r *RFConfiguration) String() string {
	return fmt.Sprintf("response.RFConfiguration")
}

func ParseRFConfiguration(data []byte) (packet.Payload, error) {
	return &RFConfiguration{
		ResponsePacket: ResponsePacket{
			ResponseCode:    data[0],
			SubResponseCode: data[1],
		},
	}, nil
}
