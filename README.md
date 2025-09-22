<h1 align="center">💀 calcAPI – Skeleton Calculator API</h1>

<p align="center">
  <img src="src/readme_image.png" alt="calcAPI Logo" width="300" />
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

## 📂 Project Structure

- [📌 Task Summary](./task.md)
- [⚙️ Terraform – Local Docker Deployment](./terraform/README.md)
- [📈 Traffic Generator](./traffic-gen/README.md)
- [📊 Prometheus + Grafana](./prometheus/README.md)
- [📦 Helm Chart](./helm/README.md)

```
calcAPI/
├── main.go
├── go.mod
├── go.sum
├── src/
│   └── readme_image.png
├── terraform/
│   └── README.md
├── traffic-gen/
│   └── README.md
├── prometheus/
│   └── README.md
├── helm/
│   └── README.md
└── task.md
```

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

Generate a key:
```bash
curl http://localhost:8080/token
```

Add it to your requests:
```
-H "X-API-Key: <your_api_key>"
```

---

## 📡 Endpoints (with curl)

### Health Check
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/health
```

### Generate Token
```bash
curl http://localhost:8080/token
```

### Addition
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/add/2/4
```

### Subtraction
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/sub/10/3
```

### Multiplication
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/multiply/3/5
```

### Division
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/divide/10/2
```

---

## 📝 Notes

- Inputs must be integers (`a`, `b`).
- Division by zero returns an error.
- Invalid inputs return `400 Bad Request`.

---

## 👤 Author

**Ron (ronthesoul)**  
[GitHub Profile](https://github.com/ronthesoul)

