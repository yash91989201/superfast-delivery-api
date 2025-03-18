FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/services/geolocation
COPY services/geolocation/go.mod services/geolocation/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/services/geolocation
RUN go build -ldflags="-s -w" -o /app/geolocation_service ./cmd/geolocation

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/geolocation_service ./

EXPOSE $PORT
CMD ["./geolocation_service"]
