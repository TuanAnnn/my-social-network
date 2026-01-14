FROM node:18-alpine

WORKDIR /app

COPY package.json .

# === THÊM DÒNG NÀY ===
# Cài đặt OpenSSL cho Alpine
RUN apk add --no-cache openssl

RUN npm install

COPY . .

# Biến giả để build (như đã làm trước đó)
ENV DATABASE_URL="postgresql://dummy:dummy@localhost:5432/dummy"

RUN npx prisma generate
RUN npm run build

EXPOSE 3000
CMD ["npm", "run", "start:prod"]