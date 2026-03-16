# Commands

## Server Commands

Run from the `go run ./cmd/server` terminal.

---

### `fetch`

Fetches market data for a given asset type from its source API and publishes it to connected clients.

```
fetch <asset>
fetch <asset> [parameters]
```

Use `-` to skip a query parameter. Parameter order varies by asset — see the relevant asset doc for details.

---

### `get`

Retrieves an existing asset list from the database and publishes it to connected clients.

```
get <asset> <list_id>
```

---

### `save`

Saves an asset list from the server cache to the database. If no custom ID is provided, the existing cache ID is used.

```
save <asset> <cache_id> [custom_id]
```

---

### `delete`

Deletes an asset list from the database. Lists persist until explicitly deleted.

```
delete <asset> <list_id>
```

---

## Client Commands

Run from the `go run ./cmd/client` terminal.

> **Note:** Every analysis operation automatically updates the current client list with the result. Use `save` before an operation to snapshot the list first.

---

### `fetch`

Gets a specific asset list from the server publisher.

```
fetch <asset> <list_id>
```

---

### `get`

Gets a specific asset list from a client publisher — i.e. one shared by another client via `save`.

```
get <asset> <list_id>
```

---

### `save`

Publishes the current client list to other clients waiting with `get`. Also adds the list to your local cache with a new ID so you can `switch` back to it later.

```
save <asset> <list_id>
```

---

### `list`

Prints the ID of the current client list and all lists stored in the local cache.

```
list <asset>
```

---

### `set`

Sets a client preference field to a given value. Preferences control the behavior of analysis commands and vary by asset type.

```
set <asset> <field_name> <value>
```

---

### `database`

Performs a database operation on the current asset list.

```
database <asset> <CREATE|READ|UPDATE|DELETE>
```

---

## Search Commands

Run from the `go run ./cmd/search` terminal.

BlightSanest includes an AI-powered semantic search engine built on [Sentence Transformers](https://www.sbert.net/) (`all-MiniLM-L6-v2`), served via FastAPI. Set `SEMANTIC_API_URL` in your `.env` to match the service URL in `docker-compose.yml`.

> **Note:** Embeddings must be generated before running a search.

---

### `embeddings`

Generates and stores vector embeddings for the current asset list.

```
embeddings <asset>
```

---

### `search`

Searches across the embedded data using natural language.

```
search <asset> <query>
```

---

## Asset Docs

For asset-specific commands, query parameters, and analysis operations see:

| Asset | Doc |
|-------|-----|
| Cryptocurrencies | [CRYPTO.md](CRYPTO.md) |
