FROM golang:1.11-alpine as builder
RUN apk add git gcc g++ make dep

WORKDIR /go/src/flyff
ADD . .

RUN dep ensure -vendor-only
RUN go get github.com/codegangsta/gin
