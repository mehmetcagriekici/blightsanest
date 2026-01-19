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
	fmt.Println("  - To use the Keyword Search on existing database index: * search keyword <your search query>")
	fmt.Println("  ### Create Database Index")
	fmt.Println("")
	fmt.Println("  - Create an Inverted Index for the entire database: * create_inverted_index")
	fmt.Println("  ! This might take a while depending on the size of your database. This will create an inverted index for the entire database -literally-. But afterwards you will be able to search for anything from the database with any query.")
}
