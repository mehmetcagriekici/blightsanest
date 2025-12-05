package main

import(
        "log"
	"slices"
	"reflect"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoSet(cs *crypto.CryptoState, args []string) {
        // add and remove fields
        stack := []string{}

        // client frames
	frames := []crypto.AvailableTimeframes{}
        // ignored coins
	coins := []string{}

        // iterate over the args
	for snake := range slices.Values(args) {
	     // check if it's a state field
	     if reflect.ValueOf(cs).FieldByName(crypto.ToCamelCase(snake)).IsValid() {
	             // if stack does already have a fieldname remove it
		     if len(stack) != 0 {
		             stack = slices.Delete(stack, 0, 1)
		     }
		     
	             // add fieldname to the stack
	             stack = slices.Insert(stack, 0, crypto.ToCamelCase(snake))
	     } else {
	             // value
		     // get the field name from the stack
		     // ClientTimeFrames and CurrentIgnoredCoins -> multiple values
		     switch fieldName := stack[0]; fieldName {
		     case "CurrentOrder":
		             updateOrder(cs, snake)
                     case "CurrentSortingField":
		             updateField(cs, snake)
		     case "ClientTimeframes":
		             // more than one
			     frames = slices.Insert(frames, len(frames), crypto.GetInputTimeframes([]string{snake})[0])
			     cs.UpdateClientTimeframes(frames)
		     case "CurrentTimeframe":
		             frame := crypto.GetInputTimeframes([]string{snake})[0]
			     cs.UpdateCurrentTimeframe(frame)
		     case "CurrentMaxRank":
		             maxRank, err := strconv.Atoi(snake)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
		     case "CurrentMinRank":
		             minRank, err := strconv.Atoi(snake)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateMarketRank(minRank, cs.CurrentMaxRank)
		     case "CurrentMaxVolume":
		             maxVolume, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateVolume(cs.CurrentMinVolume, maxVolume)
		     case "CurrentMinVolume":
		             minVolume, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateVolume(minVolume, cs.CurrentMaxVolume)
		     case "CurrentMinCirculatingSupply":
		             minCirculatingSupply, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateCirculatingSupply(minCirculatingSupply)
		     case "CurrentMaxAthChangePercentage":
		             maxAthChangePercentage, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateAthChangePercentage(maxAthChangePercentage)
		     case "CurrentMinMarketCap":
		             minMarketCap, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateMarketCap(minMarketCap, cs.CurrentMaxMarketCap)
		     case "CurrentMaxMarketCap":
		             maxMarketCap, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateMarketCap(cs.CurrentMinMarketCap, maxMarketCap)
		     case "CurrentMinPriceChangePercentage":
		             minPriceChangePercentage, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdatePriceChangePercentage(minPriceChangePercentage, cs.CurrentMaxPriceChangePercentage)
		     case "CurrentMaxPriceChangePercentage":
		              maxPriceChangePercentage, err := strconv.ParseFloat(snake, 64)
			      if err != nil {
			              log.Fatal(err)
			      }
			      cs.UpdatePriceChangePercentage(cs.CurrentMinPriceChangePercentage, maxPriceChangePercentage)
		     case "CurrentMinSwingScore":
		             minSwingScore, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateCurrentSwingScore(minSwingScore, cs.CurrentMaxSwingScore)
		     case "CurrentMaxSwingScore":
		             maxSwingScore, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateCurrentSwingScore(cs.CurrentMinSwingScore, maxSwingScore)
		     case "CurrentIgnoredCoins":
		             // multiple
			     coins = slices.Insert(coins, len(coins), snake)
		     case "CurrentMinSupply":
		             minSupply, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateSupply(minSupply)
		     case "CurrentMinVolatility":
		             minVolatility, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateCurrentVolatility(minVolatility, cs.CurrentMaxVolatility)
		     case "CurrentMaxVolatility":
		             maxVolatility, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateCurrentVolatility(cs.CurrentMinVolatility, maxVolatility)
		     case "CurrentMinGrowthPotential":
		             minGrowthPotential, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateGrowthPotential(minGrowthPotential)
		     case "CurrentMinLiquidity":
		             minLiquidity, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     cs.UpdateCurrentLiquidity(minLiquidity)
		     default:
		             log.Println("Invalid fieldname.")
		     }
	     }
	}
}