# BlightSanest — Stable Insights CLI

A CLI tool to fetch and analyze financial assets, identifying outliers using configurable financial algorithms.

> **Disclaimer:** BlightSanest is not a financial advisor. It doesn't give you answers — it gives you a reliable filter and an educated guess. If you prefer a polished GUI over raw data, this app is probably not for you.

---

## How It Works

BlightSanest uses a **publisher/subscriber architecture**. The server fetches raw data from third-party APIs and publishes it to connected clients. Multiple terminal instances can then run different analyses on the same data simultaneously — without interfering with each other.

---

## Documentation

| File | Description |
|------|-------------|
| [Getting Started](docs/GETTING_STARTED.md) | Installation, environment variables, and quick start |
| [Commands](docs/COMMANDS.md) | Server, client, and search commands |
| [Crypto Commands](docs/CRYPTO.md) | Crypto-specific commands and analysis |
| [Contributing](docs/CONTRIBUTING.md) | How to contribute |

---

## Available Assets

| Asset | Source | Status |
|-------|--------|--------|
| Cryptocurrencies | [CoinGecko API](https://www.coingecko.com/en/api) | ✅ Available |
| Stock Market | — | 🔜 Planned for V2 |
