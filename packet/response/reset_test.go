package response

import (
	"reflect"
	"testing"
)

func TestReset(t *testing.T) {
	cmd, _ := ParseReset([]byte{
		0xd5, 0x19,
	})
	c, ok := cmd.(*Reset)
	if !ok {
		t.Fail()
	}

	expected := Reset{
		ResponsePacket: ResponsePacket{
			ResponseCode:    0xd5,
			SubResponseCode: 0x19,
		},
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}
