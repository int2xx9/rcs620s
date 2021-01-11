package main

import (
	"errors"
	"time"

	"github.com/int2xx9/rcs620s"
	"github.com/int2xx9/rcs620s/packet/request"
	"github.com/int2xx9/rcs620s/packet/response"
)

func pollingCommon(r *rcs620s.RCS620S, brty request.BRTYType, data []byte) (*response.InListPassiveTarget, error) {
	target, err := r.InListPassiveTarget(brty, data)
	if err != nil {
		return nil, err
	}
	if target == nil {
		return nil, errors.New("cast error")
	}

	return target, nil
}

func pollingA(r *rcs620s.RCS620S) (*TriggerResult, error) {
	ret, err := pollingCommon(r, request.BRTYType(byte(0)), []byte{})
	if err != nil {
		return nil, err
	}

	if ret.NbTg != 1 {
		return nil, nil
	}

	var id uint64
	id |= uint64(ret.TargetData[5]) << 24
	id |= uint64(ret.TargetData[6]) << 16
	id |= uint64(ret.TargetData[7]) << 8
	id |= uint64(ret.TargetData[8]) << 0

	// TypeA card ID has 3 variants 4bytes, 7bytes and 10bytes
	// however, currently support only 4bytes ID.
	// Use first 4bytes if ID is longer than 4bytes.
	return &TriggerResult{
		ID: CardID{
			ID:     id,
			Length: 4,
			Random: (id >> 24) == 0x08,
		},
		Time: time.Now(),
		Type: NFCTypeA,
	}, nil
}

func pollingB(r *rcs620s.RCS620S) (*TriggerResult, error) {
	ret, err := pollingCommon(r, request.BRTYType(byte(3)), []byte{0x00})
	if err != nil {
		return nil, err
	}

	if ret.NbTg != 1 {
		return nil, nil
	}

	var id uint64
	id |= uint64(ret.TargetData[2]) << 24
	id |= uint64(ret.TargetData[3]) << 16
	id |= uint64(ret.TargetData[4]) << 8
	id |= uint64(ret.TargetData[5]) << 0
	return &TriggerResult{
		ID: CardID{
			ID:     id,
			Length: 4,
			Random: true,
		},
		Time: time.Now(),
		Type: NFCTypeB,
	}, nil
}

func pollingF(r *rcs620s.RCS620S) (*TriggerResult, error) {
	ret, err := pollingCommon(r, request.BRTY212, []byte{0x00, 0xff, 0xff, 0x00, 0x00})
	if err != nil {
		return nil, err
	}

	if ret.NbTg != 1 {
		return nil, nil
	}

	var id uint64
	id |= uint64(ret.TargetData[3]) << 56
	id |= uint64(ret.TargetData[4]) << 48
	id |= uint64(ret.TargetData[5]) << 40
	id |= uint64(ret.TargetData[6]) << 32
	id |= uint64(ret.TargetData[7]) << 24
	id |= uint64(ret.TargetData[8]) << 16
	id |= uint64(ret.TargetData[9]) << 8
	id |= uint64(ret.TargetData[10]) << 0
	return &TriggerResult{
		ID: CardID{
			ID:     id,
			Length: 8,
			Random: false,
		},
		Time: time.Now(),
		Type: NFCTypeF,
	}, nil
}

func polling(r *rcs620s.RCS620S, opts *AppOptions) (*TriggerResult, error) {
	pollingFuncs := map[NFCType]func(*rcs620s.RCS620S) (*TriggerResult, error){
		NFCTypeA: pollingA,
		NFCTypeB: pollingB,
		NFCTypeF: pollingF,
	}

	var result *TriggerResult
	for {
		for t, f := range pollingFuncs {
			if opts.NFCType&t == 0 {
				continue
			}

			var err error
			result, err = f(r)
			if err != nil {
				return nil, err
			}
			if result != nil {
				goto fin
			}
		}

		time.Sleep(time.Duration(opts.PollingInterval) * time.Millisecond)
	}
fin:

	if err := r.CloseCommunication(); err != nil {
		return nil, err
	}

	return result, nil
}
