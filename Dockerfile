# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o syslogserver

# Final stage
FROM alpine:3.19
RUN adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=builder /app/syslogserver .
USER appuser
EXPOSE 6601/tcp
CMD ["./syslogserver"]
