# Commands

Every command supports `--help` for usage details. For example:

```bash
./bin/server fetch --help
./bin/client calc crypto --help
```

---

## Server Commands

Run from a terminal using `go run ./cmd/server` or `./bin/server`.

---

### `fetch`

Fetches market data for a given asset type from its source API and publishes it to connected clients.

```
server fetch <asset> [parameters...]
```

Use `-` to skip a query parameter. Parameter order varies by asset — see the relevant asset doc for details.

---

### `read`

Retrieves an existing asset list from the database and publishes it to connected clients.

```
server read <asset> <list_id>
```

---

### `save`

Saves an asset list from the server cache to the database. If no custom ID is provided, the existing cache ID is used.

```
server save <asset> <cache_id> [custom_id]
```

---

### `delete`

Deletes an asset list from the database. Lists persist until explicitly deleted.

```
server delete <asset> <list_id>
```

---

## Client Commands

Run from a terminal using `go run ./cmd/client` or `./bin/client`.

> **Note:** Every analysis operation automatically updates the current client list with the result. Use `save` before an operation to snapshot the list first.

---

### `fetch`

Gets a specific asset list from the server publisher.

```
client fetch <asset> <list_id>
```

---

### `get`

Gets a specific asset list from a client publisher — i.e. one shared by another client via `save`.

```
client get <asset> <list_id>
```

---

### `save`

Publishes the current client list to other clients waiting with `get`. Also adds the list to your local cache with a new ID so you can `switch` back to it later.

```
client save <asset> <list_id>
```

---

### `list`

Prints the ID of the current client list and all lists stored in the local cache.

```
client list <asset>
```

---

### `switch`

Switches the active client list to a different one stored in the local cache.

```
client switch <asset> <cache_id>
```

---

### `set`

Sets a client preference field to a given value. Preferences control the behavior of analysis commands and vary by asset type.

```
client set <asset> <field_name> <value>
```

---

### `database`

Performs a database operation on the current asset list.

```
client database <asset> <CREATE|READ|UPDATE|DELETE>
```

---

### `rank`

Sorts the current asset list by a field.

```
client rank <asset> <asc|desc> <field>
```

---

### `group`

Clusters assets by a grouping criteria. Subcommands vary by asset — see the asset doc for available options.

```
client group <asset> <subcommand> [args...]
```

---

### `filter`

Narrows down the current list by a filtering criteria.

```
client filter <asset> <subcommand> [args...]
```

---

### `find`

Searches within the current list.

```
client find <asset> <subcommand> [args...]
```

---

### `calc`

Calculates a metric over the current list.

```
client calc <asset> <subcommand> [args...]
```

---

## Search Commands

Run from a terminal using `go run ./cmd/search` or `./bin/search`.

BlightSanest includes an AI-powered semantic search engine built on [Sentence Transformers](https://www.sbert.net/) (`all-MiniLM-L6-v2`), served via FastAPI. Set `SEMANTIC_API_URL` in your `.env` to match the service URL in `docker-compose.yml`.

> **Note:** Embeddings must be generated before running a search.

---

### `embeddings`

Generates and stores vector embeddings for the current asset list.

```
search embeddings <asset>
```

---

### `search`

Searches across the embedded data using natural language.

```
search search <asset> <query>
```

---

## Asset Docs

For asset-specific commands, query parameters, and analysis operations see:

| Asset | Doc |
|-------|-----|
| Cryptocurrencies | [CRYPTO.md](CRYPTO.md) |
