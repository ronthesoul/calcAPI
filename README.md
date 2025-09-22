# calcAPI

<p align="center">
  <img src="src/readme_image.png" alt="calcAPI Logo" width="250" />
</p>

<p align="center">
  <b>A lightweight arithmetic API built with <a href="https://github.com/gin-gonic/gin">Gin</a></b><br/>
  Featuring API key authentication, structured logging with <a href="https://github.com/uber-go/zap">Zap</a>, and secure HTTP headers.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/Framework-Gin-green?logo=go" />
  <img src="https://img.shields.io/badge/License-MIT-blue" />
</p>

---

## 📂 Project Sections

- [📌 Task Summary](./task.md)
- [⚙️ Terraform – Local Docker Deployment](./terraform/README.md)
- [📈 Traffic Generator](./traffic-gen/README.md)
- [📊 Prometheus + Grafana](./prometheus/README.md)
- [📦 Helm Chart](./helm/README.md)

---

## 🚀 Getting Started

Run the server locally:

```bash
go run main.go
```

The server starts on:

```
http://localhost:8080
```

---

## 🔑 Authentication

All endpoints (except `/health`) require an API key.

Add it to the request header:

```
X-API-Key: <your_api_key>
```

Generate a key via:

```
GET /token
```

---

## 📡 Endpoints

### Health Check
```
GET /health
```
Response:
```json
{"status":"ok"}
```

### Generate Token
```
GET /token
```
Response:
```json
{"token":"<generated_api_key>"}
```

### Addition
```
GET /add/:a/:b
```
Response:
```json
{"a":2,"b":4,"op":"add","result":6}
```

### Subtraction
```
GET /sub/:a/:b
```

### Multiplication
```
GET /multiply/:a/:b
```

### Division
```
GET /divide/:a/:b
```

---

## 📝 Notes

- Inputs must be integers (`a`, `b`).
- Division by zero returns an error.
- Invalid inputs return `400 Bad Request`.

---

## 📜 License

This project is licensed under the MIT License.
