package response

import (
	"github.com/int2xx9/rcs620s/packet"
)

type ResponsePacket struct {
	packet.Packet
	ResponseCode    byte
	SubResponseCode byte
}

type parserFunc func(data []byte) (packet.Payload, error)

var parserFuncs = map[byte]map[byte]parserFunc{
	0xd5: {
		0x03: ParseGetFirmwareVersion,
		0x05: ParseGetGeneralStatus,
		0x19: ParseReset,
		0x33: ParseRFConfiguration,
		0x4b: ParseInListPassiveTarget,
	},
}

func Parse(data []byte) (packet.Payload, error) {
	command := data[0]
	subCommand := data[1]

	if _, ok := parserFuncs[command]; !ok {
		return packet.ParseRaw(data)
	}
	if _, ok := parserFuncs[command][subCommand]; !ok {
		return packet.ParseRaw(data)
	}

	return parserFuncs[command][subCommand](data)
}
