package cmd

import(
	"fmt"
        "log"
	"strconv"

	"github.com/spf13/cobra"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var calcCryptoGrowthPotentialCmd = &cobra.Command{
	Use:   "crypto growth_potential [args...]",
	Short: "Calculate coin growth potentials, and get a rage with min growth preference, and a max market rank preference.",
	Run:    handleCryptoCalcGrowthPotential,
}

func handleCryptoCalcGrowthPotential(cmd *cobra.Command, args []string) {
        controlCalcGrowthPotential(CryptoState, args)
        list := crypto.EstimateCoinUpsidePotential(CryptoState.CurrentMinGrowthPotential,
		CryptoState.CurrentMaxRank,
		CryptoState.CurrentList)
        newID := fmt.Sprintf("calc_growth_potential_%f_%d",
		CryptoState.CurrentMinGrowthPotential,
		CryptoState.CurrentMaxRank)
	fields := []string{"ATH", "AthChangePercentage"}
	commonCryptoHandler(CryptoState, list, fields, newID)
}

func controlCalcGrowthPotential(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the client preferences for min growth potential and max market cap rank values.")
		log.Printf("Min Growth Potential: %f\n", cs.CurrentMinGrowthPotential)
		log.Printf("Max Market Cap Rank: %d\n", cs.CurrentMaxRank)
	case 1:
	        log.Println("One argument is passed. Updating the client min growth potential preference...")
		log.Printf("Max Market Cap Rank: %d\n", cs.CurrentMaxRank)
		minPotential, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateGrowthPotential(minPotential)
	case 2:
	        log.Println("Updating the client min growth potential and max market cap rank preferences...")
		minPotential, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxRank, err := strconv.Atoi(args[1])
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateGrowthPotential(minPotential)
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
	default:
	        log.Println("Invalid use of command: calc crypto growth_potential <min_potential float64> <max_market_cap_rank int>")
	}
}
