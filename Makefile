.PHONY: build
build:
	go build -o bin/terrafire ./cmd/terrafire 

.PHONY: test
test:
	go test .

.PHONY: run
run:
	go run ./cmd/terrafire

.PHONY: fmt
fmt:
	go fmt ./cmd/terrafire
	go fmt .
