# BlightSanest - Stable Insights CLI

## What is BlightSanest?
BlightSanest is a CLI tool that allows users to fetch finance assets and run operations on them. It uses the publisher/subscriber architecture to separate data and functionality by fetching the raw data from the server and publishing it to the clients. This way you can run various operations on any finance asset simultaniously from multiple terminals.

## Available Finance Assets:

### Crypto currencies:
        **<ins>Available Timeframes:</ins> __1h, 24h, 7d, 30d, 200d, 1y__**
        From the [CoinGecko API](https://www.coingecko.com/en/api) BlightSanest Server can fetch the coins with related market data from the API [endpoint](https://docs.coingecko.com/reference/coins-markets) with the server command "fetch" with the arguments "crypto" and one or multiple timeframes (1h, 24h, 7d, 30d, 200d, 1y)
	<sub>server examples</sub>
	> `fetch crypto 1h 24h`
	> `fetch crypto 1h`

        After fetching the crypto data from the server, you also need to get it from the client. BlightSanest does not perform initial calls to any APIs on the server neither on the client not to produce undesired results and not to be a burden on the API.
	<sub>client examples</sub>
	> `get crypto`        <sup>gets all crypto lists from the publisher</sup>
	> `get crypto 1h`     <sup>gets a specific crypto list from the publisher if exists</sup>
	> `get cryoto 1h 24h` <sup>gets a specific crypto list from the publisher if exists</sup>

        From the clients that has the crypto list/lists, you can perform these operations:
	<sub><ins>**After each operation if you want to update the crypto list on the client enter the command and the argument `mutate crypto` on the client terminal. Otherwise the operations will not affect the crypto list and you will run the next operation on the original list. __Mutating the list on a client will not affect the lists on the other clients nor the lists on the server__**</ins></sub>
	    1. You can see the biggest risers/fallers by sorting the coins in ascending/descending order
	        > `rank crypto asc`
		> `rank crypto desc`
	    2. You can get the coins between certain market cap ranks and filter out the coins with low liquidity
	        > `group crypto liquidity <min_market_rank int> <max_market_rank int> <min_volume float64>`
	    3. You can identify scarce coins and find undervalued gems near their lows
	        > `group crypto scarcity <min_circulating_supply float64> <max_ath_change_percentage float64>`
	    4. You can filter the coins between a total volume range
	        > `filter crypto total_volume <min_volume float64> <max_volume float64>`
	    5. You can filter the coins between a market cap range
	        > `filter crypto market_cap <min_market_cap float64> <max_market_cap float64>`
	    6. You can filter the coins between a price change percentage rate
	        > `filter crypto price_change_percentage <min_change float64> <max_change float64> <timeframe>`
	    7. You can filter the coins by their volatility by their swing rate (rate = highest_price_24h / lowest_price_24h)
	        > `filter crypto volatile <min_rate float64> <max_rate float64>`
	    8. You can filter the high risk coins by their ath change percentages and total volumes.
	        > `filter crypto high_risk <max_ath_change_percentage float64> <max_total_volume float64>`
	    9. You can filter the low risk coins by their market cap ranks and price change percentages
	        > `filter crypto low_risk <max_market_cap_rank int> <max_price_change_percentage> <timeframe>`
	    10. You can search a coin by its name
	        > `find crypto name <coin_name string>`
	    11. You can search for the coins with a new high price.
	        > `find crypto new_high_price`
	    12. You can search for the coins with a new low price.
	        > `find crypto new_low_price`
	    **To use the features 11 and 12 you need one additional coin list, the process will happen automatically but if you only have 1 coin list on the server this feature will not work. __The server caches the same timeframe coins on hourly base.__ You can either fetch 2 lists from the server with different timeframes -this will cause the server to fetch a new list from the API, or you can wait an hour and fetch another list. Otherwise the server will not publish a new list. Removing old lists from the cache happens every three hours.**
	    13. You can search for the coins with a high price spike
	        > `find crypto high_price_spike <min_prace_change_percentage float64> <timeframe>`
	    14. You can search for the coins with potential rallies
	        > `find crypto potential_rally <max_ath_change_percentage float64>`
	    15. You can search for the coins with possible token unlocks or inflation risks
	        supply_value = current_price * circulating_supply
		ignored_coins = write the names of the coins you want to ignore with a space between the names
	        > `find crypto coin_inflation <min_market_cap_rank int> <min_supply_value float64> <ignored_coins>`
	    16. You can calculate the daily range (volatility) of the coins in a range
	        volatility = (high_24 - low_24) / current_price
	        > `calc crypto volatility <min_volatility float64> <max_volatility float64>
	    17. You can calculate the coins' growth potentials in a range with minimum growth potential and a maximum market cap rank
	        growth_potential = (ATH - current_price) / current_price * 100
		> `calc crypto growth_potential <min_potential float64> <max_market_cap_rank int>`
	    18. You can calculate the coin's liquidities and set a min liquidity value
	          liquidity = total_volume / market_cap
	        > `calc crypto liquidity <min_liquidity float64>`
	    19. You can calculate the daily coin trend strength index and get the trending coins
	        **make sure are operating on a list that has 24h timeframe**
		> `calc crypto trend_strength`

-------------------------------------Initial Relase Version-----------------------------------------------------