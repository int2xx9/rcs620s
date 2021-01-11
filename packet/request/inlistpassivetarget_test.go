package request

import (
	"reflect"
	"testing"
)

func TestInListPassiveTarget(t *testing.T) {
	cmd, _ := NewInListPassiveTarget(BRTY212, []byte{0x00, 0x00, 0x00, 0x00, 0x00})
	c, ok := cmd.(*InListPassiveTarget)
	if !ok {
		t.Fail()
	}
	if c.MaxTg != 1 {
		t.Fail()
	}
	if c.BRTY != BRTY212 {
		t.Fail()
	}

	expected := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
	actual := c.InitiatorData
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, actual)
	}
}

func TestInListPassiveTarget_ToByte(t *testing.T) {
	cmd, _ := NewInListPassiveTarget(BRTY212, []byte{0x00, 0x00, 0x00, 0x00, 0x00})
	c, ok := cmd.(*InListPassiveTarget)
	if !ok {
		t.Fail()
	}

	expected := []byte{0xd4, 0x4a, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}
	actual := c.ToByte()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, actual)
	}
}
