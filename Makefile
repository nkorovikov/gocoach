.PHONY: build
build:
	go build -o bin/gocoach cmd/gocoach/main.go

.PHONY: run
run:
	go run cmd/gocoach/main.go