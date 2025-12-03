package main

import(
        "log"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/clientlogic"
)

func handleCryptoMutate(cs *crypto.CryptoState, cc *crypto.CryptoCache, operation, feature string, args []string) {
        if len(cs.CurrentList) == 0 {
	        log.Println("Client crypto list is empty...")
		return
	}
	
        switch operation {
	case clientlogic.CLIENT_RANK:
	        controlCryptoRank(cs, args)
	        list := crypto.RankCoins(cs.CurrentTimeframe, cs.CurrentOrder, cs.CurrentList)
		updateCurrentList(cs, list)
	case clientlogic.CLIENT_GROUP:
	        switch feature {
		case clientlogic.CRYPTO_GROUP_LIQUIDITY:
		        controlLiquidityArguments(cs, args)
			list := crypto.GroupHighLiquidityCoins(cs.CurrentMinRank, cs.CurrentMaxRank, cs.CurrentMinVolume, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_GROUP_SCARCITY:
		        controlScarcityArguments(cs, args)
			list := crypto.RankCoinScarcity(cs.CurrentMinCirculatingSupply, cs.CurrentMaxATHChangePercentage, cs.CurrentList)
			updateCurrentList(cs, list)
		default:
		        log.Println("### Available Crypto Grouping Features:")
			log.Printf("     %s\n", clientlogic.CRYPTO_GROUP_LIQUIDITY)
			log.Printf("     %s\n", clientlogic.CRYPTO_GROUP_SCARCITY)
			log.Fatal("Invalid crypto grouping feature.")
		}
	case clientlogic.CLIENT_FILTER:
	        switch feature {
		case clientlogic.CRYPTO_FILTER_TOTAL_VOLUME:
		        controlFilterTotalVolume(cs, args)
			list := crypto.FilterCoinVolume(cs.CurrentMinVolume, cs.CurrentMaxVolume, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FILTER_MARKET_CAP:
		        controlFilterMarketCap(cs, args)
			list := crypto.FilterCoinCap(cs.CurrentMinMarketCap, cs.CurrentMaxMarketCap, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FILTER_PRICE_CHANGE_PERCENTAGE:
		        controlFilterPriceChangePercentage(cs, args)
			list := crypto.FilterCoinPriceChange(cs.CurrentMinPriceChangePercentage, cs.CurrentMaxPriceChangePercentage, cs.CurrentTimeframe, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FILTER_VOLATILE:
		        controlFilterVolatile(cs, args)
			list := crypto.FindWildSwingCoins(cs.CurrentMinSwingScore, cs.CurrentMaxSwingScore, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FILTER_HIGH_RISK:
		        controlFilterHighRisk(cs, args)
			list := crypto.FlagRiskCoins(cs.CurrentMaxATHChangePercentage, cs.CurrentMaxVolume, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FILTER_LOW_RISK:
		        controlFilterLowRisk(cs, args)
			list := crypto.FlagSafeCoins(cs.CurrentMaxRank, cs.CurrentMaxPriceChangePercentage, cs.CurrentTimeframe, cs.CurrentList)
			updateCurrentList(cs, list)
		default:
		        log.Println("### Available Crypto Filtering Features:")
			log.Printf("     %s\n", clientlogic.CRYPTO_FILTER_TOTAL_VOLUME)
			log.Printf("     %s\n", clientlogic.CRYPTO_FILTER_MARKET_CAP)
			log.Printf("     %s\n", clientlogic.CRYPTO_FILTER_PRICE_CHANGE_PERCENTAGE)
			log.Printf("     %s\n", clientlogic.CRYPTO_FILTER_VOLATILE)
			log.Printf("     %s\n", clientlogic.CRYPTO_FILTER_HIGH_RISK)
			log.Printf("     %s\n", clientlogic.CRYPTO_FILTER_LOW_RISK)
			log.Fatal("Invalid crypto filtering feature.")
		
		}
	case clientlogic.CLIENT_FIND:
	        switch feature {
		case clientlogic.CRYPTO_FIND_NEW_HIGH_PRICE:
		        list := findNewPrice(cs, cc, crypto.CoinsNewHigh)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FIND_NEW_LOW_PRICE:
		        list := findNewPrice(cs, cc, crypto.CoinsNewLow)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FIND_HIGH_PRICE_SPIKE:
		        controlHighPriceSpike(cs, args)
			list := crypto.CoinsHighPriceSpike(cs.CurrentMinPriceChangePercentage, cs.CurrentTimeframe, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FIND_POTENTIAL_RALLY:
		        controlFindPotentialRally(cs, args)
			list := crypto.CoinsGetCloseAthChange(cs.CurrentMaxATHChangePercentage, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_FIND_COIN_INFLATION:
		        controlFindCoinInflation(cs, args)
			list := crypto.CoinsHighCirculatingSupply(cs.CurrentMinRank, cs.CurrentMinSupply, cs.CurrentIgnoredCoins, cs.CurrentList)
			updateCurrentList(cs, list)
		default:
		        log.Println("### Available Crypto Searching Features:")
			log.Printf("     %s\n", clientlogic.CRYPTO_FIND_NEW_HIGH_PRICE)
			log.Printf("     %s\n", clientlogic.CRYPTO_FIND_NEW_LOW_PRICE)
			log.Printf("     %s\n", clientlogic.CRYPTO_FIND_HIGH_PRICE_SPIKE)
			log.Printf("     %s\n", clientlogic.CRYPTO_FIND_POTENTIAL_RALLY)
			log.Printf("     %s\n", clientlogic.CRYPTO_FIND_COIN_INFLATION)
			log.Fatal("Invalid crypto searching feature.")
		
		}
	case clientlogic.CLIENT_CALC:
	        switch feature {
		case clientlogic.CRYPTO_CALC_VOLATILITY:
		        controlCalcVolatility(cs, args)
			list := crypto.CalcCoinVolatility(cs.CurrentMinVolatility, cs.CurrentMaxVolatility, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_CALC_GROWTH_POTENTIAL:
		        controlCalcGrowthPotential(cs, args)
			list := crypto.EstimateCoinUpsidePotential(cs.CurrentMinGrowthPotential, cs.CurrentMaxRank, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_CALC_LIQUIDITY:
		        controlCalcLiquidity(cs, args)
			list := crypto.CalcCoinLiquidity(cs.CurrentMinLiquidity, cs.CurrentList)
			updateCurrentList(cs, list)
		case clientlogic.CRYPTO_CALC_TREND_STRENGTH:
		        controlCalcTrendStrength(cs, args)
			list, err := crypto.CheckRealTrend(cs.CurrentTimeframe, cs.CurrentList)
			if err != nil {
			        log.Fatal(err)
			}
			updateCurrentList(cs, list)
		default:
		        log.Println("### Available Crypto Calculation Features:")
			log.Printf("     %s\n", clientlogic.CRYPTO_CALC_VOLATILITY)
			log.Printf("     %s\n", clientlogic.CRYPTO_CALC_GROWTH_POTENTIAL)
			log.Printf("     %s\n", clientlogic.CRYPTO_CALC_LIQUIDITY)
			log.Printf("     %s\n", clientlogic.CRYPTO_CALC_TREND_STRENGTH)
			log.Fatal("Invalid crypto calculation feature.")
		
		}
	default:
	        log.Println("## Available Mutate Client Operations:")
		log.Printf("    %s\n", clientlogic.CLIENT_RANK)
		log.Printf("    %s\n", clientlogic.CLIENT_GROUP)
		log.Printf("    %s\n", clientlogic.CLIENT_FILTER)
		log.Printf("    %s\n", clientlogic.CLIENT_FIND)
		log.Printf("    %s\n", clientlogic.CLIENT_CALC)
	        log.Fatal("Invalid client operation.")
        }
}

func updateCurrentList(cs *crypto.CryptoState, list []crypto.MarketData) {
	cs.UpdateCurrentList(cs.CurrentListID, list)
}
