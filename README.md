# BlightSanest - Stable Insights CLI

## What is BlightSanest?

BlightSanest is a CLI tool that allows users to fetch finance assets and analyze them, finding/identifying outliers.

## Architecture

It uses the publisher/subscriber architecture to separate data and functionality by fetching the raw data from the server and publishing it to the clients. This way you can run various operations on any finance asset simultaniously from multiple terminals.

## Quick Start:
1) Clone this repository
2) Create a .env file with the necessary variables described below
3) Make sure docker is running
4) Start the rabbitmq server from your CLI using rabbit.sh file
```
./rabbit.sh start
```
5) Start the postgresql using postgres.sh file - still testing the possibe permission issues, if you face any problem with the database please reach me.
```
./postgres.sh start
```
> or start these services manually with the correct credentials
6) Directly run the server and the client separately from different CLIs - or build them one by one, and run the executables from separate CLIs.
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
DB_URL                   # postgres connection string
```

## Motivation

I did not have time to learn finance but wanted to invest my money as widely as possible.

**Honest Disclaimer before starting:**
> !!!I am not a financial expert!!! but you can implement your own knowledge.
> If you don't like figuring things out from "raw" data, BlightSanest might not be for you. I don't want to discourage you from using my application but this is the truth. BlightSanest is a CLI app and offers genuine help for your finance moves with wide perspective. It doesn't give you answers but offers you a reliable filter. 

- Why do I need a CLI app to get financial data?
The answer is simple, because it's more flexible this way. Enough being limited by fancy looking GUIs. It's time to embrace the real data flow from the wide arms of CLIs.

- How can I trust in a system in which everything looks like a gamble without any prior knowledge and experience?
Blightsanest doesn't expect you to trust into any system. It includes known financial algorithms for you to filter out data gradually and see if the assets are reliable. It's gives you an educated guess.

- Do I need to keep this app 24/7 open?
The app allows you to send a request to the API(s) hourly with the same queries, and it has a configurealbe reaping interval value. **It also depends with your API subscription(This is based on free/demo versions)**
On current version, there is no way to save your data and exit a session - there is no storage - because I want to make the core version as reliable as possible. But before the release of version 2 I will implement a database so you can save your data and come back later.

- How do I choose where to put my money if all I get is filtered lists?
Simply trusts the lists but don't bet on one horse. BlightSanest designed to offer accurate financial algorithms on various assets - with version 1 we only have crypto - and the brilliant thing is the app is about the architecture rather than explicit answers. You can easily implement your own algorithms and **contribute** to the repo or not. What I promise is providing a stable app with default options to give you stable insights.

## Usage
To see commands in detail: [commands](https://github.com/mehmetcagriekici/blightsanest/blob/main/COMMANDS.md)

## Features

The answer to "Why does this app exist?" question is that I wanted to create a communicative way to observe data, in this case financial assets.
1) Hence the Main Feature -> Ability to see various outcomes on different terminal without mutating the other one.

2) Caching (Not to be a burden to 3rd part services) -> I used a structure that serves two different purposes, concerning caching the data. On server side, I didn't want the make the same request twice to the API(s), without blocking the user's ability to make requests from the server. And for the client side, I wanted to give the user the flexibility to work on the data from different filters, using the caching as a saving system.

3) Flexibility (Easy to debug/update and expand) -> I've built this app for me and the users to add new code without delving into the app too much. 

## Available Finance Assets:

### Crypto currencies:

From the [CoinGecko API](https://www.coingecko.com/en/api) BlightSanest Server can fetch the coins with related market data from the API [endpoint](https://docs.coingecko.com/reference/coins-markets) with the server command "fetch" with the arguments "crypto" and one or multiple queries.

After fetching the crypto data from the server, you also need to get it from the client. BlightSanest does not perform initial calls to any APIs on the server neither on the client not to produce undesired results and not to be a burden on the API.

```
CurrentList                     # Updated after each operation, can be changed with one of the existing ones in the client cache - see <list> command - with <switch> command

CurrentListID                   # Same as CurrentList - base id -> created hour in unix - rest -> latest operation preference values

CurrentOrder                    # Used with rank operation, updated after the operation if one passed.

CurrentSortingField             # Used with rank operation, updated after the operation if one passed.

ClientTimeframes                # Existing timeframes for coin price_change_percentage values. (ex: price_change_percentage_1h)

CurrentTimeframe                # One of the client timeframes

CurrentMaxRank                  # Maximum coin market cap rank preference

CurrentMinRank                  # Minimum coin market cap rank preference

CurrentMinVolume                # Minimum coin total volume preference

CurrentMaxVolume                # Maximum coin total volume preference

CurrentMinCirculatingSupply     # Minimum coin circulating supply preference

CurrentMaxAthChangePercentage   # Maximum coin all time high price change percentage preference

CurrentMinMarketCap             # Minimum coin market cap preference

CurrentMaxMarketCap             # Maximum coin market cap preference

