FROM golang:1.15.2-alpine3.12

ENV GO111MODULE=on

RUN apk apk update && \
    apk add curl git && \
    go get github.com/pilu/fresh
