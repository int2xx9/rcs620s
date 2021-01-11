package request

import "github.com/int2xx9/rcs620s/packet"

type CfgItemType byte

const (
	CfgItemRFField    CfgItemType = 0x01
	CfgItemMaxRetries CfgItemType = 0x05
	CfgItemWait       CfgItemType = 0x81
	CfgItemTimeout    CfgItemType = 0x82
)

type RFConfiguration struct {
	RequestPacket
	CfgItem           CfgItemType
	ConfigurationData []byte
}

func (g *RFConfiguration) ToByte() []byte {
	d := []byte{g.CommandCode, g.SubCommandCode, byte(g.CfgItem)}
	return append(d, g.ConfigurationData...)
}

func NewRFConfiguration(cfgItem CfgItemType, data []byte) (packet.Payload, error) {
	return &RFConfiguration{
		RequestPacket: RequestPacket{
			CommandCode:    0xd4,
			SubCommandCode: 0x32,
		},
		CfgItem:           cfgItem,
		ConfigurationData: data,
	}, nil
}

func CreateRFPowerCommand(status bool) (packet.Payload, error) {
	var data byte = 0
	if status {
		data |= 0x01
	} else {
		data &= 0xfe
	}
	return NewRFConfiguration(CfgItemRFField, []byte{data})
}
