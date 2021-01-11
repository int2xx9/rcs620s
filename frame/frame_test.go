package frame

import (
	"reflect"
	"testing"
)

func TestParseNormalFrame(t *testing.T) {
	frame, err := ParseNormalFrame([]byte{
		0x00,       // preamble
		0x00, 0xff, // start of packet
		0x06,                               // len
		0xfa,                               // lcs
		0xd5, 0x03, 0x33, 0x01, 0x30, 0x07, // data
		0xbd, // dcs
		0x00, // postamble
	})
	if err != nil {
		t.Fail()
	}

	actual := frame.Payload().ToByte()
	expected := []byte{0xd5, 0x03, 0x33, 0x01, 0x30, 0x07}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, actual)
	}
}

func TestParseAckFrame(t *testing.T) {
	frame, err := ParseAckFrame([]byte{
		0x00,       // preamble
		0x00, 0xff, //start of packet
		0x00, // len
		0xff, // lcs
		0x00, // postamble
	})
	if err != nil {
		t.Fail()
	}

	actual := frame.Payload().ToByte()
	expected := []byte{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:nil, actual:%+v", frame.Payload().ToByte())
	}
}
