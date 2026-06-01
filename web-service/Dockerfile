# Stage 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o api .

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 6769

CMD ["./api"]
