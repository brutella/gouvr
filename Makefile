GO ?= go

all: test
	
test: 
	$(GO) test ./...