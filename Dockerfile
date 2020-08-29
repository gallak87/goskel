# Mix of https://stackoverflow.com/a/52990151 and https://tutorialedge.net/golang/go-docker-tutorial/
# TODO: Make end image smaller? currently has all source
FROM golang:1.15-alpine AS builder
LABEL maintainer="George Allakhverdyan <george.a@outlook.com>"

# TODO: build in container itself? copying from local for now
#COPY server ./
WORKDIR /go/src/github.com/gallak87/goskel
COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/server

#FROM base
#COPY --from=builder /go/src/github.com/gallak87/goskel /server
#RUN "chmod +x /server"
ENTRYPOINT ["./server"]
