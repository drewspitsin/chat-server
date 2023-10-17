FROM golang:1.21.2-alpine3.18 AS builder

COPY . /github.com/drewspitsin/chat-server/grpc/source/
WORKDIR /github.com/drewspitsin/chat-server/grpc/source/

RUN go mod download
RUN go build -o ./bin/chat-server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder github.com/drewspitsin/chat-server/grpc/source/bin/chat-server .

CMD ["./chat-server"]