package response

import (
	"reflect"
	"testing"
)

func TestParseRFConfiguration(t *testing.T) {
	cmd, _ := ParseRFConfiguration([]byte{
		0xd5, 0x33,
	})
	c, ok := cmd.(*RFConfiguration)
	if !ok {
		t.Fail()
	}

	expected := RFConfiguration{
		ResponsePacket: ResponsePacket{
			ResponseCode:    0xd5,
			SubResponseCode: 0x33,
		},
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}
