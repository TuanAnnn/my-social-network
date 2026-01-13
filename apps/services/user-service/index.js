const express = require('express');
const app = express();
const PORT = 3000;

app.get('/', (req, res) => {
  res.json({
    message: "Hello from Node.js User Service!",
    timestamp: new Date(),
    service: "user-service"
  });
});

app.get('/users', (req, res) => {
  res.json([
    { id: 1, name: "Nguyen Van A", role: "Admin" },
    { id: 2, name: "Tran Thi B", role: "User" }
  ]);
});

app.listen(PORT, () => {
  console.log(`User Service running on port ${PORT}`);
});