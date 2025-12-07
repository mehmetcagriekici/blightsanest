# BlightSanest - Stable Insights CLI Commands

These commnads are desinged to include various finance assets and common/special commands for users to run analyses on their financial assets.

## Server Commands:

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

## Client Commands:

### Available Fields for <set> command:
```
<current_order>                         <current_sorting_field>
<client_timeframes>                     <current_timeframe>
<current_max_rank>                      <current_min_rank>
<current_max_volume>                    <current_min_volume>
<current_min_circulating_supply>        <current_max_ath_change_percentage>
<current_min_market_cap>                <current_max_market_cap>,
<current_min_price_change_percentage>   <current_max_price_change_percentage>
<current_min_swing_score>               <current_max_swing_score>
<current_ignored_coins>                 <current_min_supply>
<current_min_volatility>                <current_max_volatility>
<current_min_growth_potential>          <current_min_liquidity>
```

```
fetch crypto crypto_list_id # gets a specific crypto list from the server publisher if exists

get crypto crypto_list_id   # gets a specific crypto list from a client publisher if exists

save crypto crypto_list_id  # publishes a crypto list to other clients that are waiting for it with get command

list crypto                 # prints the ids of the current client list and the lists in the cache

set crypto field_name value # sets the client preference a crypto state fields to the value
```
## Client Crypto Commands

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

7. You can filter the coins by their volatility using their swing rate

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

find crypto coin_inflation <max_market_cap_rank int> <min_supply_value float64> <ignored_coins>
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