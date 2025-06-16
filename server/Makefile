.PHONY: proto run build clean

proto:
	@mkdir pb
	protoc \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/user/user.proto \
		proto/application/application.proto \
		proto/application_role/application_role.proto \
		proto/user_application_role/user_application_role.proto \
		proto/module/module.proto \
		proto/module_role_permission/module_role_permission.proto \
		proto/oauth_token/oauth_token.proto \
		proto/structural_position/structural_position.proto \
		proto/organic_unit/organic_unit.proto \
		proto/password_history/password_history.proto \
		proto/password_reset/password_reset.proto \
		proto/user_session/user_session.proto

development:
	go run cmd/grpc/main.go

build:
	go build -o bin/server cmd/grpc/main.go

clean:
	rm -rf bin/ pb/
