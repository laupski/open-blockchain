all: build

.PHONY: build
build:
	go build -o obc

.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	api/proto/blockchain.proto