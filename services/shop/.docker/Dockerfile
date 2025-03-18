FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/shop
COPY services/shop/go.mod services/shop/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/shop
RUN go build -ldflags="-s -w" -o /app/shop_service ./cmd/shop

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/shop_service ./

EXPOSE $PORT
CMD ["./shop_service"]
