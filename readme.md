# My Social Network - Cấu trúc dự án
```
my-social-network/
├── apps/
│   ├── api-gateway/         # (Node.js hoặc Nginx config)
│   ├── user-service/        # (Node.js + NestJS)
│   │   ├── src/
│   │   ├── package.json
│   │   └── Dockerfile
│   ├── post-service/        # (Node.js + NestJS)
│   ├── chat-service/        # (Go + Gin)
│   │   ├── cmd/main.go
│   │   ├── internal/handler
│   │   ├── go.mod
│   │   └── Dockerfile
│   └── media-service/       # (Go + Fiber)
├── libs/
│   ├── shared-types/        # File .proto (gRPC) dùng chung cho cả Node và Go
│   └── docker-compose.yml   # Chạy toàn bộ hệ thống
```

## Mô tả các service

### API Gateway
- **Tech stack**: Node.js hoặc Nginx config
- **Chức năng**: Điều hướng request đến các microservices

### User Service
- **Tech stack**: Node.js + NestJS
- **Chức năng**: Quản lý người dùng, authentication, authorization

### Post Service
- **Tech stack**: Node.js + NestJS
- **Chức năng**: Quản lý bài viết, timeline, tương tác

### Chat Service
- **Tech stack**: Go + Gin
- **Chức năng**: Xử lý tin nhắn real-time, chat

### Media Service
- **Tech stack**: Go + Fiber
- **Chức năng**: Upload, xử lý và lưu trữ media (ảnh, video)

## Shared Libraries

### shared-types
- Chứa file `.proto` (Protocol Buffers) để định nghĩa gRPC contracts
- Dùng chung cho cả Node.js và Go services

## Docker
- `docker-compose.yml`: Orchestration cho toàn bộ hệ thống microservices