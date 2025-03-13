# Builder stage 
FROM golang:1.24.0-alpine AS builder
WORKDIR /

# Download dependencies for the common module
COPY common ./common
WORKDIR /common
RUN go mod download

# Download dependencies for rest module
WORKDIR /
COPY . .
WORKDIR /services/shop
RUN go mod download

# Build the final binary for shop service
WORKDIR /
RUN go build -o shop_service ./services/shop/cmd/shop

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /shop_service ./
CMD ["./shop_service"]
