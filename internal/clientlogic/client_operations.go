package clientlogic

// main operations
const (
	CLIENT_SWITCH string = "switch"
	CLIENT_SAVE   string = "save"
	CLIENT_LIST   string = "list"
	CLIENT_GET    string = "get"
	CLIENT_RANK   string = "rank"
	CLIENT_GROUP  string = "group"
	CLIENT_FILTER string = "filter"
	CLIENT_FIND   string = "find"
	CLIENT_CALC   string = "calc"
	CLIENT_FETCH  string = "fetch"
)

// crypto features (sub operations)
const (
        CRYPTO_GROUP_LIQUIDITY                string = "liquidity"
	CRYPTO_GROUP_SCARCITY                 string = "scarcity"
	CRYPTO_FILTER_TOTAL_VOLUME            string = "total_volume"
	CRYPTO_FILTER_MARKET_CAP              string = "market_cap"
	CRYPTO_FILTER_PRICE_CHANGE_PERCENTAGE string = "price_change_percentage"
	CRYPTO_FILTER_VOLATILE                string = "volatile"
	CRYPTO_FILTER_HIGH_RISK               string = "high_risk"
	CRYPTO_FILTER_LOW_RISK                string = "low_risk"
	CRYPTO_FIND_NAME                      string = "name"
	CRYPTO_FIND_NEW_HIGH_PRICE            string = "new_high_price"
	CRYPTO_FIND_NEW_LOW_PRICE             string = "new_low_price"
	CRYPTO_FIND_HIGH_PRICE_SPIKE          string = "high_price_spike"
	CRYPTO_FIND_POTENTIAL_RALLY           string = "potential_rally"
	CRYPTO_FIND_COIN_INFLATION            string = "coin_inflation"
	CRYPTO_CALC_VOLATILITY                string = "volatility"
	CRYPTO_CALC_GROWTH_POTENTIAL          string = "growth_potential"
	CRYPTO_CALC_LIQUIDITY                 string = "liquidity"
	CRYPTO_CALC_TREND_STRENGTH            string = "trend_strength"
)