package crypto

import(
        "fmt"
	"slices"
	"regexp"
)

func PrintCryptoList(list []MarketData, id string, timeframes []string, fields []string) {
        frames := GetInputTimeframes(timeframes)
	
        fmt.Println("##########")
	fmt.Printf("# Crypto List: %s", id)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

        for c := range slices.Values(list) {
	        // base crypto information
	        fmt.Printf("    ## Crypto Currency: %s (%s)\n", c.Symbol, c.Name)
		fmt.Println("")
		fmt.Printf("        ### Current Price: %.4f usd\n", c.CurrentPrice)
		fmt.Printf("        ### Highest Price of the Day usd\n: %.4f", c.High24H)
		fmt.Printf("        ### Lowest Price of the Day: %.4f usd\n", c.Low24H)
		fmt.Println("")
		
		// client timeframe preferences
		for t := range slices.Values(frames) {
		        f := GetPriceChange(c, t)
			fmt.Printf("        ### Price change percentage %s: %.2f", t, f)
		}
		fmt.Println("")

                // other -optional- fields
		for field := range slices.Values(fields) {
		        val := GetCoinField(field, c)
			fmt.Printf("%s: %v\n", ToSnakeCase(field), val)
		}
		
		fmt.Println("")
		fmt.Println("")
	}
}

// Source - https://stackoverflow.com/a/56616250
// Posted by Tenusha Guruge, modified by community. See post 'Timeline' for change history
// Retrieved 2025-11-25, License - CC BY-SA 4.0
func ToSnakeCase(str string) string {
        matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   := regexp.MustCompile("([a-z0-9])([A-Z0-9])")

        snake := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
        return matchAllCap.ReplaceAllString(snake, "${1} ${2}")
}
