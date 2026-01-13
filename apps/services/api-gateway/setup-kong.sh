#!/bin/bash

# 1. Đăng ký Node.js User Service vào Kong
echo "Configuring User Service..."
curl -i -X POST http://localhost:8001/services \
  --data name=user-service \
  --data url='http://user-service:3000'

# 2. Tạo đường dẫn (Route) cho User Service
# Gọi localhost:8000/api/users -> Vào Node.js
curl -i -X POST http://localhost:8001/services/user-service/routes \
  --data 'paths[]=/api/users' \
  --data name=user-routes

# 3. Đăng ký Go Chat Service vào Kong
echo "Configuring Chat Service..."
curl -i -X POST http://localhost:8001/services \
  --data name=chat-service \
  --data url='http://chat-service:8080'

# 4. Tạo đường dẫn (Route) cho Chat Service
# Gọi localhost:8000/api/chat -> Vào Go
curl -i -X POST http://localhost:8001/services/chat-service/routes \
  --data 'paths[]=/api/chat' \
  --data name=chat-routes

echo "Done!"