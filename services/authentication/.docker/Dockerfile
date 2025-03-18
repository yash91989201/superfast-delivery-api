FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/authentication
COPY services/authentication/go.mod services/authentication/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/authentication
RUN go build -ldflags="-s -w" -o /app/authentication_service ./cmd/authentication

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/authentication_service ./

EXPOSE $PORT
CMD ["./authentication_service"]

