GO ?= go
PROTOC ?= protoc

.PHONY: all
all: coordinator

.PHONY: coordinator
coordinator: rpc domain
	$(GO) build ./cmd/coordinator

.PHONY: gateway
gateway: rpc domain
	$(GO) build ./cmd/gateway

.PHONY: rabbitmq
rabbitmq: rpc domain
	$(GO) build ./cmd/rabbitmq

.PHONY: clean
clean:
	find . -name "*\.pb\.go" -type f | xargs rm

.PHONY: rpc
rpc: proto
	mkdir -p internal/rpc
	$(PROTOC) --proto_path=proto/ --go-grpc_out=internal/rpc --go-grpc_opt=paths=source_relative  proto/*.proto

.PHONY: domain
domain: proto/domain
	mkdir -p internal/domain
	$(PROTOC) --proto_path=proto/ --go_out=internal --go_opt=paths=source_relative  proto/domain/*.proto

.PHONY: fmt
fmt:
	clang-format -i \
	proto/*.proto \
	proto/domain/*.proto

.PHONY: lint
lint:
	golangci-lint run -c .golangci.yaml --fix=false --color=always
