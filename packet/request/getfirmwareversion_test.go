package request

import (
	"reflect"
	"testing"
)

func TestGetFirmwareVersion(t *testing.T) {
	cmd, _ := NewGetFirmwareVersion()
	c, ok := cmd.(*GetFirmwareVersion)
	if !ok {
		t.Fail()
	}
	if c.CommandCode != 0xd4 {
		t.Fail()
	}
	if c.SubCommandCode != 0x02 {
		t.Fail()
	}
}

func TestGetFirmwareVersion_ToByte(t *testing.T) {
	cmd, _ := NewGetFirmwareVersion()
	c, ok := cmd.(*GetFirmwareVersion)
	if !ok {
		t.Fail()
	}

	expected := []byte{0xd4, 0x02}
	actual := c.ToByte()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, actual)
	}
}
