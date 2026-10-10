build_proto:
	protoc --proto_path=proto/v1 --go_out=pkg/pb/v1 --go_opt=paths=source_relative \
	       --go-grpc_out=pkg/pb/v1 --go-grpc_opt=paths=source_relative \
	       proto/v1/calculator.proto