CurrentMinPriceChangePercentage # Minimum coin price change percentage preference (mind CurrentTimeframe)

CurrentMaxPriceChangePercentage # Maximum coin price change percentage preference (mind CurrentTimeframe)

CurrentMinSwingScore            # Minimum coin swing score preference (swing score = coin.High24h / coin.Low24h)

CurrentMaxSwingScore            # Maximum coin swing score preference (swing score = coin.High24h / coin.Low24h)

CurrentIgnoredCoins             # Names of the coins that will be ignored while looking for high circulating supplies

CurrentMinSupply                # Minimum coin supply preference (coin supply = coin.CurrentPrice * coin.CirculatingSupply)

CurrentMinVolatility            # Minimum coin volatility score preference (volatility = (coin.High24h - coin.Low24h) / coin.CurrentPrice)

CurrentMaxVolatility            # Maximum coin volatility score preference (volatility = (coin.High24h - coin.Low24h) / coin.CurrentPrice)

CurrentMinGrowthPotential       # Minimum coin growth potential score preference (growth potential = (coin.ATH - coin.CurrentPrice) / coin.CurrentPrice * 100)

CurrentMinLiquidity             # Minimum coin liquidity score preference (liquidity = coin.TotalVolume / coin.MarketCap)

##  To Group Coins with High Liquidity:
### coin.MarketCapRank >= CurrentMinRank && coin.MarketCapRank <= CurrentMaxRank && coin.TotalVolume >= CurrentMinVolume

##  To Group Scarce Coins
### (coin.CirculatingSupply / coin.MaxSupply) >= coin.MinCirculatingSupply && coin.AthChangePercentage <= MaxAthChangePercentage

##  To Filter Coins by Their Total Volumes
### coin.TotalVolume >= CurrentMinVolume && coin.TotalVolume <= CurrentMaxVolume

##  To Filter Coins by Their Market Cap Values
### coin.MarketCap >= CurrentMinMarketCap && coin.MarketCap <= CurrentMaxMarketCap

##  To Filter Coins by Their Price Change Percentages for a Certain Timeframe
### coin.PriceChangePercentage<CurrentTimeframe> >= CurrentMinPriceChangePercentage && coin.PriceChangePercentage<CurrentTimeframe> <= CurrentMaxPriceChangePercentage

##  To Filter Coins by Their Volatility
### (coin.High24h / coin.Low24h) >= CurrentMinSwingScore && (coin.High24h / coin.Low24h) <= CurrentMaxSwingScore

##  To Filter Coins to Flag High Risk Coins
### coin.AthChangePercentage <= CurrentMaxAthChangePercentage && coin.TotalVolume <= CurrentMaxVolume

##  To Filter Coins to Flag Low Risk Coins
### coin.MarketCapRank <= CurrentMaxRank && coin.PriceChangePercentage<CurrentTimeframe> <= CurrentMaxPriceChangePercentage

##  To Search for the Coins with a High Price Spike
### coin.PriceChangePercentage<CurrentTimeframe> >= CurrentMinPriceChangePercentage

##  To Search for the Coins with Potential Rallies
### coin.AthChangePercentage <= CurrentMaxAthChangePercentage

##  To Search for the Coins with Possible Token Unlocks
### coin.MarketCapRank <= CurrentMaxRank && (coin.CurrentPrice * coin.CirculatingSupply) >= CurrentMinSupply && !slices.Contains(CurrentIgnoredCoins, coin.Name)

##  To Calculate the Daily Range of the Coins with Client Preferences
### (coin.High24h - coin.Low24h) / coin.CurrentPrice <= CurrentMaxVolatility && (coin.High24h - coin.Low24h) / coin.CurrentPrice >= CurrentMinVolatility

##  To Calculate the Growth Potential of the Coins with Client Preferences
### (coin.ATH - coin.CurrentPrice) / coin.CurrentPrice * 100 >= CurrentMinGrowthPotential && coin.MarketCapRank <= CurrentMaxRank

##  To Calculate the Liquidities of the Coins with Client Preferencs
### coin.TotalVolume / coin.MarketCap >= CurrentMinLiquidity
```
## Contributing
To see the changelog: [changelog](https://github.com/mehmetcagriekici/blightsanest/blob/main/CHANGELOG.md)
Please fork the repository and open a pull request to the `main` branch...
1) Clone the repository
```
# you can see the clonning options from the green <> code button on the main screen
git clone https://github.com/mehmetcagriekici/blightsanest.git
```
2) Run the tests
```
go test ./...
```
3) Submit a pull request

# Future Plans

## V1.1-1.9 - making the app more flexible
1) Implementing a database. -tests-
2) Implementing a search functionality with RAG. -implementatiom + tests-

## V2 - expanding the app
1) Stock Market Implementation

# License - Copied from [mit license](http://opensource.org/license/mit)
Copyright 2025 Mehmet Cagri Ekici

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
