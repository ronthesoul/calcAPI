# calcAPI – Golang Skeleton API Project

## Overview
**calcAPI** is a learning project built in **Go** using the **Gin** framework.  
The main goal was to design a clean and reusable **API skeleton** that can later serve as a template for more complex projects.  
The first implementation is a **simple calculator API** with essential middleware, observability, and deployment pipelines.

---

## Features

### Middleware
- **Authentication**  
  Validates API requests using an API key.  
  - Keys can be generated at `/token`.  
  - Keys are stored in a CSV file and expire after 1 hour.
- **Security Headers**  
  Adds browser security headers to protect against attacks such as clickjacking and enforces a CSP policy.
- **CORS**  
  Configured to control which origins are allowed to call the API.
- **Metrics**  
  Exposes `/metrics` endpoint for **Prometheus** scraping.
- **Structured Logging**  
  Integrated with **Zap logger**; logs accessible via container logs.
- **Rate Limiter**  
  Prevents abuse and mitigates DoS/DDoS attacks by limiting request rates.

### Routes
- **Calculator Operations**
  - `/add/:a/:b` → Addition
  - `/sub/:a/:b` → Subtraction
  - `/multiply/:a/:b` → Multiplication
  - `/divide/:a/:b` → Division
- **API Key Management**
  - `/token` → Generates a new API key (valid for 1 hour).
- **System Probes**
  - `/healthz` → Liveness check
  - `/readyz` → Readiness check
- **Shared Utilities**
  - `common.go` implements reusable error handling and response helpers.

### Traffic Generator
- A lightweight client written in Go to generate synthetic API traffic for testing and observability.

---

## CI/CD Pipeline
Implemented using **GitHub Actions**:
1. Clone repository.
2. Run **lint** and **unit tests**.
3. Build Docker image, run container, and validate API availability.
4. Tag image with incremented semantic version (e.g., `1.x+1.0`) and push to **Docker Hub**.
5. Automatically update:
   - Deployment manifest (used by ArgoCD).
   - Helm chart values (`values.yaml`).
6. Package Helm chart as an artifact.

---

## Deployment Options
- **Dockerfile** → build and run as a container.  
- **Kubernetes Deployment** → manifests managed by **ArgoCD**.  
- **Helm Chart** → packaged for Kubernetes deployments.  
- **Docker Compose** → runs both API and traffic generator together.  
- **Prometheus + Grafana Compose Setup** → observability stack for metrics.  
- **Terraform (AWS Provider)** → provision an EC2 instance running calcAPI.

---

## Summary
The **calcAPI** project demonstrates:
- A **production-like skeleton API** in Go with Gin.  
- Middleware for **security, observability, and reliability**.  
- Multiple deployment strategies (**Docker, K8s, Helm, Terraform**).  
- Automated **CI/CD pipeline** with Docker Hub publishing and GitOps integration (ArgoCD).  

This project provides a **professional foundation** for building and scaling more complex Go APIs.

