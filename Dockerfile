FROM golang:1.14 AS builder

WORKDIR /go/src/github.com/videocoin/cloud-validator
ADD ./ ./

RUN make build


FROM jrottenberg/ffmpeg:4.3-ubuntu1804

RUN apt-get update && \
    apt-get -y --force-yes install ffmpeg ca-certificates curl

COPY --from=builder /go/src/github.com/videocoin/cloud-validator/bin/validator /validator

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
   curl -L -k https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 --output /bin/grpc_health_probe && chmod +x /bin/grpc_health_probe

CMD [ "/validator" ]
