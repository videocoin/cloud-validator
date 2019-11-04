FROM golang:1.12.4 AS builder 

LABEL maintainer="VideoCoin"

WORKDIR /go/src/github.com/videocoin/cloud-validator

ADD ./ ./

RUN make build

FROM bitnami/minideb:jessie
RUN echo 'deb http://www.deb-multimedia.org jessie main non-free' >> /etc/apt/sources.list
RUN echo 'deb-src http://www.deb-multimedia.org jessie main non-free' >> /etc/apt/sources.list
RUN apt-get update && \
    apt-get -y --force-yes install deb-multimedia-keyring && \
    apt-get -y --force-yes install ffmpeg

COPY --from=builder /go/src/github.com/videocoin/cloud-validator/bin/validator ./
RUN install_packages curl && GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
   curl -L -k https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 --output /bin/grpc_health_probe && chmod +x /bin/grpc_health_probe
ENTRYPOINT [ "./validator" ]
