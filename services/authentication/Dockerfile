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
WORKDIR /services/authentication
RUN go mod download

# Build the final binary for authentication service
WORKDIR /
RUN go build -o authentication_service ./services/authentication/cmd/authentication

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /authentication_service ./
CMD ["./authentication_service"]
