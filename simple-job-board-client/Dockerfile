FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/app/main.go

# Stage 2
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ENTRYPOINT ["./app"]