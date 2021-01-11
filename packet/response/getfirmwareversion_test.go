package response

import (
	"reflect"
	"testing"
)

func TestParseGetFirmwareVersion(t *testing.T) {
	cmd, _ := ParseGetFirmwareVersion([]byte{
		0xd5, 0x03, 0x33, 0x01, 0x30, 0x07,
	})
	c, ok := cmd.(*GetFirmwareVersion)
	if !ok {
		t.Fail()
	}

	expected := GetFirmwareVersion{
		ResponsePacket: ResponsePacket{
			ResponseCode:    0xd5,
			SubResponseCode: 0x03,
		},
		ICType: 0x33,
		Ver:    0x130,
		NA:     0x07,
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}
