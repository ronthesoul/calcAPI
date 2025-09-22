# Terraform – Local Docker Deployment (calcAPI)

Run **calcAPI** locally in Docker using Terraform.

## Prerequisites
- Terraform ≥ 1.5
- Docker running locally
- Image available as `m4gapower/calcapi:latest` (or build it: `docker build -t m4gapower/calcapi:latest .`)

## Files
- `main.tf` – Terraform config (Docker provider, network, image, container)
- *(optional)* `variables.tf` – variable defaults

## Variables
| Name | Default | Description |
|---|---|---|
| `image` | `m4gapower/calcapi:latest` | Docker image to run |
| `container_name` | `calcapi` | Container name |
| `container_port` | `8080` | Port inside the container |
| `host_port` | `8080` | Port exposed on localhost |

## Quick Start
```bash
git clone https://github.com/ronthesoul/calcAPI.git
cd calcAPI/terraform

terraform fmt
terraform init
terraform apply -auto-approve

# Open the app
curl http://localhost:8080/healthz
# or visit http://localhost:8080
```

