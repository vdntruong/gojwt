# gojwt

## get token

```bash
curl -X POST http://localhost:8080/signin -H "Content-Type: application/json" -d '{"username": "truong", "password": "951951"}'
```

## get task status

```bash
curl -X GET http://localhost:8080/task/15 -H "Content-Type: application/json" -H "Authorization: Bearer <token-here>"
```
