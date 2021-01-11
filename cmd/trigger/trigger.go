package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

type TriggerResult struct {
	Time time.Time
	Type NFCType
	ID   CardID
}

type Trigger interface {
	Do(msg *TriggerMessage) error
}

type HTTPTrigger struct {
	Method  string
	Address string
}

var (
	ErrHTTPTrigger = xerrors.New("HTTPTrigger error")
)

func (h *HTTPTrigger) Do(msg *TriggerMessage) error {
	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	client := http.Client{}
	var req *http.Request
	switch strings.ToLower(h.Method) {
	case "get":
		url, err := url.Parse(h.Address)
		if err != nil {
			return err
		}
		q := url.Query()
		q.Set("readerData", string(json))
		url.RawQuery = q.Encode()

		req, err = http.NewRequest("GET", url.String(), bytes.NewReader([]byte{}))
	case "post":
		req, err = http.NewRequest("POST", h.Address, bytes.NewReader(json))
	default:
		return ErrHTTPTrigger
	}
	if err != nil {
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return ErrHTTPTrigger
	}

	return nil
}

type ExecuteTrigger struct {
	Path string
}

func (e *ExecuteTrigger) Do(msg *TriggerMessage) error {
	cmd := exec.Command(e.Path)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	stdin.Write(json)
	if err := stdin.Close(); err != nil {
		return err
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

type StdoutTrigger struct {
}

func (s *StdoutTrigger) Do(msg *TriggerMessage) error {
	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println(string(json))
	return nil
}
