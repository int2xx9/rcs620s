package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/int2xx9/rcs620s"
	"github.com/int2xx9/rcs620s/packet/request"
	"github.com/int2xx9/rcs620s/packet/response"
)

type cardID uint64

func (c cardID) String() string {
	return fmt.Sprintf("%016x", uint64(c))
}

type cardInfo struct {
	ID        cardID
	Random    bool
	byteCount int
}

func usage() {
	fmt.Printf("usage: %s path_to_serial\n", os.Args[0])
	os.Exit(0)
}

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

func pollingA(r *rcs620s.RCS620S) (*cardInfo, error) {
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
	return &cardInfo{
		ID:     cardID(id),
		Random: (id >> 24) == 0x08,
	}, nil
}

func pollingB(r *rcs620s.RCS620S) (*cardInfo, error) {
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
	return &cardInfo{
		ID:     cardID(id),
		Random: true,
	}, nil
}

func pollingF(r *rcs620s.RCS620S) (*cardInfo, error) {
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
	return &cardInfo{
		ID:     cardID(id),
		Random: false,
	}, nil
}

func polling(r *rcs620s.RCS620S) (*cardInfo, error) {
	pollingFuncs := []func(*rcs620s.RCS620S) (*cardInfo, error){pollingA, pollingB, pollingF}
	var cardInfo *cardInfo
	for {
		for _, f := range pollingFuncs {
			var err error
			cardInfo, err = f(r)
			if err != nil {
				return nil, err
			}
			if cardInfo != nil {
				goto fin
			}
		}

		time.Sleep(1 * time.Second)
	}
fin:

	if err := r.CloseCommunication(); err != nil {
		return nil, err
	}

	return cardInfo, nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	r, err := rcs620s.New(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = r.Close(); err != nil {
			panic(err)
		}
	}()

	if err := r.InitializeAsInitiator(); err != nil {
		panic(err)
	}

	ra, err := polling(r)
	if err != nil {
		panic(err)
	}

	random := ""
	if ra.Random {
		random = " (random)"
	}
	fmt.Printf("%s%s\n", ra.ID, random)
}
