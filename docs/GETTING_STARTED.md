# Getting Started

## Prerequisites

- [Go](https://go.dev/) — to run the server, client, and search commands
- [Docker](https://www.docker.com/) — to run RabbitMQ and PostgreSQL
- A [CoinGecko API key](https://www.coingecko.com/en/api) (free tier works)

---

## Installation

**1. Clone the repository**

```bash
git clone https://github.com/mehmetcagriekici/blightsanest.git
cd blightsanest
```

**2. Create a `.env` file** in the project root with the variables listed below.

**3. Start Docker services**

```bash
docker compose build   # only needed once
docker compose up
```

**4. Run the application** — each in its own terminal

```bash
go run ./cmd/server
go run ./cmd/client
go run ./cmd/search
```

---

## Environment Variables

| Variable | Description |
|----------|-------------|
| `COIN_GECKO_KEY` | CoinGecko API key for cryptocurrency data |
| `RABBITMQ_URL` | Connection URL to the RabbitMQ server |
| `CACHE_INTERVAL` | Time until a cache entry becomes stale and is removed |
| `SUBSCRIBER_PREFETCH` | Prefetch count for AMQP QoS |
| `DATABASE_URL` | PostgreSQL connection string |
| `SEMANTIC_API_URL` | FastAPI URL — must match the value in `docker-compose.yml` |
