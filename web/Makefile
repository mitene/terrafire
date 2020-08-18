PROTOC_GEN_TS_PATH := ./node_modules/.bin/protoc-gen-ts
PROTOC_OUT_DIR := ./src/api

.PHONY: setup
setup:
	npm install --no-progress

.PHONY: build
build:
	npm run build

.PHONY: proto
proto:
	protoc \
		--plugin="protoc-gen-ts=$(PROTOC_GEN_TS_PATH)" \
	    --js_out="import_style=commonjs,binary:$(PROTOC_OUT_DIR)" \
	    --ts_out="service=grpc-web:$(PROTOC_OUT_DIR)" \
	    -I ../api \
	    ../api/*.proto
# disable eslint
# see. https://github.com/improbable-eng/grpc-web/issues/96#issuecomment-347871452
	@for f in $(PROTOC_OUT_DIR)/*.js; do \
  		(echo '/* eslint-disable */'; cat $$f) > $$f.tmp; \
  		mv $$f.tmp $$f; \
	done
