# ☸ KubeLaunch

> Deploy applications to Kubernetes in minutes — without writing YAML.

KubeLaunch is an open-source developer platform that simplifies Kubernetes deployments for application developers.

Instead of manually creating Kubernetes manifests, configuring services, debugging deployments, and learning complex infrastructure concepts, developers can deploy applications through a simple workflow with automatic stack detection and real-time observability.

---

## 🚨 The Problem

Modern developers want to ship applications quickly.

But deploying to Kubernetes often requires learning:

- Deployments
- Services
- Ingress
- Secrets
- ConfigMaps
- kubectl
- YAML syntax
- Kubernetes networking

For many frontend and backend developers, Kubernetes becomes a deployment bottleneck instead of a productivity tool.

---

## ✅ The Solution

KubeLaunch abstracts Kubernetes complexity behind a modern developer experience.

### Simple deployment flow

```text
GitHub Repository
       ↓
Automatic stack detection
       ↓
Deployment configuration
       ↓
Kubernetes manifest generation
       ↓
Deployment to cluster
       ↓
Live logs + application monitoring
```

---

## ✨ Features

### Current (v0.1)

- Dashboard overview
- Application deployment UI
- Stack auto-detection simulation
- Live log streaming simulation
- Metrics dashboard
- Environment variable management
- Modern Kubernetes-inspired UI

---

## 🧠 Planned Kubernetes Features

- Real Kubernetes deployments via client-go
- Deployment + Service generation
- Secret & ConfigMap management
- GitHub repository analysis
- Real-time pod logs streaming
- Pod health monitoring
- Namespace management
- Multi-cluster support

---

## 🏗 Architecture

```text
Frontend (Next.js)
        ↓
Backend API (Go + Gin)
        ↓
Kubernetes API (client-go)
        ↓
Kubernetes Cluster
```

---

## 🛠 Tech Stack

### Frontend
- Next.js 14
- TypeScript
- Tailwind CSS

### Backend
- Go
- Gin
- client-go

### Infrastructure
- Kubernetes
- Docker
- WebSockets

---

## 📂 Project Structure

```text
src/
├── app/
├── components/
├── lib/
└── types/
```

---

## 🚀 Getting Started

### Prerequisites

- Node.js 18+
- npm or yarn

### Installation

```bash
npm install
npm run dev
```

Open http://localhost:3000

---

## 🗺 Roadmap

### v0.2 — Real Kubernetes integration

- Go backend with Gin
- Kubernetes integration with client-go
- Real GitHub repository analysis
- Deployment generation
- Real-time logs streaming

### v0.3 — Platform features

- Multi-cluster support
- GitHub Actions integration
- Helm installation
- Team collaboration
- CNCF Sandbox submission

---

## 🤝 Contributing

Contributions are welcome.

KubeLaunch is currently in active development and aims to improve the Kubernetes developer experience for application developers.

---

## 📄 License

Apache 2.0