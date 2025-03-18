FROM golang:1.24.0-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download && go mod verify

WORKDIR /app/gateways/graphql
COPY gateways/graphql/go.mod gateways/graphql/go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY . .

WORKDIR /app/gateways/graphql
RUN go build -ldflags="-s -w" -o /app/graphql-gateway ./cmd/graphql

FROM scratch AS final
WORKDIR /
COPY --from=builder /app/graphql-gateway ./

EXPOSE $PORT
CMD ["./graphql-gateway"]
