package packet

import (
	"reflect"
	"testing"
)

func TestNewRaw(t *testing.T) {
	cmd, _ := NewRaw([]byte{
		0x01, 0x02, 0x03,
	})
	c, ok := cmd.(*Raw)
	if !ok {
		t.Fail()
	}

	expected := Raw{
		data: []byte{0x01, 0x02, 0x03},
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}

func TestParseRaw(t *testing.T) {
	cmd, _ := ParseRaw([]byte{
		0x01, 0x02, 0x03,
	})
	c, ok := cmd.(*Raw)
	if !ok {
		t.Fail()
	}

	expected := Raw{
		data: []byte{0x01, 0x02, 0x03},
	}
	if !reflect.DeepEqual(*c, expected) {
		t.Errorf("expected:%+v, actual:%+v", expected, *c)
	}
}
