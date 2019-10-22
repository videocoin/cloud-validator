FROM golang:1.12.4 AS builder 

LABEL maintainer="VideoCoin"

WORKDIR /go/src/github.com/videocoin/cloud-validator

ADD ./ ./

RUN apt-get update && apt-get -y install libmediainfo-dev
RUN make build

FROM bitnami/minideb:jessie
RUN apt-get update && apt-get -y install mediainfo ffmpeg
COPY --from=builder /go/src/github.com/videocoin/cloud-validator/bin/validator ./
ENTRYPOINT [ "./validator" ]
