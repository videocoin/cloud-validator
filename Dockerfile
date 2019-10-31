FROM golang:1.12.4 AS builder 

LABEL maintainer="VideoCoin"

WORKDIR /go/src/github.com/videocoin/cloud-validator

ADD ./ ./

RUN apt-get update && apt-get -y install libmediainfo-dev
RUN make build

FROM bitnami/minideb:jessie
RUN apt-get update && apt-get -y install mediainfo libmediainfo-dev wget xz-utils
RUN wget --no-check-certificate -O /tmp/ffmpeg.tar.xz https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
RUN tar -xJC /usr/bin --strip-components=1 -f /tmp/ffmpeg.tar.xz
RUN rm -rf \
    /usr/bin/ffmpeg-10bit \
    /usr/bin/ffserver \
    /tmp/ffmpeg.tar.xz
COPY --from=builder /go/src/github.com/videocoin/cloud-validator/bin/validator ./
RUN install_packages curl && GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
   curl -L -k https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 --output /bin/grpc_health_probe && chmod +x /bin/grpc_health_probe
ENTRYPOINT [ "./validator" ]
