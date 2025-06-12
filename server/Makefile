.PHONY: proto run build clean

# Compilar archivos .proto
proto:
	@mkdir -p pb
	protoc \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/user.proto

# Ejecutar el servidor
development:
	go run cmd/grpc/main.go

# Compilar binario
build:
	go build -o bin/server cmd/grpc/main.go

# Limpiar binario y código generado
clean:
	rm -rf bin/ pb/
