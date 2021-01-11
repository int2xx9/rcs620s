package rcs620s

import (
	"errors"
	"time"

	"github.com/int2xx9/rcs620s/frame"
	"github.com/int2xx9/rcs620s/packet"
	"github.com/int2xx9/rcs620s/packet/request"
	"github.com/int2xx9/rcs620s/packet/response"
)

var (
	ErrInvalidResponse = errors.New("invalid response")
)

type RCS620S struct {
	Datalink       *Datalink
	lastAntennaOff time.Time
}

func New(deviceFile string) (*RCS620S, error) {
	datalink, err := NewDatalink(deviceFile)
	if err != nil {
		return nil, err
	}
	return &RCS620S{
		Datalink: datalink,
	}, nil
}

func (r *RCS620S) Close() error {
	return r.Datalink.Close()
}

// InitializeAsInitiator initializes RCS620 to run as initiator mode.
func (r *RCS620S) InitializeAsInitiator() error {
	if err := r.Reset(); err != nil {
		return err
	}

	req, _ := request.NewRFConfiguration(request.CfgItemWait, []byte{0xff})
	if _, err := r.Datalink.Request(req); err != nil {
		panic(err)
	}

	req, _ = request.NewRFConfiguration(request.CfgItemMaxRetries, []byte{0xff, 0x00, 0x00})
	if _, err := r.Datalink.Request(req); err != nil {
		panic(err)
	}

	if err := r.SetAntenna(false); err != nil {
		return err
	}

	return nil
}

// Reset resets RCS620S.
// This method improves noise immunity at the same time.
func (r *RCS620S) Reset() error {
	cmd, _ := request.NewReset()
	if _, err := r.Datalink.Request(cmd); err != nil {
		return err
	}

	if err := r.Datalink.WriteACK(); err != nil {
		return err
	}

	time.Sleep(20 * time.Millisecond)

	if err := r.ImproveNoiseImmunity(); err != nil {
		return err
	}

	return nil
}

// ApplyInListPassiveTargetWorkaround fixes no response issue.
// Read section 9 in below document for more details.
// https://www.sony.co.jp/Products/felica/business/tech-support/data/fp_rcs620s_command_2.11.pdf
func (r *RCS620S) ApplyInListPassiveTargetWorkaround() (bool, error) {
	writeData, err := packet.NewRaw([]byte{0xd4, 0x06, 0x63, 0x37, 0x63, 0x31})
	if err != nil {
		return false, err
	}
	if err := r.Datalink.WritePacket(writeData); err != nil {
		return false, err
	}

	f, err := r.Datalink.ReadFrame()
	if err != nil || !frame.IsAck(f) {
		return false, err
	}

	f, err = r.Datalink.ReadFrame()
	if err != nil {
		return false, err
	}

	resp1 := f.Payload().ToByte()
	if (resp1[2] & 0x08) == 0 {
		return false, nil
	}

	// need workaround

	z := resp1[3] & 0xf0
	dataFrame, _ := packet.NewRaw([]byte{0xd4, 0x08, 0x63, 0x31, z | 0x08, 0x63, 0x31, z | 0x00})
	if err := r.Datalink.WritePacket(dataFrame); err != nil {
		return false, err
	}

	f, err = r.Datalink.ReadFrame()
	if err != nil || !frame.IsAck(f) {
		return false, err
	}

	_, err = r.Datalink.ReadFrame()
	if err != nil {
		return false, err
	}

	return true, nil
}

// ImproveNoiseImmunity improves noise immunity.
// Read section 9 in below document for more details.
// https://www.sony.co.jp/Products/felica/business/tech-support/data/fp_rcs620s_command_2.11.pdf
func (r *RCS620S) ImproveNoiseImmunity() error {
	writeData, err := packet.NewRaw([]byte{0xd4, 0x08, 0x63, 0x0a, 0x40})
	if err != nil {
		return err
	}
	if err := r.Datalink.WritePacket(writeData); err != nil {
		return err
	}

	f, err := r.Datalink.ReadFrame()
	if err != nil || !frame.IsAck(f) {
		return err
	}

	f, err = r.Datalink.ReadFrame()
	if err != nil {
		return err
	}

	return nil
}

// FirmwareVersion gets firmware version of RCS620S
func (r *RCS620S) FirmwareVersion() (*response.GetFirmwareVersion, error) {
	f, _ := request.NewGetFirmwareVersion()
	res, err := r.Datalink.Request(f)
	if err != nil {
		return nil, err
	}
	version, ok := res.(*response.GetFirmwareVersion)
	if !ok {
		return nil, ErrInvalidResponse
	}
	return version, nil
}

// GeneralStatus gets general status of RCS620S
func (r *RCS620S) GeneralStatus() (*response.GetGeneralStatus, error) {
	f, _ := request.NewGetGeneralStatus()
	res, err := r.Datalink.Request(f)
	if err != nil {
		return nil, err
	}
	status, ok := res.(*response.GetGeneralStatus)
	if !ok {
		return nil, ErrInvalidResponse
	}
	return status, nil
}

// SetAntenna enables or disables antenna
func (r *RCS620S) SetAntenna(enable bool) error {
	if enable {
		// Wait 100ms since last antenna is turned off
		// This is required for clearing captured target (see command reference section 8.2.5)
		waitTime := r.lastAntennaOff.Add(100 * time.Millisecond).Sub(time.Now())
		if waitTime > 0 {
			time.Sleep(waitTime)
		}
	}

	req, _ := request.CreateRFPowerCommand(enable)
	_, err := r.Datalink.Request(req)
	if err != nil {
		return err
	}

	if !enable {
		r.lastAntennaOff = time.Now()
	}

	return nil
}

// InListPassiveTarget detects a card
func (r *RCS620S) InListPassiveTarget(brty request.BRTYType, initiatorData []byte) (*response.InListPassiveTarget, error) {
	r.SetAntenna(true)

	if _, err := r.ApplyInListPassiveTargetWorkaround(); err != nil {
		panic(err)
	}

	req, _ := request.NewInListPassiveTarget(brty, initiatorData)
	res, err := r.Datalink.Request(req)
	if err != nil {
		return nil, err
	}

	target, ok := res.(*response.InListPassiveTarget)
	if !ok {
		return nil, nil
	}

	return target, nil
}

func (r *RCS620S) CloseCommunication() error {
	if err := r.Datalink.WriteACK(); err != nil {
		return err
	}

	if err := r.InitializeAsInitiator(); err != nil {
		return err
	}

	return nil
}
