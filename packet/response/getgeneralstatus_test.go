package response

import (
	"reflect"
	"testing"
)

func TestParseGetGeneralStatus(t *testing.T) {
	cmd, _ := ParseGetGeneralStatus([]byte{
		0xd5, 0x05, 0xff, 0x02, 0x01, 0x03, 0x04, 0x05, 0x10, 0x00, 0x06,
	})
	c, ok := cmd.(*GetGeneralStatus)
	if !ok {
		t.Fail()
	}

	expected := GetGeneralStatus{
		ResponsePacket: ResponsePacket{
			ResponseCode:    0xd5,
			SubResponseCode: 0x05,
		},
		Err:   0xff,
		NA1:   0x02,
		NbTg:  0x01,
		Tg1:   0x03,
		BrRx1: 0x04,
		BrTx1: 0x05,
		Type1: 0x10,
		NA2:   0x00,
		Mode:  0x06,
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}

func TestParseGetGeneralStatus_NoNbTg(t *testing.T) {
	cmd, _ := ParseGetGeneralStatus([]byte{
		0xd5, 0x05, 0xff, 0x02, 0x00, 0x00, 0x06,
	})
	c, ok := cmd.(*GetGeneralStatus)
	if !ok {
		t.Fail()
	}

	expected := GetGeneralStatus{
		ResponsePacket: ResponsePacket{
			ResponseCode:    0xd5,
			SubResponseCode: 0x05,
		},
		Err:  0xff,
		NA1:  0x02,
		NbTg: 0x00,
		NA2:  0x00,
		Mode: 0x06,
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}
