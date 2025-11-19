package crypto

import(
        "fmt"
)

func PrintCryptoOptions() {
	fmt.Println("Available Crypto Features:")
	fmt.Println("1) Rankings and Leaderboards:")
	fmt.Println("1.1) To see the biggest risers: * rankcoins asc <timeframe>")
	fmt.Println("1.2) To see the biggest fallers: * rankcoins desc <timeframe>")
	fmt.Println("1.3) To see the coins between market cap ranks and filter out the low-liquidity ones: * groupcoins <min rank> <max rank> <min volume>")
	fmt.Println("1.4) To identify scarce assets and to find undervalued gems near their lows (circulating_supply / max_supply = scarcity_score): * scarcecoins <circulating supply> <max supply> <ath change percentage>")
	fmt.Println("2) Filtering and Search Options:")
	fmt.Println("2.1) To filter the coins by total volume: * filtercoins total_volume <min volume> <max volume>")
	fmt.Println("2.2) To filter the coins by market cap: * filtercoins market_cap <min cap> <max cap>")
	fmt.Println("2.3) To filter the coins by their price change percentage: * filtercoins price_change_percentage <min change> <max change> <timeframe>")
	fmt.Println("2.4) To filter the volatile coins by their swing rate (high_24h / low_24h = swing_rate): * filtercoins volatile <min rate> <max rate>")
        fmt.Println("2.5) To search a coin by its name: * searchcoins name <name>")
	fmt.Println("2.6) To get high risk coins (overhyped - near zero ath change percentage - or illiquid - low total volume -): * searchcoins high_risk <max ath change> <max volume>")
	fmt.Println("2.7) To get low risk coins (high market cap rank and stable price change - percentage-): * searchcoins low_risk <max market rank> <max price change> <timeframe>")
	fmt.Println("3) Alerts and Notifications:")
	fmt.Println("3.1) To get the coins with a new high price: * alertcoins new_high")
	fmt.Println("3.2) To get the coins with a new low price: * alertcoins new_low")
	fmt.Println("3.3) To get the coins with a high price soike: * alertcoins high_spike <min change> <timeframe>")
	fmt.Println("3.4) To get the coins with potential rallies (low ath change percentage): * alertcoins potential_rally <max ath change>")
	fmt.Println("3.5) To get the coins with with possible token unlocks or inflation risks: (supply_value = circulating_supply * current_price): * alertcoins high_circ_supply <min market rank> <supply value> <ignored coin 1> <ignored coin 2> <ignored coin 3>...")
	fmt.Println("4) Analytics and Insights")
	fmt.Println("4.1) To calculate daily range (volatility = (high_24 - low_24) / current_price) - only works with 24H timeframe -: * calccoins daily_range <max volatility>")
	fmt.Println("4.2) To calculate the coin liquidity (turnover_ratio = total_volume / market_cap): * calccoins liquidity <min liquidity>")
	fmt.Println("4.3) To calculate the trend strength index - only available for 24H timeframe - (price change and market cap change > 0): * calccoins trend_strength <timeframe>")
}

func PrintAvailableCryptoTimeframes() {
        fmt.Println("Available timeframes:")
	fmt.Println("One hour: 1h")
	fmt.Println("One day: 24h")
	fmt.Println("One week: 7d")
	fmt.Println("One month: 30d")
	fmt.Println("Two hundred days: 200d")
	fmt.Println("One year: 1d")
}
