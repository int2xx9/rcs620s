package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

var (
	ErrInvalidOption         = xerrors.New("invalid option")
	ErrInvalidHTTPTrigger    = xerrors.Errorf("invalid httptrigger %v", ErrInvalidOption)
	ErrInvalidExecuteTrigger = xerrors.Errorf("invalid executetrigger %v", ErrInvalidOption)
)

type AppOptions struct {
	Serial          string
	Onetime         bool
	Verbose         bool
	NFCType         NFCType
	PollingInterval time.Duration
	DetectInterval  time.Duration
	Stdout          bool
	Triggers        []Trigger
}

func parseHTTPTrigger(arg string) (Trigger, error) {
	methodStr := strings.ToLower(arg[0:strings.Index(arg, ":")])
	var method string
	var url string
	switch methodStr {
	case "get":
		method = "GET"
		url = arg[strings.Index(arg, ":")+1:]
	case "post":
		method = "POST"
		url = arg[strings.Index(arg, ":")+1:]
	default:
		method = "POST"
		url = arg
	}
	if !strings.Contains(arg, "://") {
		return nil, ErrInvalidHTTPTrigger
	}
	return &HTTPTrigger{
		Method:  method,
		Address: url,
	}, nil
}

func parseExecuteTrigger(arg string) (Trigger, error) {
	_, err := os.Stat(arg)
	if err != nil {
		return nil, ErrInvalidExecuteTrigger
	}

	return &ExecuteTrigger{
		Path: arg,
	}, nil
}

func parseTriggerArgument(arg string) (Trigger, error) {
	var trigger Trigger
	var err error
	if strings.Contains(arg, "://") {
		trigger, err = parseHTTPTrigger(arg)
	} else {
		trigger, err = parseExecuteTrigger(arg)
	}
	return trigger, err
}

func parseTriggerArguments(args []string) ([]Trigger, error) {
	triggers := []Trigger{}
	for _, arg := range args {
		trigger, err := parseTriggerArgument(arg)
		if err != nil {
			return nil, err
		}
		triggers = append(triggers, trigger)
	}
	return triggers, nil
}

func validateArguments(opts *AppOptions) error {
	if opts.PollingInterval < 1 {
		return ErrInvalidOption
	}
	if opts.DetectInterval < 1 {
		return ErrInvalidOption
	}
	return nil
}

func parseArguments() *AppOptions {
	var (
		serial          = flag.String("serial", "/dev/serial0", "Path to serial device of RC-S620S")
		onetime         = flag.Bool("onetime", false, "Enable onetime mode")
		verbose         = flag.Bool("verbose", false, "Enable verbose output")
		nfcType         = flag.String("type", "ABF", "Detect NFC type (A, B, F)")
		pollingInterval = flag.Int("pollingInterval", 1000, "Interval for polling in ms")
		detectInterval  = flag.Int("detectInterval", 1000, "Interval for reading next card in ms")
		stdout          = flag.Bool("stdout", false, "Write json to stdout")
	)

	flag.Parse()

	opts := &AppOptions{
		Serial:          "/dev/serial0",
		Onetime:         false,
		Verbose:         false,
		NFCType:         NFCTypeA | NFCTypeB | NFCTypeF,
		PollingInterval: time.Duration(1000),
		DetectInterval:  time.Duration(1000),
		Stdout:          false,
		Triggers:        []Trigger{},
	}

	if serial != nil {
		opts.Serial = *serial
	}
	if onetime != nil {
		opts.Onetime = *onetime
	}
	if verbose != nil {
		opts.Verbose = *verbose
	}
	if nfcType != nil {
		t := NFCType(0)
		for _, c := range *nfcType {
			switch c {
			case 'a', 'A':
				t |= NFCTypeA
			case 'b', 'B':
				t |= NFCTypeB
			case 'f', 'F':
				t |= NFCTypeF
			}
		}
		opts.NFCType = t
	}
	if pollingInterval != nil {
		opts.PollingInterval = time.Duration(*pollingInterval)
	}
	if detectInterval != nil {
		opts.DetectInterval = time.Duration(*detectInterval)
	}
	if stdout != nil {
		opts.Stdout = *stdout
	}
	triggerArgs, err := parseTriggerArguments(flag.Args())
	if err != nil {
		panic(err)
	}
	opts.Triggers = triggerArgs

	if opts.Stdout {
		opts.Triggers = append(opts.Triggers, &StdoutTrigger{})
	}

	if err := validateArguments(opts); err != nil {
		panic(err)
	}

	return opts
}
