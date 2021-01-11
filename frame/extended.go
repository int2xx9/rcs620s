package frame

import "errors"

type Extended struct {
}

func NewExtendedFrame(data []byte) (*Extended, error) {
	if len(data) < 256 || len(data) > 265 {
		return nil, errors.New("aa")
	}
	return nil, nil
}
