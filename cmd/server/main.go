package main

import(
	"log"

        "github.com/joho/godotenv"
)

func main() {
        // load environment variables
	if err := godotenv.Load(); err != nil {
	        log.Fatal(err)
	}
}