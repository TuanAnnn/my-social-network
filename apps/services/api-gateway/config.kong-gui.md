# Quick Start

Sau khi chạy `docker-compose up -d`, hệ thống sẽ hoạt động tại:

* **Cổng truy cập (Client):** `http://localhost:8000`
* **Cổng cấu hình (Admin):** `http://localhost:8001`

# Hướng dẫn Config trên Kong GUI (Konga)

## 1. Đầu tiên điền name và Kong Url
   - **Name:** `Social App Kong`
   - **Kong Admin Url:** `http://kong:8001`

Truy cập: `http://localhost:1337`

## 1. Kết nối Node.js Service


### Bước 1: Tạo Service
1. Menu **SERVICES** -> **+ ADD NEW SERVICE**.
2. Điền thông tin:
   - **Name:** `user-service`
   - **Protocol:** `http`
   - **Host:** `user-service`
   - **Port:** `3000`
   - **Path:** `/`
3. Bấm **SUBMIT SERVICE**.

### Bước 2: Tạo Route
1. Bấm vào tên **`user-service`** vừa tạo.
2. Tab **ROUTES** -> **+ ADD ROUTE**.
3. Điền thông tin:
   - **Name:** `user-route`
   - **Paths:** Gõ `/api/users` -> **BẮT BUỘC BẤM ENTER** (để biến thành thẻ màu xám).
   - **Protocols:** Gõ `http` (Enter), `https` (Enter).
4. Bấm **SUBMIT ROUTE**.

## 2. Kết nối Go Chat Service

### Bước 1: Tạo Service
1. Menu **SERVICES** -> **+ ADD NEW SERVICE**.
2. Điền thông tin:
   - **Name:** `chat-service`
   - **Protocol:** `http`
   - **Host:** `chat-service`
   - **Port:** `8080`
   - **Path:** `/`
3. Bấm **SUBMIT SERVICE**.

## 2. Kết nối Go Chat Service

### Bước 1: Tạo Service
1. Menu **SERVICES** -> **+ ADD NEW SERVICE**.
2. Điền thông tin:
   - **Name:** `chat-service`
   - **Protocol:** `http`
   - **Host:** `chat-service`
   - **Port:** `8080`
   - **Path:** `/`
3. Bấm **SUBMIT SERVICE**.

### Bước 2: Tạo Route
1. Bấm vào tên **`chat-service`** vừa tạo.
2. Tab **ROUTES** -> **+ ADD ROUTE**.
3. Điền thông tin:
   - **Name:** `chat-route`
   - **Paths:** Gõ `/api/chat` -> **BẮT BUỘC BẤM ENTER** (để biến thành thẻ màu xám).
   - **Protocols:** Gõ `http` (Enter), `https` (Enter).
   - **Strip Path:** `No` (Lưu ý: Chọn No để Go nhận được đầy đủ đường dẫn).
4. Bấm **SUBMIT ROUTE**.