# Build stage
FROM golang:1.23.1-alpine AS builder

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

RUN go build -o reverse-proxy-editor .

# Final stage
FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/reverse-proxy-editor .

EXPOSE 7777

CMD ["./reverse-proxy-editor"]
