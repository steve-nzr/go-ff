FROM golang:1.11-alpine as builder
RUN apk add git gcc g++ make dep

WORKDIR /go/src/flyff
ADD Gopkg.lock .
ADD Gopkg.toml .

RUN dep ensure -vendor-only
