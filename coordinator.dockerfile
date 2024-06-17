FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum Makefile /app/
COPY cmd /app/cmd
COPY internal /app/internal
COPY proto /app/proto
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go \
               google.golang.org/grpc/cmd/protoc-gen-go-grpc
ENV CGO_ENABLED=0
RUN make coordinator

FROM scratch
WORKDIR /app
COPY --from=builder /app/coordinator /app/coordinator
COPY coordinator.example.toml /app/coordinator.toml
ENTRYPOINT ["/app/coordinator"]