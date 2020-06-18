ALL := ./cmd/terrafire ./internal/api ./internal/controller ./internal/database ./internal/manifest ./internal/runner ./internal/server ./internal/utils

.PHONY: setup
setup:
	go get -u github.com/GeertJohan/go.rice/rice \
	          github.com/golang/protobuf/protoc-gen-go

.PHONY: build
build: web
	go build -o dist/terrafire ./cmd/terrafire

.PHONY: web
web:
	make -C web build
	rice embed-go -i ./internal/server

.PHONY: proto
proto:
	protoc -I api --go_out=plugins=grpc:./internal/api --go_opt=paths=source_relative api/*.proto

.PHONY: test
test:
	go test $(ALL)

.PHONY: fmt
fmt:
	go fmt $(ALL)

.PHONY: push-docker-dev
push-docker-dev:
	docker build . -t terrafire -f build/package/Dockerfile
	docker tag terrafire mitene/terrafire:dev-$$(git symbolic-ref --short HEAD)
	docker push mitene/terrafire:dev-$$(git symbolic-ref --short HEAD)
