# Builder stage 
FROM golang:1.23.0-alpine AS builder
WORKDIR /

# Download dependencies for the common module
COPY common ./common
WORKDIR /common
RUN go mod download

# Download dependencies for user module
WORKDIR /
COPY . .
WORKDIR /services/user
RUN go mod download

# Build the final binary for user service
WORKDIR /
RUN go build -o user_service ./services/user/cmd/user

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /user_service ./
CMD ["./user_service"]
