package rcs620s

import (
	"io"
	"time"

	"github.com/goburrow/serial"
	"github.com/int2xx9/rcs620s/frame"
	"github.com/int2xx9/rcs620s/packet"
)

type Datalink struct {
	port serial.Port
}

func NewDatalink(deviceFile string) (*Datalink, error) {
	config := serial.Config{
		Address:  deviceFile,
		BaudRate: 115200,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  2 * time.Second,
	}
	port, err := serial.Open(&config)
	if err != nil {
		return nil, err
	}
	return &Datalink{
		port: port,
	}, nil
}

func (d *Datalink) Close() error {
	return d.port.Close()
}

func (d *Datalink) WriteFrame(f frame.Packet) error {
	if _, err := d.WriteByte(f.ToByte()); err != nil {
		return err
	}
	return nil
}

func (d *Datalink) WriteByte(f []byte) (int, error) {
	return d.port.Write(f)
}

func (d *Datalink) WritePacket(f packet.Payload) error {
	p, err := frame.Create(f)
	if err != nil {
		return err
	}
	return d.WriteFrame(p)
}

func (d *Datalink) WriteACK() error {
	ackFrame, _ := frame.NewAckFrame()
	return d.WriteFrame(ackFrame)
}

func (d *Datalink) read(len int) ([]byte, error) {
	b := make([]byte, len)
	_, err := io.ReadFull(d.port, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (d *Datalink) ReadFrame() (frame.Packet, error) {
	// need first 4bytes or 7bytes to read whole packet
	head, err := d.read(4)
	if err != nil {
		return nil, err
	}

	var tail []byte
	if head[3] == 0xff {
		// extended frame
		panic("not implemented")
	} else if head[3] == 0x00 {
		// ack frame
		tail, err = d.read(2)
	} else {
		// normal frame
		tail, err = d.read(int(head[3]) + 3)
	}
	if err != nil {
		return nil, err
	}

	data := append(head, tail...)
	fr, err := frame.Parse(data)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func (d *Datalink) Request(packet packet.Payload) (packet.Payload, error) {
	writeFrame, err := frame.Create(packet)
	if err != nil {
		return nil, err
	}

	if err := d.WriteFrame(writeFrame); err != nil {
		return nil, err
	}

	f, err := d.ReadFrame()
	if err != nil || !frame.IsAck(f) {
		return nil, err
	}

	f, err = d.ReadFrame()
	if err != nil {
		return nil, err
	}

	return f.Payload(), nil
}
