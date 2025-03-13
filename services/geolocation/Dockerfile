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
WORKDIR /services/geolocation
RUN go mod download

# Build the final binary for geolocation service
WORKDIR /
RUN go build -o geolocation_service ./services/geolocation/cmd/geolocation

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /geolocation_service ./
CMD ["./geolocation_service"]
