## ユーザー作成
POST /register
- development
```
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "password": "secret123"}'
```
