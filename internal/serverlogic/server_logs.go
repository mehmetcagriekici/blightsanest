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
	fmt.Println("--- To see the available commands: * help")
	fmt.Println("--- To fetch crypto data with custom queries: * fetch crypto ...")
}

func PrintServerQuit() {
        fmt.Println("Quiting the server...")
}