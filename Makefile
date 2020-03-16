.PHONY: build
build:
	go build -o bin/terrafire ./cmd/terrafire

.PHONY: test
test: export TERRAFIRE_REPORT_GITHUB_OWNER = mitene
test: export TERRAFIRE_REPORT_GITHUB_REPO  = terrafire
test: export TERRAFIRE_REPORT_GITHUB_ISSUE = 1
test:
	go test .

.PHONY: run
run:
	go run ./cmd/terrafire

.PHONY: fmt
fmt:
	go fmt ./cmd/terrafire
	go fmt .
