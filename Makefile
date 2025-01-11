PROTO_PATH ?= pkg/proto
OUTPUT_PATH ?= ./pkg/gen/go

PROTO_FILES ?= \
    organization/organization.proto \
    warehouse/warehouse.proto \
	product/product.proto \
	crm/crm.proto \
	user/user.proto

proto:
	@protoc -I $(PROTO_PATH) $(addprefix $(PROTO_PATH)/, $(PROTO_FILES)) \
		--go_out=$(OUTPUT_PATH) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTPUT_PATH) --go-grpc_opt=paths=source_relative

build:
	@go build -o bin/core ./cmd/core/main.go

run: build
	@./bin/core --config=./config/local.yaml

test:
	@go test -v ./...

migrate:
	@go run ./cmd/migrate/main.go --config=./config/local.yaml --migrations-path=./migrations
