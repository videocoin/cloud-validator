FROM golang:latest AS builder 

LABEL maintainer="Videocoin" description="verify input vs output"


WORKDIR /go/src/github.com/videocoin/cloud-validator

ADD ./ ./

RUN make build

FROM jrottenberg/ffmpeg:4.0-ubuntu AS release

RUN apt update && apt upgrade -y
RUN apt install ca-certificates -y

COPY --from=builder /go/src/github.com/videocoin/cloud-validator/bin/validator ./

ENTRYPOINT [ "./validator" ]
