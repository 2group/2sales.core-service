PROTO_PATH       ?= pkg/proto
OUTPUT_PATH_GO   ?= pkg/gen/go
OUTPUT_PATH_PY   ?= pkg/gen/py
OUTPUT_PATH_CPP  ?= pkg/gen/cpp

PROTO_FILES ?= \
	user/user.proto \
	organization/organization.proto \
	product/product.proto \
	crm/crm.proto \
	order/order.proto \
	advertisement/advertisement.proto \
	customer/customer.proto \
	service/service.proto \
	b2c_service_order/b2c_service_order.proto \
	# warehouse/warehouse.proto \

proto:
	@protoc -I $(PROTO_PATH) $(addprefix $(PROTO_PATH)/, $(PROTO_FILES)) \
		--go_out=$(OUTPUT_PATH_GO) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTPUT_PATH_GO) --go-grpc_opt=paths=source_relative \
		# --cpp_out=$(OUTPUT_PATH_CPP) \
		# --grpc_out=$(OUTPUT_PATH_CPP) --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`

	# source venv/bin/activate && \
	# python -m grpc_tools.protoc \
	# 	-I $(PROTO_PATH) \
	# 	$(addprefix $(PROTO_PATH)/, $(PROTO_FILES)) \
	# 	--python_out=$(OUTPUT_PATH_PY) \
	# 	--grpc_python_out=$(OUTPUT_PATH_PY)

build:
	@go build -o bin/core ./cmd/core/main.go

run: build
	@./bin/core --config=./config/local.yaml

test:
	@go test -v ./...

migrate:
	@go run ./cmd/migration/main.go --config=./config/local.yaml --migrations-path=./migrations

