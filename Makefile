build:
	mkdir -p bin
	go build -o bin/main cmd/main.go

build-labels:
	mkdir -p bin
	go build -o bin/ginkgo-labels tools/ginkgo-labels/main.go

test:
	go test -v ./...

all: build build-labels