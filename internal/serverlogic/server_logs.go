package serverlogic

import(
        "fmt"
)

// server introduction
func PrintServerIntroduction() {
        fmt.Println("Welcome to the BlightSanest Publisher...")
	PrintAvailableTimeframes()
}

func PrintAvailableTimeframes() {
        fmt.Println("Available timeframes:")
	fmt.Println("One hour: 1H")
	fmt.Println("One day: 24H")
	fmt.Println("One week: 7D")
	fmt.Println("One month: 30D")
	fmt.Println("Two hundred days: 200D")
	fmt.Println("One year: 1Y")
}