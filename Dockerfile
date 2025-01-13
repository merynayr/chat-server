FROM golang:1.23-alpine AS builder

COPY . /github.com/merynayr/chat-server/source/
WORKDIR /github.com/merynayr/chat-server/source/

RUN go mod download
RUN go build -o ./bin/test_server cmd/grpc_server/main.go

FROM alpine:latest


WORKDIR /root/

COPY --from=builder /github.com/merynayr/chat-server/source/bin/chat_server .

COPY local.env .

CMD ["./chat_server", "-config-path=local.env"]