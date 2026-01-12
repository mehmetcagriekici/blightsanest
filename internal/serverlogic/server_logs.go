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
	fmt.Println("--- To get a crypto list from the database * get crypto <crypto_list_id>")
	fmt.Println("--- To save a crypto list from the server cache to database with a custom id * save crypto <crypto_list_cache_id> <custom_id>")
}

func PrintServerQuit() {
        fmt.Println("Quiting the server...")
}
