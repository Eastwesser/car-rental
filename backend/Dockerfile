# Build stage - Updated to Go 1.24.2
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copy only module files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy backend source code
COPY backend/ ./backend

# Build
WORKDIR /app/backend
RUN CGO_ENABLED=0 GOOS=linux go build -o backend .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/backend/backend .
COPY --from=builder /app/backend/images ./images

EXPOSE 8080
CMD ["./backend"]
