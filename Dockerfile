FROM golang:1.12.4 AS builder 

LABEL maintainer="Videocoin" description="verify input vs output"

WORKDIR /go/src/github.com/videocoin/cloud-validator

ADD ./ ./

RUN make build

FROM bitnami/minideb:jessie
COPY --from=builder /go/src/github.com/videocoin/cloud-validator/bin/validator ./
ENTRYPOINT [ "./validator" ]
