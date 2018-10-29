FROM flyff-deps as builder
WORKDIR /go/src/flyff

ADD . .

RUN sh build.sh

FROM alpine:3.8
WORKDIR /go/bin
COPY --from=builder /go/src/flyff/bin .
COPY --from=builder /go/src/flyff/config/docker.env ../config/.env
