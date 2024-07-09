FROM golang:1.22 AS builder
WORKDIR /app
ENV CGO_ENABLED=0
RUN apt-get update && apt-get install -y protobuf-compiler
COPY go.mod go.sum Makefile /app/
RUN go mod download
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go \
               google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install github.com/rabbitmq/amqp091-go
COPY cmd /app/cmd
COPY db /app/db
COPY internal /app/internal
COPY proto /app/proto

RUN make rabbitmq

FROM scratch
WORKDIR /app
COPY db /app/db
COPY internal /app/internal
COPY --from=builder /app/rabbitmq /app/rabbitmq
COPY rabbitmq.example.toml /app/rabbitmq.toml
ENTRYPOINT ["/app/rabbitmq"]