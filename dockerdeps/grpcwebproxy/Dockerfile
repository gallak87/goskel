#### Copied from: https://github.com/grpc/grpc-web/blob/master/net/grpc/gateway/docker/grpcwebproxy/Dockerfile
FROM golang:alpine

RUN apk add --no-cache curl git ca-certificates && \
  rm -rf /var/lib/apt/lists/*

ARG VERSION=0.12.0

WORKDIR /tmp

RUN curl -sS https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN wget https://github.com/improbable-eng/grpc-web/archive/v$VERSION.tar.gz

WORKDIR /go/src/github.com/improbable-eng/

RUN tar -zxf /tmp/v$VERSION.tar.gz -C .
RUN mv grpc-web-$VERSION grpc-web

WORKDIR /go/src/github.com/improbable-eng/grpc-web

RUN dep ensure && \
  go install ./go/grpcwebproxy

#ADD ./etc/localhost.crt /etc
#ADD ./etc/localhost.key /etc

ENTRYPOINT [ "/bin/sh", "-c", "exec /go/bin/grpcwebproxy \
  --backend_addr=goskel:9090 \
  --run_tls_server=false \
  --allow_all_origins " ]
