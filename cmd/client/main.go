package main

import "github.com/mehmetcagriekici/blightsanest/cmd/client/cmd"

func main() {
	cmd.Execute()
}


func main() {
        for {
	        // Get data from the server
		if words[0] == clientlogic.CLIENT_FETCH {
		        if words[1] == clientlogic.ASSET_CRYPTO {
				handleCryptoFetch(cryptoCache,
				                  cryptoState,
						  conn,
						  words[2:],
						  cryptoSubscriptionManager)
                                continue
			}

                        continue
		}

                // get data from other clients
		if words[0] == clientlogic.CLIENT_GET {
		        if words[1] == clientlogic.ASSET_CRYPTO {
			        handleCryptoGet(cryptoCache,
				                cryptoState,
						conn,
						words[2:],
						cryptoSubscriptionManager)
				continue
			}

                        continue
		}

                // ranking features
                if words[0] == clientlogic.CLIENT_RANK {
		        if words[1] == clientlogic.ASSET_CRYPTO {
			        handleCryptoRank(cryptoState, words[2:])
				continue
			}
		}

                // grouping features
		if words[0] == clientlogic.CLIENT_GROUP {
		        if !crypto.ControlFeatureSub(words) {
			        continue
			}

		        if words[1] == clientlogic.ASSET_CRYPTO {
			        switch words[2] {
				case clientlogic.CRYPTO_GROUP_LIQUIDITY:
			               	handleCryptoGroupLiquidity(cryptoState, words[3:])
				case clientlogic.CRYPTO_GROUP_SCARCITY:
					handleCryptoGroupScarcity(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto grouping option. Available: <liquidity> <scarcity>")
				}
				continue
			}
		}

                // filtering features
		if words[0] == clientlogic.CLIENT_FILTER {
		        if !crypto.ControlFeatureSub(words) {
			        continue
			}

                        if words[1] == clientlogic.ASSET_CRYPTO {
                                switch words[2] {
			        case "total_volume":
					handleCryptoFilterTotalVolume(cryptoState, words[3:])
			        case "market_cap":
					handleCryptoFilterMarketCap(cryptoState, words[3:])
				case "price_change_percentage":
					handleCryptoFilterPriceChangePercentage(cryptoState, words[3:])
				case "volatile":
					handleCryptoFilterVolatile(cryptoState, words[3:])
				case "high_risk":
					handleCryptoFilterHighRisk(cryptoState, words[3:])
				case "low_risk":
					handleCryptoFilterLowRisk(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto filtering option. Available: <total_volume> <market_cap> <price_change_percentage> <volatile> <high_risk> <low_risk>")
			        }

			continue
		        }
		}

                // searcing features
		if words[0] == clientlogic.CLIENT_FIND {
		        if !crypto.ControlFeatureSub(words) {
			        continue
			}

                        if words[1] == clientlogic.ASSET_CRYPTO {
			        switch words[2] {
				case clientlogic.CRYPTO_FIND_NAME:
				        handleCryptoFindName(cryptoState, words[3])
				case clientlogic.CRYPTO_FIND_NEW_HIGH_PRICE:
				        handleCryptoNewHighPrice(cryptoState, cryptoCache)
				case clientlogic.CRYPTO_FIND_NEW_LOW_PRICE:
				        handleCryptoNewLowPrice(cryptoState, cryptoCache)
				case clientlogic.CRYPTO_FIND_HIGH_PRICE_SPIKE:
					handleCryptoNewPriceSpike(cryptoState, words[3:])
				case clientlogic.CRYPTO_FIND_POTENTIAL_RALLY:
					handleCryptoFindPotentialRally(cryptoState, words[3:])
				case clientlogic.CRYPTO_FIND_COIN_INFLATION:
				        handleCryptoFindCoinInflation(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto search command. Available: <name>, <new_high_price>, <high_price_spike>, <potential_rally>, <coin_inflation>")
				}
				continue
			}
		}

                // calculating features
		if words[0] == clientlogic.CLIENT_CALC {
		        if !crypto.ControlFeatureSub(words) {
			        continue
			}

                        if words[1] == clientlogic.ASSET_CRYPTO {
			        switch words[2] {
				case clientlogic.CRYPTO_CALC_VOLATILITY:
				        handleCryptoCalcVolatility(cryptoState, words[3:])
				case clientlogic.CRYPTO_CALC_GROWTH_POTENTIAL:
					handleCryptoCalcGrowthPotential(cryptoState, words[3:])
				case clientlogic.CRYPTO_CALC_LIQUIDITY:
					handleCryptoCalcLiquidity(cryptoState, words[3:])
				case clientlogic.CRYPTO_CALC_TREND_STRENGTH:
					handleCryptoCalcTrendStrength(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto calculation command. Available: <volatility>, <growth_potential>, <liquidity>, <trend_strength>")
				}
				continue
			}
		}

	}
}
