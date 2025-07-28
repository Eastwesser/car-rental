# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o backend .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/backend .
COPY --from=builder /app/images ./images

EXPOSE 8080
CMD ["./backend"]