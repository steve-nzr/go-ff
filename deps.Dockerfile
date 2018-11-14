FROM golang:1.11-alpine
RUN apk add git gcc g++ make dep

WORKDIR /go/src/flyff
ADD . .

RUN go get -v ./...
