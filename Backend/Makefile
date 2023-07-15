proto-auth:
	protoc -I ./common/proto/auth -I ./common/proto \
	--go_out ./common/proto/auth/generated --go_opt paths=source_relative \
	--go-grpc_out ./common/proto/auth/generated --go-grpc_opt paths=source_relative \
	./common/proto/auth/auth.proto

proto-application:
	protoc -I ./common/proto/application -I ./common/proto \
	--go_out ./common/proto/application/generated --go_opt paths=source_relative \
	--go-grpc_out ./common/proto/application/generated --go-grpc_opt paths=source_relative \
	./common/proto/application/application.proto

proto-all: proto-auth proto-application