# ---------- Build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Copy go mod files first (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project (from root context)
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server

# ---------- Runtime stage ----------
FROM alpine:3.19

WORKDIR /app

# Create data directory
RUN mkdir -p /app/data

# Copy binary
COPY --from=builder /app/server /app/server

EXPOSE 8080

ENTRYPOINT ["/app/server"]
