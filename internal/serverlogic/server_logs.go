package serverlogic

import(
        "fmt"
)

// server introduction
func PrintServerIntroduction() {
        fmt.Println("Welcome to the BlightSanest Server! Your local publisher...")
	PrintServerHelp()
}

func PrintServerHelp() {
        fmt.Println("Available Server Features and Commands:")
	fmt.Println("--- To quit the current session: * quit")
	fmt.Println("--- To fetch crypto data with custom queries: * fetch crypto <1h/24h/7d/30d/200d/1y - one or many with one space>")
}

func PrintServerQuit() {
        fmt.Println("Quiting the server...")
	fmt.Println("You no longer have access to the server from your clients!")
}