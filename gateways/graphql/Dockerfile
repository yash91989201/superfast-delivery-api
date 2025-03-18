FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Cache dependencies for common module
COPY common/go.mod common/go.sum ./common/
WORKDIR /app/common
RUN go mod download

# Cache dependencies for graphql module
WORKDIR /app/gateways/graphql
COPY gateways/graphql/go.mod gateways/graphql/go.sum ./
RUN go mod download

# Now copy the actual source code
WORKDIR /app
COPY . .

# Build the GraphQL Gateway
WORKDIR /app/gateways/graphql
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/graphql-gateway ./cmd/graphql

# Final minimal image
FROM scratch
WORKDIR /
COPY --from=builder /app/graphql-gateway ./
EXPOSE $PORT
CMD ["./graphql-gateway"]
