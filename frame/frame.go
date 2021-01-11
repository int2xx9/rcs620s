package frame

import (
	"errors"

	"github.com/int2xx9/rcs620s/packet"
	"github.com/int2xx9/rcs620s/packet/response"
)

type Packet interface {
	Payload() packet.Payload
	ToByte() []byte
}

func CalcLCS(length byte) byte {
	return uint8(0x100 - uint16(length))
}

func CalcDCS(payload []byte) byte {
	var sum byte
	for _, b := range payload {
		sum += b
	}
	return uint8(0x100 - uint16(sum&0xff))
}

var (
	ErrInvalidFrame = errors.New("Invalid frame")
)

func Create(p packet.Payload) (Packet, error) {
	len := len(p.ToByte())
	if len <= 255 {
		return NewNormalFrame(p)
	} else if len > 255 && len <= 265 {
		panic("not implemented")
	}
	return nil, ErrWrongDataLength
}

func Parse(data []byte) (Packet, error) {
	if err := validateFrameCommon(data); err != nil {
		return nil, err
	}
	if data[3] == 0xff && data[4] == 0xff {
		return ParseExtendedFrame(data)
	} else if data[3] == 0x00 {
		return ParseAckFrame(data)
	} else {
		return ParseNormalFrame(data)
	}
}

func validateFrameCommon(data []byte) error {
	if data[0] != 0x00 {
		return ErrInvalidFrame
	}
	if data[1] != 0x00 {
		return ErrInvalidFrame
	}
	if data[2] != 0xff {
		return ErrInvalidFrame
	}
	if data[len(data)-1] != 0x00 {
		return ErrInvalidFrame
	}
	return nil
}

func ParseNormalFrame(data []byte) (Packet, error) {
	if CalcLCS(data[3]) != data[4] {
		return nil, ErrInvalidFrame
	}

	expectedSize := 7 + data[3]
	if int(expectedSize) != len(data) {
		return nil, ErrInvalidFrame
	}

	d, err := response.Parse(data[5 : len(data)-2])
	if err != nil {
		return nil, err
	}

	return NewNormalFrame(d)
}

func ParseExtendedFrame(data []byte) (Packet, error) {
	return nil, errors.New("Not implemented yet")
}

func ParseAckFrame(data []byte) (Packet, error) {
	if len(data) != 6 {
		return nil, ErrInvalidFrame
	}
	if data[3] != 0x00 {
		return nil, ErrInvalidFrame
	}
	if data[4] != 0xff {
		return nil, ErrInvalidFrame
	}
	return NewAckFrame()
}
