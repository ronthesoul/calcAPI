# traffic-gen

Small Go load-generator that sends HTTP requests to **calcAPI** to simulate real traffic (health checks and calculator endpoints). Useful for local testing, demos, and metrics.

---

## Features
- Constant-rate request generation (RPS)
- Multiple workers (concurrency)
- Optional **X-API-Key** header
- Endpoint mix (/, /health, /add/:a/:b, /sub/:a/:b, /mul/:a/:b, /div/:a/:b)
- Prometheus-style basic timings/logs to stdout

---

## Quick Start

### 1) Run with Docker
```bash
# Build image (from this folder)
docker build -t calcapi-traffic:local .
```