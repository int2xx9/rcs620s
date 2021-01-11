package request

import (
	"reflect"
	"testing"
)

func TestGetGeneralStatus(t *testing.T) {
	cmd, _ := NewGetGeneralStatus()
	c, ok := cmd.(*GetGeneralStatus)
	if !ok {
		t.Fail()
	}

	expected := GetGeneralStatus{
		RequestPacket: RequestPacket{
			CommandCode:    0xd4,
			SubCommandCode: 0x04,
		},
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}
