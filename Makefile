ALL := ./cmd/terrafire ./core ./database ./runner ./server ./utils

.PHONY: build
build:
	(cd ui && npm run-script build)
	rice embed-go -i ./server
	go build -o bin/terrafire ./cmd/terrafire

.PHONY: test
test:
	go test $(ALL)

.PHONY: fmt
fmt:
	go fmt $(ALL)
