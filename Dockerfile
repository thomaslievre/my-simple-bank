# Build stage
FROM golang:1.22.1-alpine3.19 AS builder
WORKDIR /app
COPY cmd ./cmd
# COPY internal ./cmd
COPY . .

RUN go build -o main cmd/api/main.go
# RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
# COPY app.env .
COPY scripts .
# COPY --chown=user:group ./scripts/wait-for.sh ./scripts
# COPY --chown=user:group ./scripts/start.sh ./scripts
# COPY --chown=777 scripts/wait-for.sh .
# COPY db/migration ./db/migration

# RUN chmod +x /app/wait-for.sh
# RUN chmod +x /app/start.sh
# RUN ls -la /app
# RUN chmod 0755 /app/wait-for.sh
# RUN ["chmod", "+x", "/usr/src/app/start.sh"]
EXPOSE 8080 9090
CMD [ "/app/main" ]
ENTRYPOINT ["/app/start.sh" ]