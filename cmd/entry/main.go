package main

import (
	"fmt"
	"os"

	"github.com/int2xx9/rcs620s"
	"github.com/int2xx9/rcs620s/packet/request"
)

func usage() {
	fmt.Printf("usage: %s path_to_serial\n", os.Args[0])
	os.Exit(0)
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

	_, err = r.FirmwareVersion()
	if err != nil {
		panic(err)
	}

	if err = r.InitializeAsInitiator(); err != nil {
		panic(err)
	}

	target, err := r.InListPassiveTarget(request.BRTY212, []byte{0x00, 0xff, 0xff, 0x00, 0x00})
	if err != nil {
		panic(err)
	}
	if target == nil {
		panic("nocard")
	}

	fmt.Printf("targetData: %+v\n", target.TargetData)

	status, err := r.GeneralStatus()
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %s\n", status)

	if err := r.CloseCommunication(); err != nil {
		panic(err)
	}
}
