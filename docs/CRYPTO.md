# Crypto Commands

Cryptocurrency data is sourced from the [CoinGecko API](https://www.coingecko.com/en/api). You will need a CoinGecko API key set as `COIN_GECKO_KEY` in your `.env`.

---

## Server Commands

### `fetch crypto`

Fetches cryptocurrency market data and publishes it to connected clients.

**Without query parameters:**

```
fetch crypto
```

**With query parameters:**

Parameter order: `ids > names > symbols > include_tokens > category > order > per_page > page > sparkline > price_change_percentage > precision`

Use `-` to skip a parameter.

```
fetch crypto - Bitcoin
# skips ids, fetches by name

fetch crypto - - btc
# skips ids and names, fetches by symbol

fetch crypto - - - top - market_cap_asc - - - 1h,24h,7d
# skips ids, names, symbols, per_page, page, sparkline, and precision
```

For full parameter details see the [CoinGecko /coins/markets endpoint docs](https://docs.coingecko.com/reference/coins-markets).

---

### `get crypto`

Retrieves an existing crypto list from the database and publishes it to connected clients.

```
get crypto <list_id>
```

---

### `save crypto`

Saves a crypto list from the server cache to the database. If no custom ID is provided, the existing cache ID is used.

```
save crypto <cache_id> [custom_id]
```

---

### `delete crypto`

Deletes a crypto list from the database. Lists persist until explicitly deleted.

```
delete crypto <list_id>
```

---

## Client Commands

### `fetch crypto`

Gets a specific crypto list from the server publisher.

```
fetch crypto <list_id>
```

---

### `get crypto`

Gets a specific crypto list from a client publisher — i.e. one shared by another client via `save`.

```
get crypto <list_id>
```

---

### `save crypto`

Publishes the current client list to other clients waiting with `get`. Also adds the list to your local cache with a new ID so you can `switch` back to it later.

```
save crypto <list_id>
```

---

### `list crypto`

Prints the ID of the current client list and all lists stored in the local cache.

```
list crypto
```

---

### `set crypto`

Sets a client preference field that controls the behavior of analysis commands.

```
set crypto <field_name> <value>
```

**Available fields:**

| Field | Field |
|-------|-------|
| `current_order` | `current_sorting_field` |
| `client_timeframes` | `current_timeframe` |
| `current_max_rank` | `current_min_rank` |
| `current_max_volume` | `current_min_volume` |
| `current_min_circulating_supply` | `current_max_ath_change_percentage` |
| `current_min_market_cap` | `current_max_market_cap` |
| `current_min_price_change_percentage` | `current_max_price_change_percentage` |
| `current_min_swing_score` | `current_max_swing_score` |
| `current_ignored_coins` | `current_min_supply` |
| `current_min_volatility` | `current_max_volatility` |
| `current_min_growth_potential` | `current_min_liquidity` |

---

### `database crypto`

Performs a database operation on the current crypto list.

```
database crypto <CREATE|READ|UPDATE|DELETE>
```

---

## Analysis Commands

All operations update the current client list with their result. Use `save` before an operation if you want to preserve the previous state.

---

### `rank` — Sort Coins

Sort the current list by any field to surface top risers, fallers, or other rankings.

```
rank crypto <asc|desc> <field>
```

**Available fields:** `current_price`, `market_cap`, `market_cap_rank`, `market_cap_change_percentage`, `total_volume`, `high_24h`, `low_24h`, `ath`, `price_change_percentage`, `ath_change_percentage`, `max_supply`, `circulating_supply`

**Examples:**

```
rank crypto asc  current_price
rank crypto desc market_cap_rank
rank crypto asc  price_change_percentage
```

---

### `group` — Cluster Coins by Criteria

#### `group crypto liquidity`

Returns coins within a market cap rank range that meet a minimum volume threshold — useful for finding liquid assets.

```
group crypto liquidity <min_rank> <max_rank> <min_volume>
```

**Formula:** `MarketCapRank >= min_rank && MarketCapRank <= max_rank && TotalVolume >= min_volume`

---

#### `group crypto scarcity`

Identifies scarce coins that are still far from their all-time highs — potential undervalued gems.

```
group crypto scarcity <min_circulating_supply> <max_ath_change_percentage>
```

**Formula:** `(CirculatingSupply / MaxSupply) >= min_circulating_supply && AthChangePercentage <= max_ath_change_percentage`

---

### `filter` — Narrow Down Coins

#### By Total Volume

```
filter crypto total_volume <min> <max>
```

#### By Market Cap

```
filter crypto market_cap <min> <max>
```

#### By Price Change Percentage

```
filter crypto price_change_percentage <min> <max> <timeframe>
```

> Mind the `current_timeframe` set in your client state.

#### By Swing Rate

```
filter crypto volatile <min_rate> <max_rate>
```

**Formula:** `rate = High24h / Low24h`

#### High Risk Coins

Flags coins with a large drop from ATH and low volume — high speculation, low liquidity.

```
filter crypto high_risk <max_ath_change_percentage> <max_total_volume>
```

#### Low Risk Coins

Flags coins with a strong market position and stable price movement.

```
filter crypto low_risk <max_market_cap_rank> <max_price_change_percentage> <timeframe>
```

---

### `find` — Search for Coins

#### By Name

```
find crypto name <coin_name>
```

#### New High / New Low Price

Compares the current list against a second cached list to find coins hitting new highs or lows.

```
find crypto new_high_price
find crypto new_low_price
```

> **Requirements:** Two crypto lists must exist in the server cache. Either fetch two lists with different timeframes, or wait an hour — the server caches lists hourly and clears old ones every 3 hours. This feature will not work with only one cached list.

#### High Price Spike

Finds coins with a significant price increase over a given timeframe.

```
find crypto high_price_spike <min_price_change_percentage> <timeframe>
```

#### Potential Rally

Finds coins that are far from their ATH and may have room to grow.

```
find crypto potential_rally <max_ath_change_percentage>
```

#### Coin Inflation / Token Unlock Risk

Identifies coins with high circulating supply values that may face inflation or unlock pressure. Use `ignored_coins` to exclude known stablecoins or irrelevant assets (space-separated names).

```
find crypto coin_inflation <max_market_cap_rank> <min_supply_value> [ignored_coins]
```

**Formula:** `supply_value = CurrentPrice * CirculatingSupply`

---

### `calc` — Calculate Metrics

#### Volatility (Daily Range)

```
calc crypto volatility <min> <max>
```

**Formula:** `volatility = (High24h - Low24h) / CurrentPrice`

#### Growth Potential

```
calc crypto growth_potential <min_potential> <max_market_cap_rank>
```

**Formula:** `growth_potential = (ATH - CurrentPrice) / CurrentPrice * 100`

#### Liquidity Score

```
calc crypto liquidity <min_liquidity>
```

**Formula:** `liquidity = TotalVolume / MarketCap`

#### Trend Strength Index

Calculates the daily trend strength and returns trending coins.

```
calc crypto trend_strength
```

> **Requirement:** The current list must have been fetched with the `24h` timeframe.
