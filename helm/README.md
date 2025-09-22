# calcAPI – Helm Chart

This folder contains a **Helm chart** to deploy `calcAPI` onto Kubernetes.

---

## 📦 Prerequisites
- Kubernetes cluster (kind, k3s, minikube, EKS, etc.)
- [Helm 3+](https://helm.sh/docs/intro/install/)

---

## 🚀 Install

From the repo root or directly inside this `helm/` folder:

```bash
helm install calcapi ./helm
```

This will create a Deployment, Service, and required metadata.

---

## ⚙️ Values

You can override defaults with a `values.yaml` file or `--set` flags.

| Key              | Default                | Description                     |
|------------------|------------------------|---------------------------------|
| `replicaCount`   | `1`                    | Number of pod replicas          |
| `image.repository` | `m4gapower/calcapi`  | Docker image repo               |
| `image.tag`      | `latest`               | Image tag                       |
| `image.pullPolicy` | `IfNotPresent`       | Image pull policy               |
| `service.type`   | `ClusterIP`            | Service type (ClusterIP/NodePort/LoadBalancer) |
| `service.port`   | `8080`                 | Service port                    |


---

## 📝 Example Overrides

Create `my-values.yaml`:

```yaml
replicaCount: 2

image:
  repository: m4gapower/calcapi
  tag: "1.18.0"
  pullPolicy: IfNotPresent

service:
  type: LoadBalancer
  port: 80

resources:
  requests:
    cpu: "200m"
    memory: "256Mi"
  limits:
    cpu: "500m"
    memory: "500Mi"
```

Install with overrides:

```bash
helm install calcapi ./helm -f my-values.yaml
```

---

## 🔄 Upgrade

When you change values:

```bash
helm upgrade calcapi ./helm -f my-values.yaml
```

---

## 🧹 Uninstall

```bash
helm uninstall calcapi
```

