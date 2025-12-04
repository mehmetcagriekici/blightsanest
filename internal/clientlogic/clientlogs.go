package clientlogic

import(
        "fmt"
)

func PrintClientIntroduction() {
        fmt.Println("Welcome to the BlightSanest Client...")
	fmt.Println("To see the client manual:      * manual")
	fmt.Println("To see the available commands: * help")
	fmt.Println("To quit:                       * quit")
}

func PrintClientHelp() {
        fmt.Println("# Available client commands: -To see more details please see the manual or README.md-")
	fmt.Println("## CLI helper commands:")
	fmt.Println("")
	fmt.Println("    * help: print available client commands.")
	fmt.Println("    * manual: print the client manual.")
	fmt.Println("")
	fmt.Println("## Asset Features Commands:")
	fmt.Println("")
	fmt.Println("    * fetch:  fetch the asset's data from the server if available.")
	fmt.Println("    * get:    get the asset's data that is published from another client.")
	fmt.Println("    * list:   print the IDs of the existing lists in the client")
	fmt.Println("    * save:   save the current list to the client cache and publish it for other clients.")
	fmt.Println("    * rank:   sort the assets.")
	fmt.Println("    * group:  get a group of assets by a specific financial feature.")
	fmt.Println("    * filter: filter the assets with available fields.")
	fmt.Println("    * find:   find an asset or assets with a specific field.")
	fmt.Println("    * calc:   get assets after calculating for available credentials.")
	fmt.Println("")
	fmt.Println("Too see more details, please see the manual.")
}

func PrintClientManual() {
        fmt.Println("# BlightSanest Client Manual")
	fmt.Println("")
	fmt.Println("To see the available client commands: * help")
	fmt.Println("")
	fmt.Println("## Available Assets:")
	fmt.Println("")
	fmt.Println("### Crypto Currencies:")
	fmt.Println("")
	PrintCryptoHelp()
}
