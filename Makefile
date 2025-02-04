PROTO_PATH       ?= pkg/proto
OUTPUT_PATH_GO   ?= pkg/gen/go
OUTPUT_PATH_PY   ?= pkg/gen/py

PROTO_FILES ?= \
	user/user.proto \
    organization/organization.proto \
    warehouse/warehouse.proto \
	product/product.proto \
	crm/crm.proto \
	order/order.proto \
	advertisement/advertisement.proto

proto:
	@protoc -I $(PROTO_PATH) $(addprefix $(PROTO_PATH)/, $(PROTO_FILES)) \
		--go_out=$(OUTPUT_PATH_GO) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTPUT_PATH_GO) --go-grpc_opt=paths=source_relative

	source venv/bin/activate && \
	python -m grpc_tools.protoc \
		-I $(PROTO_PATH) \
		$(addprefix $(PROTO_PATH)/, $(PROTO_FILES)) \
		--python_out=$(OUTPUT_PATH_PY) \
		--grpc_python_out=$(OUTPUT_PATH_PY)


build:
	@templ generate
	@go build -o bin/core ./cmd/core/main.go

run: build
	@./bin/core --config=./config/local.yaml

test:
	@go test -v ./...

migrate:
	@go run ./cmd/migration/main.go --config=./config/local.yaml --migrations-path=./migrations
