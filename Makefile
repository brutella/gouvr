GO ?= go

all: test
	
test: 
	$(GO) test ./...

bbb:
	GOOS=linux GOARCH=arm GOARM=7 $(GO) build -o uvr1611 _example/uvr1611_linux.go

rpi:
	GOOS=linux GOARCH=arm GOARM=6 $(GO) build -o uvr1611 _example/uvr1611_linux.go