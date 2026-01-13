# --- SỬA DÒNG NÀY (1.20 -> 1.23) ---
FROM golang:1.23-alpine

WORKDIR /app

# 1. Cài Git
RUN apk add --no-cache git

# 2. Proxy
ENV GOPROXY=https://proxy.golang.org,direct

# 3. Copy code
COPY main.go .

# 4. Tự động tải thư viện
RUN go mod init chat-service || true
RUN go mod tidy

# 5. Build
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]