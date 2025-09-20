# Build stage
FROM golang:1.24.7-alpine3.22 AS builder
WORKDIR /app
COPY cmd ./cmd
# COPY internal ./cmd
COPY . .

# Build API executable
RUN go build -o main cmd/api/main.go

# Build migration executable
RUN go build -o migrate cmd/migrate/main.go

# Run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
COPY app.env .
COPY scripts .
COPY db/migration ./db/migration

RUN chmod +x /app/main /app/migrate /app/start.sh

EXPOSE 8080
ENTRYPOINT ["/app/start.sh" ]