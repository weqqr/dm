FROM golang:1.22 AS builder
WORKDIR /app
ENV CGO_ENABLED=0
RUN apt-get update && apt-get install -y protobuf-compiler
COPY go.mod go.sum Makefile /app/
RUN go mod download
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go \
               google.golang.org/grpc/cmd/protoc-gen-go-grpc
COPY cmd /app/cmd
COPY db /app/db
COPY internal /app/internal
COPY proto /app/proto

RUN make gateway

FROM scratch
WORKDIR /app
COPY db /app/db
COPY internal /app/internal
COPY --from=builder /app/gateway /app/gateway
COPY gateway.example.toml /app/gateway.toml
ENTRYPOINT ["/app/gateway"]