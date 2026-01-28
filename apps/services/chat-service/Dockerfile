# Stage 1: Builder (Dùng để build code Go)
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

# Stage 2: Runner (Cái này mới quan trọng!)
FROM alpine:latest

WORKDIR /root/

# === KIỂM TRA DÒNG NÀY ===
# Nó phải nằm ở dưới chữ FROM alpine:latest
RUN apk add --no-cache tzdata
# =========================

COPY --from=builder /app/main .
EXPOSE 3001
CMD ["./main"]