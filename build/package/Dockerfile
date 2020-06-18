FROM hashicorp/terraform:0.12.25 AS terraform

# build node
FROM node:alpine AS build-node

WORKDIR /build

COPY web/package*.json ./
RUN npm ci

COPY web .
RUN npm run build

# build go
FROM golang:alpine AS build-go

RUN apk update \
    && apk --no-cache add git make gcc g++ \
    && go get github.com/GeertJohan/go.rice/rice

WORKDIR /build

COPY go.* ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY --from=build-node /build/build ./web/build
RUN rice embed-go -i ./internal/server \
    && go build ./cmd/terrafire

# main
FROM alpine

RUN apk update && apk --no-cache add git

WORKDIR /app

COPY --from=terraform /bin/terraform /usr/local/bin/
COPY --from=build-go /build/terrafire /usr/local/bin/

ENTRYPOINT ["terrafire"]
