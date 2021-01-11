package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/int2xx9/rcs620s"
)

type NFCType byte

func (t NFCType) String() string {
	list := []string{}
	if t&NFCTypeA != 0 {
		list = append(list, "TypeA")
	}
	if t&NFCTypeB != 0 {
		list = append(list, "TypeB")
	}
	if t&NFCTypeF != 0 {
		list = append(list, "TypeF")
	}
	if len(list) > 0 {
		return strings.Join(list, ",")
	} else {
		return "(none)"
	}
}

const (
	NFCTypeA NFCType = 1 << iota
	NFCTypeB
	NFCTypeF
)

type CardID struct {
	ID     uint64
	Length int
	Random bool
}

func (c *CardID) String() string {
	idstr := fmt.Sprintf("%016x", c.ID)
	idstr = idstr[16-c.Length*2 : len(idstr)]

	idstr2 := ""
	for len(idstr) > 0 {
		idstr2 += idstr[0:2]
		idstr = idstr[2:]
		if len(idstr) > 0 {
			idstr2 += ":"
		}
	}

	return idstr2
}

type TriggerMessage struct {
	Time   int64
	Type   string
	ID     string
	Random bool
}

func main() {
	opts := parseArguments()
	if len(opts.Triggers) == 0 {
		fmt.Fprintf(os.Stderr, "Warning: no triggers\n")
	}

	r, err := rcs620s.New(opts.Serial)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = r.Close(); err != nil {
			panic(err)
		}
	}()

	if err := r.Reset(); err != nil {
		panic(err)
	}

	if err := r.ImproveNoiseImmunity(); err != nil {
		panic(err)
	}

	for {
		result, err := polling(r, opts)
		if err != nil {
			panic(err)
		}

		for _, value := range opts.Triggers {
			trigger := value
			go func() {
				err := trigger.Do(&TriggerMessage{
					ID:     result.ID.String(),
					Time:   result.Time.Unix(),
					Type:   result.Type.String(),
					Random: result.ID.Random,
				})
				if err != nil {
					fmt.Fprintf(os.Stderr, "trigger error: %s\n", err.Error())
				}
			}()
		}

		random := ""
		if result.ID.Random {
			random = " (random)"
		}
		fmt.Fprintf(os.Stderr, "%s%s\n", result.ID.String(), random)

		if opts.Onetime {
			break
		}

		time.Sleep(opts.DetectInterval * time.Millisecond)
	}
}
