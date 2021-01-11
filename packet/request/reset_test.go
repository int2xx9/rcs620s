package request

import (
	"reflect"
	"testing"
)

func TestReset(t *testing.T) {
	cmd, _ := NewReset()
	c, ok := cmd.(*Reset)
	if !ok {
		t.Fail()
	}
	expected := Reset{
		RequestPacket: RequestPacket{
			CommandCode:    0xd4,
			SubCommandCode: 0x18,
		},
		Function: 0x01,
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}

func TestReset_ToByte(t *testing.T) {
	cmd, _ := NewReset()
	c, ok := cmd.(*Reset)
	if !ok {
		t.Fail()
	}
	expected := []byte{0xd4, 0x18, 0x01}
	actual := c.ToByte()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, actual)
	}
}
