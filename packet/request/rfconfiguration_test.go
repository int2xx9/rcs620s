package request

import (
	"reflect"
	"testing"
)

func TestRFConfiguration(t *testing.T) {
	cmd, _ := NewRFConfiguration(CfgItemRFField, []byte{0x01})
	c, ok := cmd.(*RFConfiguration)
	if !ok {
		t.Fail()
	}

	if c.CfgItem != CfgItemRFField {
		t.Fail()
	}

	if !reflect.DeepEqual(c.ConfigurationData, []byte{0x01}) {
		t.Fail()
	}
}

func TestCreateRFPowerCommand_Enable(t *testing.T) {
	cmd, _ := CreateRFPowerCommand(true)
	c, ok := cmd.(*RFConfiguration)
	if !ok {
		t.Fail()
	}

	if c.CfgItem != CfgItemRFField {
		t.Fail()
	}

	if !reflect.DeepEqual(c.ConfigurationData, []byte{0x01}) {
		t.Fail()
	}
}

func TestCreateRFPowerCommand_Disable(t *testing.T) {
	cmd, _ := CreateRFPowerCommand(false)
	c, ok := cmd.(*RFConfiguration)
	if !ok {
		t.Fail()
	}

	if c.CfgItem != CfgItemRFField {
		t.Fail()
	}

	if !reflect.DeepEqual(c.ConfigurationData, []byte{0x00}) {
		t.Fail()
	}
}
