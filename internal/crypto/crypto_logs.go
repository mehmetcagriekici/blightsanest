package crypto

import(
        "fmt"
	"slices"
)

func PrintCryptoList(list []MarketData, id string, timeframes []string) {
        frames := GetInputTimeFrames(timeframes)
	
        fmt.Println("##########")
	fmt.Printf("# Crypto List: %s", id)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

        for c := range slices.Values(list) {
	        fmt.Printf("    ## Crypto Currency: %s (%s)\n", c.Symbol, c.Name)
		fmt.Println("")
		fmt.Printf("        ### Current Price: %.4f usd\n", c.CurrentPrice)
		fmt.Printf("        ### Highest Price of the Day usd\n: %.4f", c.High24H)
		fmt.Printf("        ### Lowest Price of the Day: %.4f usd\n", c.Low24H)
		fmt.Println("")
		for t := range slices.Values(frames) {
		        f := GetPriceChange(c, t)
			fmt.Printf("        ### Price change percentage %s: %.2f%", t, f)
		}
		fmt.Println("")
		fmt.Println("")
	}
}