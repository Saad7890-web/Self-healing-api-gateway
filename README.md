# 🚀 Self-Healing API Gateway in Go

A **lightweight, self-healing API Gateway** built from scratch in **Go**.  
It automatically detects failing services, reroutes traffic, and brings services back gradually, all while exposing metrics and admin APIs. Think of it as a **mini Istio / Envoy**, written entirely from scratch.

---

## 🧠 What It Is

This project is **not your usual CRUD app** or URL shortener. It’s **infrastructure software**, designed to handle real-world traffic and failures.  

Key highlights:  
- Detect failing backend services automatically  
- Route traffic away from unhealthy nodes  
- Gradually reintegrate healthy services  
- Expose metrics & admin APIs for observability  

---

## 🤯 Why This Project Stands Out

Most developers build simple apps. Few venture into **systems-level, production-grade infrastructure**. This project demonstrates:

- **Systems thinking** – handling service failures and recovery automatically  
- **Concurrency mastery** – goroutines, atomic operations, and channels  
- **Networking & reverse proxy skills** – routing, load balancing, circuit breaking  
- **Go expertise** – idiomatic, production-ready code  

> This is the kind of project that **interviewers LOVE**.

---

## 🔥 Core Features

### 1️⃣ Smart Reverse Proxy
- Accepts incoming HTTP requests and forwards them to backend services  
- Load balances using:
  - **Round-Robin**
  - **Least Connections**
  - **Weighted Routing**  
- **Built with:** `net/http`, `httputil.ReverseProxy`

### 2️⃣ Health Checking Engine (Self-Healing Core)
- Periodically pings backend services  
- Marks services as:
  - `HEALTHY`
  - `DEGRADED`
  - `DOWN`  
- Automatically removes unhealthy nodes from routing  
- **Built with:** Goroutines, `time.Ticker`, atomic state updates

### 3️⃣ Circuit Breaker (Netflix-Style)
- Tracks:
  - Failure rate
  - Timeout rate  
- Opens the circuit when failures exceed thresholds  
- Prevents cascading failures across services  
- **Implementation includes:** Sliding window counters, state machine logic

### 4️⃣ Rate Limiting (Per Client / Per Route)
- Protects backend services with **Token Bucket** or **Leaky Bucket** algorithms  
- Handles **per-client** and **per-route** limits  
- **Built with:** `sync.Map` and time-based refill logic

### 5️⃣ Distributed Tracing IDs
- Injects `X-Request-ID` headers  
- Logs request flows across services for **observability**  
- Enables tracing and debugging in complex distributed systems

---

## 🛠 Tech Stack
- **Language:** Go  
- **Concurrency:** Goroutines, channels, atomic operations  
- **HTTP:** `net/http`, `httputil.ReverseProxy`  
- **Observability:** Custom request IDs, logging, metrics endpoints  

---

## ⚡ Why You Should Care
- Shows **real-world production skills**  
- Demonstrates **systems design thinking**  
- Great **resume booster** for backend, Go, and infrastructure roles  
- Perfect **interview talking point** for distributed systems, networking, and resiliency  

---

## 📈 Next Steps / Improvements
- Add **dynamic service discovery** (e.g., via Consul or etcd)  
- Integrate **Prometheus metrics** for observability  
- Add **configurable routing policies** and **failover strategies**  
- Implement **advanced load balancing** (least latency, weighted round-robin, etc.)

---

### 💡 TL;DR
> **A self-healing, production-grade API gateway in Go, showcasing systems thinking, concurrency, networking, and Go mastery.**
