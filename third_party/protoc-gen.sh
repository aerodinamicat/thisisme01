protoc --proto_path=api/v1 --proto_path=third_party --go-grpc_out=pkg/api/v1 user-service.proto
protoc --proto_path=api/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 user-service.proto
protoc --proto_path=api/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/v1 user-service.proto
