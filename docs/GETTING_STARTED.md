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

**4. Build the binaries** — run once, then use the compiled executables

```bash
go build -o bin/server ./cmd/server
go build -o bin/client ./cmd/client
go build -o bin/search ./cmd/search
```

Or run directly without building:

```bash
go run ./cmd/server
go run ./cmd/client
go run ./cmd/search
```

Each program runs in its own terminal. Start the server first, then connect clients and the search engine as needed.

---

## Usage

All three programs use [Cobra](https://github.com/spf13/cobra)-based CLIs. Each command has built-in help available via the `--help` flag:

```bash
./bin/client --help
./bin/client fetch --help
./bin/client calc crypto --help
```

---

## Environment Variables

| Variable | Description |
|----------|-------------|
| `COIN_GECKO_KEY` | CoinGecko API key for cryptocurrency data |
| `RABBITMQ_URL` | Connection URL to the RabbitMQ server |
| `CACHE_INTERVAL` | Time (in hours) until a cache entry becomes stale and is removed |
| `SUBSCRIBER_PREFETCH` | Prefetch count for AMQP QoS |
| `DATABASE_URL` | PostgreSQL connection string |
| `SEMANTIC_API_URL` | FastAPI URL — must match the value in `docker-compose.yml` |
