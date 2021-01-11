package request

import (
	"github.com/int2xx9/rcs620s/packet"
)

type RequestPacket struct {
	packet.Packet
	CommandCode    byte
	SubCommandCode byte
}
