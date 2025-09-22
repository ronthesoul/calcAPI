# ArgoCD Setup for calcAPI

This section explains how to deploy **ArgoCD** into your Kubernetes cluster and configure it to manage the `calcAPI` manifests in this repo.

---

## 1. Install ArgoCD

Apply the official ArgoCD install manifests into the `argocd` namespace:

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

## 2. Check if the pods are running

```bash
kubectl get pods -n argocd
```

## 3. Access the ArgoCD UI
Expose the ArgoCD server (choose one of the following):

```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```
Then open: https://localhost:8080

## 4. Get Initial Admin Password
```bash
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d; echo
```
## 5. Copy the app.yaml to a file and apply it
```bash
kubectl apply -f app.yaml
```
### 6. Verify Deployment
Check that ArgoCD synced the app:
```bash
kubectl get applications -n argocd
kubectl get pods -n default
```
