package packet

import "fmt"

type Raw struct {
	data []byte
}

func (r *Raw) ToByte() []byte {
	d := make([]byte, len(r.data))
	copy(d, r.data)
	return d
}

func (r *Raw) String() string {
	return fmt.Sprintf("packet.Raw")
}

func NewRaw(data []byte) (Payload, error) {
	d := make([]byte, len(data))
	copy(d, data)
	return &Raw{
		data: d,
	}, nil
}

func ParseRaw(data []byte) (Payload, error) {
	d := make([]byte, len(data))
	copy(d, data)
	return &Raw{
		data: d,
	}, nil
}
