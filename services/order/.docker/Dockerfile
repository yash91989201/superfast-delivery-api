FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/order
COPY services/order/go.mod services/order/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/order
RUN go build -ldflags="-s -w" -o /app/order_service ./cmd/order

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/order_service ./

EXPOSE $PORT
CMD ["./order_service"]
