# Builder stage 
FROM golang:1.23.0-alpine AS builder
WORKDIR /

# Download dependencies for the common module
COPY common ./common
WORKDIR /common
RUN go mod download

# Download dependencies for rest module
WORKDIR /
COPY . .
WORKDIR /services/product
RUN go mod download

# Build the final binary for product service
WORKDIR /
RUN go build -o product_service ./services/product/cmd/product

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /product_service ./
CMD ["./product_service"]

