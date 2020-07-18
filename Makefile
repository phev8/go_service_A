.PHONY: api

PROTO_BUILD_DIR = ../../..
PROTO_TARGET = ./pkg/api


help:
	@echo "Service building targets"
	@echo "  api: compile protobuf files for go"

api:
	git submodule update --recursive --remote
	find "$(PROTO_TARGET)" -type f -delete
	find ./api/service_proto/service_a/*.proto -maxdepth 1 -type f -exec protoc {} --proto_path=./api/service_proto --go_out=plugins=grpc:$(PROTO_BUILD_DIR) \;

