FROM golang:1.11-alpine
RUN apk add git gcc g++ make dep

WORKDIR /go/src/go-ff
ADD . .

RUN go get -v ./...
RUN go get github.com/cortesi/modd/cmd/modd
