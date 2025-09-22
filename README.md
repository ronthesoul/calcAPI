# calcAPI

<img src="src/readme_image.png" alt="AWS Permissions Example" width="800">


A simple arithmetic API built with [Gin](https://github.com/gin-gonic/gin).  
Supports API key authentication, structured logging with [zap](https://github.com/uber-go/zap), and security headers.

---


## 📂 Project Sections

- [Task Summery](./task.md)
- [Terraform – Local Docker Deployment](./terraform/README.md)
- [Traffic Generator](./traffic-gen/README.md)
- [Prometheus + Grafana](./prometheus/README.md)
- [Helm Chart](./helm/README.md)


## Getting Started

Run the server:

```bash
go run main.go
```

By default the server listens on:  
```
http://localhost:8080
```

---

## Authentication

All endpoints (except `/health`) require an API key.  
Pass it in the request headers:

```
X-API-Key: <your_api_key>
```

You can generate a new API key using the `/token` endpoint.

---

## Endpoints

### 1. Health Check
Check if the server is alive.

```
GET /health
```

**Example:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{"status":"ok"}
```

---

### 2. Generate Token
Generates a new API key.

```
GET /token
```

**Example:**
```bash
curl http://localhost:8080/token
```

**Response:**
```json
{"token":"<generated_api_key>"}
```

---

### 3. Addition
Add two integers.

```
GET /add/:a/:b
```

**Example:**
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/add/2/4
```

**Response:**
```json
{"a":2,"b":4,"op":"add","result":6}
```

---

### 4. Subtraction
Subtract the second integer from the first.

```
GET /sub/:a/:b
```

**Example:**
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/sub/10/3
```

**Response:**
```json
{"a":10,"b":3,"op":"sub","result":7}
```

---

### 5. Multiplication
Multiply two integers.

```
GET /multiply/:a/:b
```

**Example:**
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/multiply/3/5
```

**Response:**
```json
{"a":3,"b":5,"op":"multiply","result":15}
```

---

### 6. Division
Divide the first integer by the second.

```
GET /divide/:a/:b
```

**Example:**
```bash
curl -H "X-API-Key: <your_api_key>" http://localhost:8080/divide/10/2
```

**Response:**
```json
{"a":10,"b":2,"op":"divide","result":5}
```

---

## Notes
- All operations expect integer values for `a` and `b`.  
- If you provide invalid inputs, the API will return a `400 Bad Request`.  
- Division by zero will return an error response.

