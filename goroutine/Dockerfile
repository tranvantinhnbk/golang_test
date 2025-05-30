# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
