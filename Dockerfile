# Mix of https://stackoverflow.com/a/52990151 and https://tutorialedge.net/golang/go-docker-tutorial/
# TODO: Make end image smaller? currently has all source
FROM golang:1.19 AS builder
LABEL maintainer="George Allakhverdyan <george.a@outlook.com>"

WORKDIR /go/src/github.com/gallak87/goskel
COPY . .
RUN go get -u ./...
RUN GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/server

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/gallak87/goskel /app
RUN ls -l /app
EXPOSE 8080 9090
ENTRYPOINT ["./server"]
