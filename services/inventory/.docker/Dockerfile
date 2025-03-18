FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/inventory
COPY services/inventory/go.mod services/inventory/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/inventory
RUN go build -ldflags="-s -w" -o /app/inventory_service ./cmd/inventory

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/inventory_service ./

EXPOSE $PORT
CMD ["./inventory_service"]
