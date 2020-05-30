ALL := ./cmd/terrafire ./internal/controller ./internal/database ./internal/executor ./internal/server ./internal/utils ./internal

.PHONY: build
build: web
	go build -o dist/terrafire ./cmd/terrafire

.PHONY: web
web:
	(cd web && npm run build)
	rice embed-go -i ./internal/server

.PHONY: test
test:
	go test $(ALL)

.PHONY: fmt
fmt:
	go fmt $(ALL)

.PHONY: run
run:
	go run ./cmd/terrafire
