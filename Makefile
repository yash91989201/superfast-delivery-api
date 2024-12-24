include .env

proto-def:
	protoc --proto_path="common/pb" \
		--go_out="common/pb" --go_opt=paths=source_relative \
		--go-grpc_out="common/pb" --go-grpc_opt=paths=source_relative \
		common/pb/*.proto

graphql-def:
	cd gateways/graphql && go run github.com/99designs/gqlgen generate
