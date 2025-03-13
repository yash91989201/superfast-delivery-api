# Builder stage 
FROM golang:1.24.0-alpine AS builder
WORKDIR /

# Download dependencies for the common module
COPY common ./common
WORKDIR /common
RUN go mod download

# Download dependencies for graphql module
WORKDIR /
COPY . .
WORKDIR /gateways/graphql
RUN go mod download

# Build the final binary for graphql gateway
WORKDIR /
RUN go build -o graphql-gateway ./gateways/graphql/cmd/graphql

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /graphql-gateway ./
EXPOSE $PORT
CMD ["./graphql-gateway"]
