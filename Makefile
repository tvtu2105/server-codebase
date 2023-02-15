GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOFMT := "goimports"
GO_PATH :=$GOPATH
SERVICE_PATH?=""

depend:
	@go get -v \
	golang.org/x/lint/golint \
	github.com/golangci/golangci-lint/cmd/golangci-lint \
	golang.org/x/tools/cmd/goimports \
	github.com/gogo/protobuf/protoc-gen-gogo@latest \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest \
	github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest \
	github.com/grpc-ecosystem/grpc-gateway \
	github.com/gogo/googleapis



generate: ## Generate proto
	protoc \
		-I $(SERVICE_PATH)/proto/ \
		-I $(GO_PATH)/src \
		-I $(GO_PATH)/pkg/mod/grpc-ecosystem/grpc-gateway/ \
		-I proto/common-proto/ \
		-I vendor/ \
		--descriptor_set_out=./descriptors.protoset\
		--include_source_info --include_imports -I. \
		--gogo_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(SERVICE_PATH)/pb \
		--grpc-gateway_out=\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(SERVICE_PATH)/pb \
		--swagger_out=$(SERVICE_PATH)/docs \
		--govalidators_out=gogoimport=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(SERVICE_PATH)/pb \
		$(SERVICE_PATH)/proto/*.proto