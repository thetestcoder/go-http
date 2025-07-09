#stage 1
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o my-app ./cmd/basic-http

#stage 2
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy only the binary from builder
COPY --from=builder /app/my-app .

# Copy config files if needed (assuming they're in config directory)
COPY --from=builder /app/internal/config ./config

EXPOSE 8888

CMD ["./my-app"]