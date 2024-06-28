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

RUN make coordinator

FROM scratch
WORKDIR /app
COPY db /app/db
COPY internal /app/internal
COPY --from=builder /app/coordinator /app/coordinator
COPY coordinator.example.toml /app/coordinator.toml
ENTRYPOINT ["/app/coordinator"]