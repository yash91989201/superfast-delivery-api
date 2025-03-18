FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/gateways/rest
COPY gateways/rest/go.mod gateways/rest/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/gateways/rest
RUN go build -ldflags="-s -w" -o /app/rest-gateway ./cmd/rest

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/rest-gateway ./

EXPOSE $PORT
CMD ["./rest-gateway"]
