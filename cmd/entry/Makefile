.PHONY: default copy test

default: entry

entry: $(shell find ../.. -name \*.go | grep -v _test.go$)
	GOOS=linux GOARCH=arm GOARM=6 go build -o entry ./main.go

test:
	go test -v ./...

copy: rcs620s
	scp rcs620s pi@raspberrypi.local:~/rcs620s

run: rcs620s
	ssh pi@raspberrypi.local /home/pi/rcs620s

