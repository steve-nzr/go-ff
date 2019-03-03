FROM golang:1-alpine
WORKDIR /go/src/go-ff

RUN apk add git gcc g++ make dep
RUN go get github.com/cortesi/modd/cmd/modd

ADD . .

RUN go get -v ./...
