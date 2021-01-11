package response

import (
	"reflect"
	"testing"
)

func TestParseInListPassiveTarget(t *testing.T) {
	cmd, _ := ParseInListPassiveTarget([]byte{
		0xd5, 0x4b, 0x01, 0x00,
	})
	c, ok := cmd.(*InListPassiveTarget)
	if !ok {
		t.Fail()
	}
	if c.ResponseCode != 0xd5 {
		t.Fail()
	}
	if c.SubResponseCode != 0x4b {
		t.Fail()
	}
	if c.NbTg != 0x01 {
		t.Fail()
	}
	expected := []byte{0x00}
	actual := c.TargetData
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, actual)
	}
}
