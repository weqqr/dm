GO ?= go
PROTOC ?= protoc

.PHONY: all
all: coordinator

.PHONY: coordinator
coordinator: rpc domain
	$(GO) build ./cmd/coordinator

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