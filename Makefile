
VERSION = $(shell printf "%s.%s" \
		$$(git rev-list --count HEAD) \
		$$(git rev-parse --short HEAD))


run:
	@echo :: start the crawler
	go get -v -d
	go run main.go

build:  $(OUTPUT)
	CGO_ENABLED=0 GOOS=linux go build -o bin/app \
		-ldflags "-X main.version=$(VERSION)" \
		-gcflags "-trimpath $(GOPATH)/src"

all: build run


$(OUTPUT):
	mkdir -p $(OUTPUT)