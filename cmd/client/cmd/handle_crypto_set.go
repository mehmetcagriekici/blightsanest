package cmd

import (
        "log"
	"slices"
	"reflect"
	"strconv"

	"github.com/spf13/cobra"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var setCryptoCmd = &cobra.Command{
	Use: "crypto [args...]",
	Short: "Set a crypto field preference."
	Args: cobra.MinimumNArgs(1),
	Run: handleCryptoSet,
}

func handleCryptoSet(cmd *cobra.Command, args []string) {
        // add and remove fields
        stack := []string{}

        // client frames
	frames := []crypto.AvailableTimeframes{}
        // ignored coins
	coins := []string{}

        // iterate over the args
	for snake := range slices.Values(args) {
	     // check if it's a state field
	     if reflect.ValueOf(CryptoState).FieldByName(crypto.ToCamelCase(snake)).IsValid() {
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
		             updateOrder(CryptoState, snake)
                     case "CurrentSortingField":
		             updateField(CryptoState, snake)
		     case "ClientTimeframes":
		             // more than one
			     frames = slices.Insert(frames, len(frames), crypto.GetInputTimeframes([]string{snake})[0])
			     CryptoState.UpdateClientTimeframes(frames)
		     case "CurrentTimeframe":
		             frame := crypto.GetInputTimeframes([]string{snake})[0]
			     CryptoState.UpdateCurrentTimeframe(frame)
		     case "CurrentMaxRank":
		             maxRank, err := strconv.Atoi(snake)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateMarketRank(CryptoState.CurrentMinRank, maxRank)
		     case "CurrentMinRank":
		             minRank, err := strconv.Atoi(snake)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateMarketRank(minRank, CryptoState.CurrentMaxRank)
		     case "CurrentMaxVolume":
		             maxVolume, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateVolume(CryptoState.CurrentMinVolume, maxVolume)
		     case "CurrentMinVolume":
		             minVolume, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateVolume(minVolume, CryptoState.CurrentMaxVolume)
		     case "CurrentMinCirculatingSupply":
		             minCirculatingSupply, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateCirculatingSupply(minCirculatingSupply)
		     case "CurrentMaxAthChangePercentage":
		             maxAthChangePercentage, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateAthChangePercentage(maxAthChangePercentage)
		     case "CurrentMinMarketCap":
		             minMarketCap, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateMarketCap(minMarketCap, CryptoState.CurrentMaxMarketCap)
		     case "CurrentMaxMarketCap":
		             maxMarketCap, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateMarketCap(CryptoState.CurrentMinMarketCap, maxMarketCap)
		     case "CurrentMinPriceChangePercentage":
		             minPriceChangePercentage, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdatePriceChangePercentage(minPriceChangePercentage, CryptoState.CurrentMaxPriceChangePercentage)
		     case "CurrentMaxPriceChangePercentage":
		              maxPriceChangePercentage, err := strconv.ParseFloat(snake, 64)
			      if err != nil {
			              log.Fatal(err)
			      }
			      CryptoState.UpdatePriceChangePercentage(CryptoState.CurrentMinPriceChangePercentage, maxPriceChangePercentage)
		     case "CurrentMinSwingScore":
		             minSwingScore, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateCurrentSwingScore(minSwingScore, CryptoState.CurrentMaxSwingScore)
		     case "CurrentMaxSwingScore":
		             maxSwingScore, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateCurrentSwingScore(CryptoState.CurrentMinSwingScore, maxSwingScore)
		     case "CurrentIgnoredCoins":
		             // multiple
			     coins = slices.Insert(coins, len(coins), snake)
		     case "CurrentMinSupply":
		             minSupply, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateSupply(minSupply)
		     case "CurrentMinVolatility":
		             minVolatility, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateCurrentVolatility(minVolatility, CryptoState.CurrentMaxVolatility)
		     case "CurrentMaxVolatility":
		             maxVolatility, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateCurrentVolatility(CryptoState.CurrentMinVolatility, maxVolatility)
		     case "CurrentMinGrowthPotential":
		             minGrowthPotential, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateGrowthPotential(minGrowthPotential)
		     case "CurrentMinLiquidity":
		             minLiquidity, err := strconv.ParseFloat(snake, 64)
			     if err != nil {
			             log.Fatal(err)
			     }
			     CryptoState.UpdateCurrentLiquidity(minLiquidity)
		     default:
		             log.Println("Invalid fieldname.")
		     }
	     }
	}
}
