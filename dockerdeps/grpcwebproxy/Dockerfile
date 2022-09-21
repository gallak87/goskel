#### Copied from: https://github.com/jfyne/docker-grpcwebproxy/blob/master/Dockerfile
FROM golang:alpine as builder

RUN apk update && apk upgrade
RUN go install github.com/improbable-eng/grpc-web/go/grpcwebproxy@latest

FROM alpine
COPY --from=builder /go/bin/grpcwebproxy .
EXPOSE 8080 9090

## TODO: parameterize addr
CMD ["sh", "-c", "./grpcwebproxy --backend_addr=goskel:9090 --run_tls_server=false --allow_all_origins"]
# --backend_addr=goskel:9090 --run_tls_server=false --allow_all_origins
