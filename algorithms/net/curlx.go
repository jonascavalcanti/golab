package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Get command-line arguments excluding the program name
	args := os.Args[1:]

	// Define the base URL prefix
	basePath := "https://"

	// Iterate through each URL provided as argument
	for _, url := range args {
		// Check if the URL contains the base path, if not prepend it
		if !strings.HasPrefix(url, basePath) {
			url = basePath + url
		}

		fmt.Printf("Checking URL: %s\n", url)

		// Perform HTTP GET request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error occurred while fetching URL: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Print status code
		fmt.Println("Status Code:", resp.StatusCode)

		// Copy response body to standard output
		bytesCopied, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Error occurred while reading response body: %v\n", err)
			os.Exit(1)
		}

		// Print the number of bytes copied
		fmt.Printf("\nBytes copied: %d\n", bytesCopied)
	}

}
