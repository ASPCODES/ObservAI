# ObservAI

**One gateway to monitor, manage and control all your API traffic.**

 ObservAI is an API Gateway & Observability Platform designed for developers and modern engineering teams. It provides a centralized layer to authenticate, monitor, analyze and control API traffic while offering real-time visibility into system performance.

Instead of connecting applications directly to services, requests pass through ObservAI where they can be authenticated, rate-limited, logged and analyzed before reaching their destination.

---

## Why ObservAI?

Modern applications rely on dozens of APIs:

* Internal microservices
* Payment providers
* Authentication services
* Third-party integrations
* AI providers

As API traffic grows, teams often struggle to answer basic questions:

* How many requests are being processed?
* Which endpoints are receiving the most traffic?
* What is the average latency?
* Which API keys are generating the highest load?
* Why are requests failing?
* How can traffic be controlled and monitored in real time?

 ObservAI solves these problems by acting as a centralized API control plane.

---

## Key Features

### API Gateway

Route all API traffic through a single gateway layer.

### API Key Management

Generate, manage and revoke API keys for projects and applications.

### Request Logging

Capture request metadata, response status codes and latency metrics.

### Real-Time Analytics

Monitor API traffic as it happens through live dashboards and WebSocket-powered updates.

### Rate Limiting

Protect backend services from abuse and excessive traffic.

### Multi-Project Support

Manage multiple applications and environments from a single account.

### Observability Dashboard

Track:

* Total Requests
* Success Rate
* Error Rate
* Average Latency
* Traffic Trends
* Top Endpoints

---

## Architecture Overview

```text
                           Client Applications
                    (Web, Mobile, Backend Services)
                                      |
                                      |
                               HTTPS Requests
                                      |
                                      |
                    +--------------------------------+
                    |        API Gateway             |
                    |         (Golang)               |
                    +--------------------------------+
                                      |
          ----------------------------------------------------------------
          |                        |                    |                 |
          |                        |                    |                 |
     Authentication          Rate Limiter         Request Logger      Metrics Engine
      & API Keys                (Redis)              (Go)              (Go)
          |                        |                    |                 |
          |                        |                    |                 |
     PostgreSQL                Redis Cache        PostgreSQL        Redis Cache
       (Prisma)                                                   (Realtime Stats)
          |                                                         |
          |                                                         |
          -----------------------------------------------------------
                                      |
                                      |
                              Event Publisher
                                      |
                                      |
                              WebSocket Hub
                                      |
                                      |
                           Next.js Dashboard
```

---

## Technology Stack

### Frontend

* Next.js
* React
* Tailwind CSS

### Backend

* Golang
* Gin/Fiber
* JWT Authentication

### Database

* PostgreSQL
* Prisma ORM

### Infrastructure

* Redis
* Docker
* Docker Compose

### Real-Time Communication

* WebSockets

---

## Core Components

### Authentication Service

Provides user authentication and access control.

### Project Management

Allows users to organize APIs and environments into projects.

### API Key Service

Secure API key generation, validation and revocation.

### Gateway Service

Processes incoming requests and applies policies such as authentication and rate limiting.

### Metrics Engine

Aggregates traffic data and generates performance insights.

### Observability Layer

Provides monitoring, analytics and traffic visibility.

---

## Dashboard

ObservAI provides a modern developer-focused dashboard for monitoring API traffic.

### Overview

* Total Requests
* Success Rate
* Average Latency
* Error Rate

### Projects

Manage multiple applications and environments.

### API Keys

Generate and control project-level API credentials.

### Requests

Inspect recent API activity and response metrics.

### Analytics

Visualize traffic patterns and performance trends.

---


## Vision

ObservAI aims to become the control plane for modern API infrastructure.

A single platform where developers can:

* Monitor API traffic
* Manage API access
* Analyze performance
* Enforce policies
* Gain real-time operational visibility

without building custom observability and gateway solutions from scratch.

---

**Built for developers who want complete visibility into their API infrastructure.**
