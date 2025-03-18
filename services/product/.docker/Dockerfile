FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/product
COPY services/product/go.mod services/product/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/product
RUN go build -ldflags="-s -w" -o /app/product_service ./cmd/product

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/product_service ./

EXPOSE $PORT
CMD ["./product_service"]
