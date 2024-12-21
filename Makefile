include .env

proto_def:
	protoc --proto_path="common/pb" \
		--go_out="common/pb" --go_opt=paths=source_relative \
		--go-grpc_out="common/pb" --go-grpc_opt=paths=source_relative \
		common/pb/*.proto

authentication_db_up:
	migrate -path ./services/authentication/db/migrations \
	-database "mysql://${AUTHENTICATION_DB_USER}:${AUTHENTICATION_DB_PASSWORD}@tcp(localhost:${AUTHENTICATION_DB_PORT})/${AUTHENTICATION_DB_NAME}" \
	up

authentication_db_down:
	migrate  -path ./services/authentication/db/migrations \
	-database "mysql://${AUTHENTICATION_DB_USER}:${AUTHENTICATION_DB_PASSWORD}@tcp(localhost:${AUTHENTICATION_DB_PORT})/${AUTHENTICATION_DB_NAME}" \
	down

