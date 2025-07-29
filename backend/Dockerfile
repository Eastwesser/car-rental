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

# Copy wait-for-it.sh FIRST
COPY backend/wait-for-it.sh .
RUN chmod +x wait-for-it.sh

# Then copy the binary
COPY --from=builder /app/backend-binary ./backend
COPY --from=builder /app/backend/images ./images

# Install dependencies
RUN apk add --no-cache bash postgresql-client

EXPOSE 8080
CMD ["./wait-for-it.sh", "db", "5432", "postgres", "luxury", "carsharing", "--", "./backend"]