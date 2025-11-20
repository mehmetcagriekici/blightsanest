package clientlogic

import(
        "fmt"
)

func PrintClientIntroduction() {
        fmt.Println("Welcome to the BlightSanest Client...")
	fmt.Println("To see the client manual: * manual")
}

func PrintClientHelp() {
        fmt.Println("# Available client commands:")
	fmt.Println("--- manual: See client manual")
	fmt.Println("--- get: Load assets data to the client <get crypto>")
}

func PrintClientManual() {
        fmt.Println("### BlightSanest Client Manual ###")
	fmt.Println("")
	fmt.Println("To see the available client commands: * help")
	fmt.Println("")
	fmt.Println("")
}