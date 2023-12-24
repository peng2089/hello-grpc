PROTO_FILES=$(shell find api -name *.proto)
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: protoc
protoc:
	protoc --proto_path=./api \
		--go_out=paths=source_relative:./api \
		--go_opt=paths=source_relative \
		--go-grpc_out=paths=source_relative:./api \
		$(PROTO_FILES)

.PHONY: generate
generate:
	go mod tidy
	go generate ./...

.PHONY: build
build:
	mkdir -p bin/ && go build -o ./bin/ ./...

.PHONY: all
all:
	make init;
	make protoc;

.PHONY: clean
clean:
	find api -type f -name *.pb.go -delete