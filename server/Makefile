.PHONY: proto run build clean

proto:
	@mkdir pb
	protoc \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/user/user.proto \
		proto/application/application.proto

development:
	go run cmd/grpc/main.go

build:
	go build -o bin/server cmd/grpc/main.go

clean:
	rm -rf bin/ pb/
