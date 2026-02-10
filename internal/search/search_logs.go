package search

import(
	"fmt"
)

func PrintSearchHelp() {
	fmt.Println("BlightSanest Search Engine Manual")
	fmt.Println("The purpose of this search engine is that giving the user absolute freedom with their queries.")
	fmt.Println("")
	fmt.Println("## Search")
	fmt.Println("")
	fmt.Println("To use the search on the database:")
	fmt.Println("")
	fmt.Println("* search <limit int> <your query ...string>")
	fmt.Println("")
	fmt.Println("After calculating BM25 and cosine similarity scores separately for keyword and semantic search. Hybrid search creates an rrf score based on the ranks of the results, and presents the first <limit> results.")
	fmt.Println("")
	fmt.Println("Before using the search feature please built the inverted and semantic search indexes...")
	fmt.Println("")
	fmt.Println("  ## Internal Sub Search Algorithms")
	fmt.Println("")
	fmt.Println("  ### Keyword Search")
	fmt.Println("")
	fmt.Println("> Keyword search uses BM25 algorithm. Displays up to 5 results. Keyword search is not semantic, it will look for the exact query matches. It's internally built with Go.")
	fmt.Println("")
	fmt.Println("  ### Semantic Search")
	fmt.Println("> Semantic Search uses sentence transormers mini model with cosine similarity. It's an external http server built with python and fastapi.")
	fmt.Println("  ### Creating Database Indexes - each must be individually created before search")
	fmt.Println("")
	fmt.Println("  - Create an Inverted Index for the keyword search from the entire database: * create_inverted_index")
	fmt.Println("")
	fmt.Println("  - Create a Semantic Search Index from the database: * create_semantic_index")
	fmt.Println("")
}
