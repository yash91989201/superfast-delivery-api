FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/delivery
COPY services/delivery/go.mod services/delivery/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/delivery
RUN go build -ldflags="-s -w" -o /app/delivery_service ./cmd/delivery

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/delivery_service ./

EXPOSE $PORT
CMD ["./delivery_service"]
