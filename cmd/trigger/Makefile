.PHONY: default copy test

default: trigger

trigger: $(shell find ../.. -name \*.go | grep -v _test.go$)
	GOOS=linux GOARCH=arm GOARM=6 go build -o trigger ./main.go ./cmdline.go ./trigger.go ./reader.go

test:
	go test -v ./...

copy: trigger
	scp trigger pi@raspberrypi.local:~/trigger

run: trigger
	ssh pi@raspberrypi.local /home/pi/trigger

