.PHONY: default copy test

default: uidreader

uidreader: $(shell find ../.. -name \*.go | grep -v _test.go$)
	GOOS=linux GOARCH=arm GOARM=6 go build -o uidreader ./main.go

test:
	go test -v ./...

copy: uidreader
	scp uidreader pi@raspberrypi.local:~/uidreader

run: uidreader
	ssh pi@raspberrypi.local /home/pi/uidreader

