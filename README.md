# Go Microservice App

This repository contains a microservice setup written in **Go**.  
The system consists of **6 services** that communicate using **gRPC** and **RabbitMQ**, with **Caddy** as the web server.

## 🧰 Tech Stack
- Go
- Docker / Docker Swarm
- RabbitMQ
- PostgreSQL
- MongoDB (Logger)
- Mailhog (Email testing)
- gRPC
- Caddy

---

## 📦 Services

| Service | Description |
|--------|-------------|
| **Front-End Service** | User-facing interface that communicates with the Broker Service. |
| **Broker Service** | Routes incoming requests from the Front-End to the correct service. |
| **Authentication Service** | Simple authentication using PostgreSQL. |
| **Listener Service** | Listens to RabbitMQ and handles message events. |
| **Logger Service** | Stores event logs to MongoDB. |
| **Mailer Service** | Sends emails through Mailhog. |

---

## 🌐 Web Server
The app uses **Caddy** as the web server and reverse proxy.

---

## 🐳 Docker & Makefile

The repository includes:
- `Dockerfile`
- `docker-compose.yml` (rename from provided example)
- `Makefile`

These help with building and running the services.

### Running with Make

> **Windows Users**: If `make` is not installed:
```powershell
cd project

```powershell
choco install make

```powershell
mv docker-compose.example.yml docker-compose.yml'

```powershell
make up_build

🐝 Deploy Using Docker Swarm

Initialize Docker Swarm:
```powershell
docker swarm init

```powershell
docker stack deploy -c docker-stack.yml your-chosen-app-name

Check running service:
```powershell
docker stack services your-chosen-app-name

Remove service:
```powershell
docker stack rm your-chosen-app-name

