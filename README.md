# ARCP — Adaptive Resource Coordination Protocol

ARCP is a lightweight, distributed protocol for Kubernetes that enables pods to share local resource metrics (CPU, memory) and coordinate usage dynamically in real-time.

### ✨ Features

- Pod-to-pod negotiation over gRPC
- Periodic exchange of resource metrics
- Suggested scaling or throttling actions
- Deployable as a sidecar or DaemonSet

### 📦 Getting Started

```bash
docker build -t your-dockerhub-user/arcp:latest .
kubectl apply -f k8s/deployment.yaml
