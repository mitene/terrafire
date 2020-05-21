ALL := ./cmd/terrafire ./controller ./database ./executor ./server ./utils .

.PHONY: build
build: ui
	go build -o dist/terrafire ./cmd/terrafire

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
