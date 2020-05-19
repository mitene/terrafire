ALL := ./cmd/terrafire ./core ./database ./runner ./server ./utils

.PHONY: build
build: ui
	go build -o bin/terrafire ./cmd/terrafire

.PHONY: ui
ui:
	(cd ui && npm run build)
	rice embed-go -i ./server

.PHONY: test
test:
	go test $(ALL)

.PHONY: fmt
fmt:
	go fmt $(ALL)

.PHONY: run
run:
	go run ./cmd/terrafire
