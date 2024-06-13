FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod /app
COPY cmd /app/cmd
RUN go build ./cmd/coordinator

FROM scratch
WORKDIR /app
COPY --from=builder /app/coordinator /app/coordinator
CMD ["/app/coordinator"]
