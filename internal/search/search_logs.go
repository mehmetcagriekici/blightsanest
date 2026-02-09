package search

import(
	"fmt"
)

func PrintSearchHelp() {
	fmt.Println("BlightSanest Search Engine Manual")
	fmt.Println("The purpose of this search engine is that giving the user absolute freedom with their queries.")
	fmt.Println("  ## Main Features")
	fmt.Println("")
	fmt.Println("  ### Keyword Search")
	fmt.Println("")
	fmt.Println("  - To use the Keyword Search on existing database index: * search keyword <asset_type> <your search query>")
	fmt.Println(" - Example: search keyword crypto find the tokens with 100 dollars current price.")
	fmt.Println("")
	fmt.Println("> Keyword search uses BM25 algorithm. Displays up to 5 results. Keyword search is not semantic, it will look for the exact query matches.")
	fmt.Println("")
	fmt.Println("To use the Semantic Search on existing database index: * search semantic <limit> <query>")
	fmt.Println(" - Example: search 10 crypto find the tokens with 100 dollars current price.")
	fmt.Println("  ### Creating Database Indexes - must be created before search")
	fmt.Println("")
	fmt.Println("  - Create an Inverted Index for the keyword search from the entire database: * create_inverted_index")
	fmt.Println("")
	fmt.Println("  - Create a Semantic Search Index from the database: * create_semantic_index")
}
