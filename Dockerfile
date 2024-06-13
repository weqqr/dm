FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum cmd/ /app/
RUN go build cmd/coordinator

FROM scratch
WORKDIR /app
COPY --from=builder /app/coordinator /app
CMD ["/app/coordinator"]
