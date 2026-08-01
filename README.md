# GoRisk AI Core 🛡️⚡

An enterprise-grade, resilient, event-driven credit risk evaluation engine built with **Golang**, **React**, **PostgreSQL**, **MongoDB**, **Redis**, **Apache Kafka**, and **Google Gemini AI**.

Designed with **Clean (Hexagonal) Architecture** and containerized for high-availability cloud deployments (AWS EKS simulation).

---

## 🏗️ Architecture Overview

The system operates as a distributed microservice handling asynchronous events and fault-tolerant AI evaluations:

```text
[ React Dashboard (Port 5173) ]
             │
             ▼ (HTTP / REST)
  [ Gin HTTP Server (Go - Port 8080) ]
             │
   ┌─────────┼───────────────────────────┬────────────────────────┐
   ▼         ▼                           ▼                        ▼
[ GORM ]  [ Redis Cache ]      [ MongoDB Atlas ]      [ Apache Kafka Broker ]
 (Postgres) (Cache-Aside)     (Immutable Audit Logs)  (Topic: risk_evaluations)
   │
   ▼
[ Gemini 3.5/3.6 AI Engine ] (Exponential Backoff + Graceful Degradation)

```

### Key Engineering Features

* **Dual Persistence:** Structured transactional records in **PostgreSQL** paired with unstructured, immutable audit trails in **MongoDB**.
* **Cache-Aside Strategy:** Primary application queries hit **Redis** memory with a 5-minute TTL before hitting the relational database.
* **Event-Driven Architecture (EDA):** Asynchronous message broadcasting to **Apache Kafka (KRaft mode)** on every credit evaluation.
* **Resilient AI Pipeline:** Credit analysis powered by **Google Gemini API** with automatic **Exponential Backoff retries (HTTP 429)** and **Graceful Degradation** fallbacks.
* **CORS & Security:** Configured cross-origin resource sharing middleware for Vite/React frontend integration.

---

## 🛠️ Tech Stack

* **Backend:** Go 1.26+, Gin Framework, GORM, Go-Redis, Segmentio Kafka-Go, GenAI SDK.
* **Frontend:** React 18, Vite, Tailwind CSS v4 (Dark Mode Financial Dashboard).
* **Databases & Event Broker:** PostgreSQL 15, MongoDB 6, Redis Alpine, Apache Kafka 7.7 (KRaft).
* **DevOps & Cloud:** Docker, Multi-Stage Dockerfiles, Docker Compose, Kubernetes Manifests (`/k8s` for AWS EKS).

---

## 🚀 Quick Start (Single Command Deployment)

### Prerequisites

* [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running.
* A valid `GEMINI_API_KEY` from Google AI Studio.

### Step-by-Step Instructions

1. **Clone the repository:**
```bash
git clone [https://github.com/your-username/gorisk-ai-core.git](https://github.com/your-username/gorisk-ai-core.git)
cd gorisk-ai-core

```


2. **Set your Gemini API Key in your terminal:**
* **PowerShell (Windows):**
```powershell
$env:GEMINI_API_KEY="your_actual_api_key_here"

```


* **Bash / Linux / macOS:**
```bash
export GEMINI_API_KEY="your_actual_api_key_here"

```




3. **Spin up the entire infrastructure:**
```bash
docker-compose up -d --build

```


4. **Initialize Kafka Topic (First time run only):**
```bash
docker exec gorisk-kafka kafka-topics --create --topic risk_evaluations_events --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

```


5. **Access the Applications:**
* **React Financial Dashboard:** [http://localhost:5173](http://localhost:5173)
* **Go Backend Health Check:** [http://localhost:8080/health](http://localhost:8080/health)



---

## 🧪 Running Unit Tests

To run the deterministic business rule and structure tests locally:

```bash
cd gorisk-ai-core
go test -v ./internal/services/...

```

---

## 📂 Project Structure

```text
.
├── docker-compose.yml           # Master Infrastructure Orchestrator
├── k8s/                         # AWS EKS Kubernetes Manifests
├── gorisk-ai-core/              # Golang Microservice (Hexagonal Architecture)
│   ├── Dockerfile               # Multi-stage Go build
│   ├── cmd/main.go              # Server Entrypoint
│   ├── internal/                # Private domain logic
│   └── pkg/database/            # Database drivers
└── frontend-dashboard/          # React.js SPA
    ├── Dockerfile               # Nginx Static Server build
    └── src/                     # Dashboard Components

```

---

## 👨‍💻 Author

**Jesus Castillo** - *Backend & Cloud Systems Engineer*
