.PHONY: build
build:
	go build -o bin/terrafire ./cmd/terrafire 

.PHONY: test
test:
	go test ./terrafire