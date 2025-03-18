FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/user
COPY services/user/go.mod services/user/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/user
RUN go build -ldflags="-s -w" -o /app/user_service ./cmd/user

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/user_service ./

EXPOSE $PORT
CMD ["./user_service"]
