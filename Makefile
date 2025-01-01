include .env

proto-def:
	find common/pb -name "*.proto" -exec buf format --write {} \;
	protoc --proto_path="common/proto" \
		--go_out="common/pb" --go_opt=paths=source_relative \
		--go-grpc_out="common/pb" --go-grpc_opt=paths=source_relative \
		common/proto/*.proto

graphql-def:
	cd gateways/graphql && go run github.com/99designs/gqlgen generate
