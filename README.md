# BlightSanest - Stable Insights CLI

## What is BlightSanest?
BlightSanest is a CLI tool that allows users to fetch finance assets and analyze them, finding/identifying outliers. It uses the publisher/subscriber architecture to separate data and functionality by fetching the raw data from the server and publishing it to the clients. This way you can run various operations on any finance asset simultaniously from multiple terminals.

## Dev Logs
If you face any bugs problems or something not clear, please do reach me from mehmetcagriekici@gmail.com
### Initial Relase V1
1) Faced an issue with REPL and pubsub architecture, fixed the cancellation logic.
2) Faced few bugs concerning routing. Fixed the issue with routing and binding keys.
3) Faced an issue with caching lists, and client distributions, fixed queue consumption and async data flow logic
4) Faced a bug that causes clients to exit after fetching a new list with new args from the server. Loosened strict key-id match check, eliminating the bug due to key creation process.
5) Facing an issue while updating clients, Solved creating another channel for inner client publishings.
6) Added more server query parameters for the API to improve the UX.
7) Added DLX and manual acknowledgement for crypto routing.

## How to Use:
1) Create a .env file with the necessary variables described below
2) Make sure docker is running
3) Start the rabbitmq server from your CLI using rabbit.sh file
```
./rabbit.sh start
```
3) Directly run the server and the client separately from different CLIs - or build them one by one, and run the executables from separate CLIs.
```
go run ./cmd/client
go run ./cmd/server

go build ./cmd/client
go build ./cmd/server
```

## Enviromental Variables:
```
COIN_GECKO_KEY           # coin gecko api key for crypto currencies
RABBIT_CONNECTION_STRING # url to the rabbitmq server
CACHE_INTERVAL           # time until a crypto cache entry becomes stalei and removed
SUBSCRIBER_PREFETCH      # prefetch count for amqp Qos
```

## Available Finance Assets:

### Crypto currencies:

**<ins>Available Timeframes:</ins> __1h, 24h, 7d, 30d, 200d, 1y__**

From the [CoinGecko API](https://www.coingecko.com/en/api) BlightSanest Server can fetch the coins with related market data from the API [endpoint](https://docs.coingecko.com/reference/coins-markets) with the server command "fetch" with the arguments "crypto" and one or multiple queries.

Server Examples:

1) Without any query parameters:
```
fetch crypto # fetches the data without any query parameters

```
2) Query parameters order: ids > names > symbols > include_tokens > category > order > per_page > page > sparkline > price_change_percentage > percision
Make sure to visit the API docs -endpoint- to see how query parameters work.
**This order is strict! Skip the parameter with minus sign (-) see the examples below.**
```
fetch crypto - Bitcoin                                  # omits ids and the rest of the parameters after the names.

fetch crypto - - btc                                    # omits ids, names and the rest of the parameters after the symbols.

fetch crypto - - - top - market_cap_asc - - - 1h,24h,7d # omits ids, names, symbols, category, per_page, page, sparkline, and percision.
```

After fetching the crypto data from the server, you also need to get it from the client. BlightSanest does not perform initial calls to any APIs on the server neither on the client not to produce undesired results and not to be a burden on the API.

Client Examples:

```
fetch crypto crypto_list_id # gets a specific crypto list from the server publisher if exists

get crypto crypto_list_id   # gets a specific crypto list from a client publisher if exists

save crypto crypto_list_id  # publishes a crypto list to other clients that are waiting for it with get command

list crypto                 # prints the ids of the current client list and the lists in the cache
```

From the clients that has the crypto list, you can perform these operations:

**Each operation will update the current client list on the state. You can publish the client list to other clients before the operation with <save> command. Save command will also add the client list to the client cache with a new ID, you can later <switch> to it.**

1. You can see the biggest risers/fallers by sorting the coins in ascending/descending order
Available Fields: <current_price>, <market_cap>, <market_cap_rank>, <market_cap_change_percentage>, <total_volume>, <high_24h>, <low_24h>, <ath>, <price_change_percentage (with an existing timeframe in the client state)>, <ath_change_percentage>, <max_supply>, <circulating_supply>

```
rank crypto asc current_price
    
rank crypto desc market_cap_rank

rank crypto asc price_change_percentage
```

2. You can get the coins between certain market cap ranks and filter out the coins with low liquidity

```
group crypto liquidity <min_market_rank int> <max_market_rank int> <min_volume float64>
```

3. You can identify scarce coins and find undervalued gems near their lows

```
group crypto scarcity <min_circulating_supply float64> <max_ath_change_percentage float64>
```

4. You can filter the coins between a total volume range

```
filter crypto total_volume <min_volume float64> <max_volume float64>
```

5. You can filter the coins between a market cap range

```
filter crypto market_cap <min_market_cap float64> <max_market_cap float64>
```

6. You can filter the coins between a price change percentage rate

```
filter crypto price_change_percentage <min_change float64> <max_change float64> <timeframe>
```

7. You can filter the coins by their volatility by their swing rate

```
# rate = highest_price_24h / lowest_price_24h

filter crypto volatile <min_rate float64> <max_rate float64>
```

8. You can filter the high risk coins by their ath change percentages and total volumes.

```
filter crypto high_risk <max_ath_change_percentage float64> <max_total_volume float64>
```

9. You can filter the low risk coins by their market cap ranks and price change percentages

```
filter crypto low_risk <max_market_cap_rank int> <max_price_change_percentage> <timeframe>
```

10. You can search a coin by its name

```
find crypto name <coin_name string>
```

11. You can search for the coins with a new high price.

```
find crypto new_high_price
```

12. You can search for the coins with a new low price.

```
find crypto new_low_price
```
     
**To use the features 11 and 12 you need one additional coin list, the process will happen automatically but if you only have 1 coin list on the server this feature will not work. __The server caches the same timeframe coins on hourly base.__ You can either fetch 2 lists from the server with different timeframes -this will cause the server to fetch a new list from the API, or you can wait an hour and fetch another list. Otherwise the server will not publish a new list. Removing old lists from the cache happens every three hours.**

13. You can search for the coins with a high price spike

```
find crypto high_price_spike <min_prace_change_percentage float64> <timeframe>
```

14. You can search for the coins with potential rallies

```
find crypto potential_rally <max_ath_change_percentage float64>
```

15. You can search for the coins with possible token unlocks or inflation risks

```
# supply_value = current_price * circulating_supply

# ignored_coins = write the names of the coins you want to ignore with a space between the names

find crypto coin_inflation <min_market_cap_rank int> <min_supply_value float64> <ignored_coins>
```

16. You can calculate the daily range (volatility) of the coins in a range

```
# volatility = (high_24 - low_24) / current_price  

calc crypto volatility <min_volatility float64> <max_volatility float64>
```

17. You can calculate the coins' growth potentials in a range with minimum growth potential and a maximum market cap rank

```
# growth_potential = (ATH - current_price) / current_price * 100

calc crypto growth_potential <min_potential float64> <max_market_cap_rank int>
```

18. You can calculate the coin's liquidities and set a min liquidity value

```
# liquidity = total_volume / market_cap

calc crypto liquidity <min_liquidity float64>
```

19. You can calculate the daily coin trend strength index and get the trending coins

```
# make sure are operating on a list that has 24h timeframe

calc crypto trend_strength
```

-------------------------------------Initial Relase Version-----------------------------------------------------
# Future Plans
## Stock Market Implementation