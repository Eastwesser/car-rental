# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/backend/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend-binary .

# Final stage
FROM alpine:latest
WORKDIR /app

# Install dependencies first for better caching
RUN apk add --no-cache bash postgresql-client

# Copy wait-for-it.sh and binary
COPY backend/wait-for-it.sh .
RUN chmod +x wait-for-it.sh
COPY --from=builder /app/backend-binary ./backend
COPY --from=builder /app/backend/images ./images

EXPOSE 8080
CMD ["./wait-for-it.sh", "db", "5432", "postgres", "luxury", "carsharing", "--", "./backend"]